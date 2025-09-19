// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalCommandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTerminalCommandsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTerminalCommandsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTerminalCommandsRequest
	GetRegionId() *string
	SetTerminalSessionToken(v string) *ListTerminalCommandsRequest
	GetTerminalSessionToken() *string
}

type ListTerminalCommandsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s ListTerminalCommandsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalCommandsRequest) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTerminalCommandsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTerminalCommandsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTerminalCommandsRequest) GetTerminalSessionToken() *string {
	return s.TerminalSessionToken
}

func (s *ListTerminalCommandsRequest) SetPageNumber(v int32) *ListTerminalCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTerminalCommandsRequest) SetPageSize(v int32) *ListTerminalCommandsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTerminalCommandsRequest) SetRegionId(v string) *ListTerminalCommandsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTerminalCommandsRequest) SetTerminalSessionToken(v string) *ListTerminalCommandsRequest {
	s.TerminalSessionToken = &v
	return s
}

func (s *ListTerminalCommandsRequest) Validate() error {
	return dara.Validate(s)
}
