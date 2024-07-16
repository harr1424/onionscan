package protocol

import (
	"github.com/harr1424/onionscan/config"
	"github.com/harr1424/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
