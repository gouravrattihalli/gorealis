// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package aurora

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

const AURORA_EXECUTOR_NAME = "AuroraExecutor"

var ACTIVE_STATES map[ScheduleStatus]bool
var SLAVE_ASSIGNED_STATES map[ScheduleStatus]bool
var LIVE_STATES map[ScheduleStatus]bool
var TERMINAL_STATES map[ScheduleStatus]bool

const GOOD_IDENTIFIER_PATTERN = "^[\\w\\-\\.]+$"
const GOOD_IDENTIFIER_PATTERN_JVM = "^[\\w\\-\\.]+$"
const GOOD_IDENTIFIER_PATTERN_PYTHON = "^[\\w\\-\\.]+$"

var ACTIVE_JOB_UPDATE_STATES map[JobUpdateStatus]bool

const BYPASS_LEADER_REDIRECT_HEADER_NAME = "Bypass-Leader-Redirect"

func init() {
	ACTIVE_STATES = map[ScheduleStatus]bool{
		9:  true,
		17: true,
		6:  true,
		0:  true,
		13: true,
		12: true,
		2:  true,
		1:  true,
		16: true,
	}

	SLAVE_ASSIGNED_STATES = map[ScheduleStatus]bool{
		9:  true,
		17: true,
		6:  true,
		13: true,
		12: true,
		2:  true,
		1:  true,
	}

	LIVE_STATES = map[ScheduleStatus]bool{
		6:  true,
		13: true,
		12: true,
		17: true,
		2:  true,
	}

	TERMINAL_STATES = map[ScheduleStatus]bool{
		4: true,
		3: true,
		5: true,
		7: true,
	}

	ACTIVE_JOB_UPDATE_STATES = map[JobUpdateStatus]bool{
		0:  true,
		1:  true,
		2:  true,
		3:  true,
		9:  true,
		10: true,
	}

}
