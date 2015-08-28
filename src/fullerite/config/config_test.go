package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var testConfiguration = `{
    "prefix": "test.",
    "interval": 10,
    "defaultDimensions": {
        "application": "fullerite",
        "host": "dev33-devc"
    },

    "diamond_collectors_path": "src/diamond/collectors",
    "diamond_collectors": {
        "CPUCollector": {"enabled": true, "interval": 10},
        "PingCollector": {"enabled": true, "target_google": "google.com", "interval": 10, "bin": "/bin/ping"}
    },

    "collectors": {
        "Test": {
            "metricName": "TestMetric",
            "interval": 10
        },
        "Diamond":{
            "port": "19191",
            "interval": 10
        }
    },

    "handlers": {
        "Graphite": {
            "server": "10.40.11.51",
            "port": "2003",
            "timeout": 2
        },
        "SignalFx": {
            "authToken": "secret_token",
            "endpoint": "https://ingest.signalfx.com/v2/datapoint",
            "interval": 10,
            "timeout": 2
        }
    }
}
`
var tmpTestFile string

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.ErrorLevel)
	if f, err := ioutil.TempFile("/tmp", "fullerite"); err == nil {
		f.WriteString(testConfiguration)
		tmpTestFile = f.Name()
		f.Close()
		defer os.Remove(tmpTestFile)
	}
	os.Exit(m.Run())
}

func TestGetInt(t *testing.T) {
	assert := assert.New(t)

	val := GetAsInt("10", 123)
	assert.Equal(val, 10)

	val = GetAsInt("notanint", 123)
	assert.Equal(val, 123)

	val = GetAsInt(12.123, 123)
	assert.Equal(val, 12)

	val = GetAsInt(12, 123)
	assert.Equal(val, 12)
}

func TestGetFloat(t *testing.T) {
	assert := assert.New(t)

	val := GetAsFloat("10", 123)
	assert.Equal(val, 10.0)

	val = GetAsFloat("10.21", 123)
	assert.Equal(val, 10.21)

	val = GetAsFloat("notanint", 123)
	assert.Equal(val, 123.0)

	val = GetAsFloat(12.123, 123)
	assert.Equal(val, 12.123)
}

func TestParseExampleConfig(t *testing.T) {
	_, err := ReadConfig(tmpTestFile)
	assert.Nil(t, err, "should succeed")
}
