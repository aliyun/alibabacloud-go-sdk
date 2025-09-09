// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRTCLiveRoomsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListRTCLiveRoomsRequest
	GetAppId() *string
	SetPageNo(v int32) *ListRTCLiveRoomsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRTCLiveRoomsRequest
	GetPageSize() *int32
}

type ListRTCLiveRoomsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRTCLiveRoomsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRTCLiveRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListRTCLiveRoomsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListRTCLiveRoomsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRTCLiveRoomsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRTCLiveRoomsRequest) SetAppId(v string) *ListRTCLiveRoomsRequest {
	s.AppId = &v
	return s
}

func (s *ListRTCLiveRoomsRequest) SetPageNo(v int32) *ListRTCLiveRoomsRequest {
	s.PageNo = &v
	return s
}

func (s *ListRTCLiveRoomsRequest) SetPageSize(v int32) *ListRTCLiveRoomsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRTCLiveRoomsRequest) Validate() error {
	return dara.Validate(s)
}
