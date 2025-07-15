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
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The action to be performed. Valid values:
	//
	// 	- **start**: forcibly starts recording.
	//
	// 	- **stop**: forcibly stops recording. If the live stream is interrupted for longer than a specific latency, a recording is generated.
	//
	// 	- **cancel_delay**: resets the latency for stream interruption and completely stops recording. If the recording task is stopped when you perform this action, a recording is generated.
	//
	// 	- **restart**: forcibly restarts recording. If the live stream is being recorded when you perform this action, a recording is generated.
	//
	// >  **stop*	- forcibly stops recording. By default, a recording is generated after 180 seconds. **cancel_delay*	- resets the latency for stream interruption from 180 seconds to 0 seconds. This means that a recording is generated immediately.
	//
	// This parameter is required.
	//
	// example:
	//
	// start
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. Make sure that you specify the correct stream name. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
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
