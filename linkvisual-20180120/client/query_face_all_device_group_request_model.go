// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIotInstanceId(v string) *QueryFaceAllDeviceGroupRequest
	GetIotInstanceId() *string
	SetIsolationId(v string) *QueryFaceAllDeviceGroupRequest
	GetIsolationId() *string
	SetPageNo(v int32) *QueryFaceAllDeviceGroupRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryFaceAllDeviceGroupRequest
	GetPageSize() *int32
}

type QueryFaceAllDeviceGroupRequest struct {
	// The instance ID.
	//
	// example:
	//
	// TestIotInstanceId
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryFaceAllDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryFaceAllDeviceGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceAllDeviceGroupRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceAllDeviceGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllDeviceGroupRequest) SetIotInstanceId(v string) *QueryFaceAllDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetIsolationId(v string) *QueryFaceAllDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetPageNo(v int32) *QueryFaceAllDeviceGroupRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) SetPageSize(v int32) *QueryFaceAllDeviceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
