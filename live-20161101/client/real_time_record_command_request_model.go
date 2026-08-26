// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealTimeRecordCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *RealTimeRecordCommandRequest
	GetAppName() *string
	SetCommand(v string) *RealTimeRecordCommandRequest
	GetCommand() *string
	SetDomainName(v string) *RealTimeRecordCommandRequest
	GetDomainName() *string
	SetOwnerId(v int64) *RealTimeRecordCommandRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RealTimeRecordCommandRequest
	GetRegionId() *string
	SetStreamName(v string) *RealTimeRecordCommandRequest
	GetStreamName() *string
}

type RealTimeRecordCommandRequest struct {
	// The name of the application to which the stream belongs. You can view the AppName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The operation action. Valid values:
	//
	// - **start**: forcibly starts recording. This must be called as the first operation and cannot be called again before stopping.
	//
	// - **stop**: forcibly pauses recording. After the stream interruption delay (180 seconds by default) elapses, a recording is generated. This can only be called after start or restart. To generate the file immediately after calling stop, call cancel_delay.
	//
	// - **cancel_delay**: immediately terminates the wait and generates a recording, completely stopping recording. This must be called after stop to generate the file in advance.
	//
	// - **restart**: forcibly restarts recording. If recording is in progress before restart, a file is immediately generated. This can only be called when the task is in the started or stopped state.
	//
	// This parameter is required.
	//
	// example:
	//
	// start
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The streamer\\"s streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stream name. Make sure that the StreamName is correct. You can view the StreamName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// > This operation supports only single-stream operations and does not support wildcards.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s RealTimeRecordCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RealTimeRecordCommandRequest) GoString() string {
	return s.String()
}

func (s *RealTimeRecordCommandRequest) GetAppName() *string {
	return s.AppName
}

func (s *RealTimeRecordCommandRequest) GetCommand() *string {
	return s.Command
}

func (s *RealTimeRecordCommandRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RealTimeRecordCommandRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RealTimeRecordCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RealTimeRecordCommandRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *RealTimeRecordCommandRequest) SetAppName(v string) *RealTimeRecordCommandRequest {
	s.AppName = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetCommand(v string) *RealTimeRecordCommandRequest {
	s.Command = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetDomainName(v string) *RealTimeRecordCommandRequest {
	s.DomainName = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetOwnerId(v int64) *RealTimeRecordCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetRegionId(v string) *RealTimeRecordCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetStreamName(v string) *RealTimeRecordCommandRequest {
	s.StreamName = &v
	return s
}

func (s *RealTimeRecordCommandRequest) Validate() error {
	return dara.Validate(s)
}
