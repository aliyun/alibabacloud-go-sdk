// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveTranscodeJobCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v string) *SendLiveTranscodeJobCommandRequest
	GetCommand() *string
	SetJobId(v string) *SendLiveTranscodeJobCommandRequest
	GetJobId() *string
}

type SendLiveTranscodeJobCommandRequest struct {
	// The operation command. Only the stop command is supported. This command is used to stop a transcoding job.
	//
	// This parameter is required.
	//
	// example:
	//
	// stop
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The ID of the transcoding job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SendLiveTranscodeJobCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s SendLiveTranscodeJobCommandRequest) GoString() string {
	return s.String()
}

func (s *SendLiveTranscodeJobCommandRequest) GetCommand() *string {
	return s.Command
}

func (s *SendLiveTranscodeJobCommandRequest) GetJobId() *string {
	return s.JobId
}

func (s *SendLiveTranscodeJobCommandRequest) SetCommand(v string) *SendLiveTranscodeJobCommandRequest {
	s.Command = &v
	return s
}

func (s *SendLiveTranscodeJobCommandRequest) SetJobId(v string) *SendLiveTranscodeJobCommandRequest {
	s.JobId = &v
	return s
}

func (s *SendLiveTranscodeJobCommandRequest) Validate() error {
	return dara.Validate(s)
}
