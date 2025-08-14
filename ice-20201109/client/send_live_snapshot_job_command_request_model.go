// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveSnapshotJobCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v string) *SendLiveSnapshotJobCommandRequest
	GetCommand() *string
	SetJobId(v string) *SendLiveSnapshotJobCommandRequest
	GetJobId() *string
}

type SendLiveSnapshotJobCommandRequest struct {
	// The operation command.
	//
	// Valid values:
	//
	// 	- stop
	//
	// 	- restart
	//
	// 	- start
	//
	// This parameter is required.
	//
	// example:
	//
	// start
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The ID of the snapshot job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SendLiveSnapshotJobCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s SendLiveSnapshotJobCommandRequest) GoString() string {
	return s.String()
}

func (s *SendLiveSnapshotJobCommandRequest) GetCommand() *string {
	return s.Command
}

func (s *SendLiveSnapshotJobCommandRequest) GetJobId() *string {
	return s.JobId
}

func (s *SendLiveSnapshotJobCommandRequest) SetCommand(v string) *SendLiveSnapshotJobCommandRequest {
	s.Command = &v
	return s
}

func (s *SendLiveSnapshotJobCommandRequest) SetJobId(v string) *SendLiveSnapshotJobCommandRequest {
	s.JobId = &v
	return s
}

func (s *SendLiveSnapshotJobCommandRequest) Validate() error {
	return dara.Validate(s)
}
