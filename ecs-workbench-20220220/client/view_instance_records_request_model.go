// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewInstanceRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ViewInstanceRecordsRequest
	GetInstanceId() *string
	SetRegionId(v string) *ViewInstanceRecordsRequest
	GetRegionId() *string
	SetTerminalSessionToken(v string) *ViewInstanceRecordsRequest
	GetTerminalSessionToken() *string
}

type ViewInstanceRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	TerminalSessionToken *string `json:"TerminalSessionToken,omitempty" xml:"TerminalSessionToken,omitempty"`
}

func (s ViewInstanceRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewInstanceRecordsRequest) GoString() string {
	return s.String()
}

func (s *ViewInstanceRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ViewInstanceRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewInstanceRecordsRequest) GetTerminalSessionToken() *string {
	return s.TerminalSessionToken
}

func (s *ViewInstanceRecordsRequest) SetInstanceId(v string) *ViewInstanceRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ViewInstanceRecordsRequest) SetRegionId(v string) *ViewInstanceRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *ViewInstanceRecordsRequest) SetTerminalSessionToken(v string) *ViewInstanceRecordsRequest {
	s.TerminalSessionToken = &v
	return s
}

func (s *ViewInstanceRecordsRequest) Validate() error {
	return dara.Validate(s)
}
