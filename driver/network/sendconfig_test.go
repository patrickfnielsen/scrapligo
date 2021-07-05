package network_test

import (
	"fmt"
	"github.com/scrapli/scrapligo/driver/core"
	"github.com/scrapli/scrapligo/util/testhelper"
	"strings"
	"testing"
)

func TestSendConfig(t *testing.T) {
	configsMap := platformConfigsMap()

	for _, platform := range core.SupportedPlatforms() {
		f := testhelper.SendConfigTestHelper(platform, strings.Join(configsMap[platform], "\n"))
		t.Run(fmt.Sprintf("Platform=%s", platform), f)
	}
}
