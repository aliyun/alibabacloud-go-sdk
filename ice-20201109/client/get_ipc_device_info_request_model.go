// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpcDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapability(v string) *GetIpcDeviceInfoRequest
	GetCapability() *string
	SetDeviceId(v string) *GetIpcDeviceInfoRequest
	GetDeviceId() *string
	SetEndTime(v string) *GetIpcDeviceInfoRequest
	GetEndTime() *string
	SetPageNo(v int32) *GetIpcDeviceInfoRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetIpcDeviceInfoRequest
	GetPageSize() *int32
	SetStartTime(v string) *GetIpcDeviceInfoRequest
	GetStartTime() *string
}

type GetIpcDeviceInfoRequest struct {
	// Capability. Default value is understand. Valid values:
	//
	// - understand: understanding
	//
	// - understand-reid: understanding with reid
	//
	// - search: search
	//
	// example:
	//
	// understand
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
	// Device ID.
	//
	// example:
	//
	// d123
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// End Time. UTC time in the format yyyy-MM-ddTHH:mm:ssZ. The difference between EndTime and StartTime must not exceed 180 days.
	//
	// example:
	//
	// 2017-02-11T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Page number for paging. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size, which specifies the number of entries displayed per page. Default value is 10. Maximum value is 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start Time. UTC time in the format yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetIpcDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIpcDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetIpcDeviceInfoRequest) GetCapability() *string {
	return s.Capability
}

func (s *GetIpcDeviceInfoRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetIpcDeviceInfoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetIpcDeviceInfoRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetIpcDeviceInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetIpcDeviceInfoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetIpcDeviceInfoRequest) SetCapability(v string) *GetIpcDeviceInfoRequest {
	s.Capability = &v
	return s
}

func (s *GetIpcDeviceInfoRequest) SetDeviceId(v string) *GetIpcDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *GetIpcDeviceInfoRequest) SetEndTime(v string) *GetIpcDeviceInfoRequest {
	s.EndTime = &v
	return s
}

func (s *GetIpcDeviceInfoRequest) SetPageNo(v int32) *GetIpcDeviceInfoRequest {
	s.PageNo = &v
	return s
}

func (s *GetIpcDeviceInfoRequest) SetPageSize(v int32) *GetIpcDeviceInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetIpcDeviceInfoRequest) SetStartTime(v string) *GetIpcDeviceInfoRequest {
	s.StartTime = &v
	return s
}

func (s *GetIpcDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
