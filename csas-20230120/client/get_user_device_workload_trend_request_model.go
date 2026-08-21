// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeviceWorkloadTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceTag(v string) *GetUserDeviceWorkloadTrendRequest
	GetDeviceTag() *string
	SetFrom(v int64) *GetUserDeviceWorkloadTrendRequest
	GetFrom() *int64
	SetTo(v int64) *GetUserDeviceWorkloadTrendRequest
	GetTo() *int64
	SetWorkloadType(v string) *GetUserDeviceWorkloadTrendRequest
	GetWorkloadType() *string
}

type GetUserDeviceWorkloadTrendRequest struct {
	// The endpoint device ID. You can obtain this value from the following operations:
	//
	// - [GetUserDevice](~~GetUserDevice~~): Queries the details of a user endpoint device.
	//
	// - [ListUserDevices](~~ListUserDevices~~): Queries user endpoint devices in batches.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// The start time of the query time range. This value is a UNIX timestamp in seconds. The value must be greater than or equal to 0 and less than the value of To.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1769998785
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// The end time of the query time range. This value is a UNIX timestamp in seconds. The value must be greater than the value of From.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1771986521
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
	// The workload type. Valid values:
	//
	// - **cpu**: CPU usage.
	//
	// - **mem**: memory usage.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s GetUserDeviceWorkloadTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceWorkloadTrendRequest) GoString() string {
	return s.String()
}

func (s *GetUserDeviceWorkloadTrendRequest) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *GetUserDeviceWorkloadTrendRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetUserDeviceWorkloadTrendRequest) GetTo() *int64 {
	return s.To
}

func (s *GetUserDeviceWorkloadTrendRequest) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *GetUserDeviceWorkloadTrendRequest) SetDeviceTag(v string) *GetUserDeviceWorkloadTrendRequest {
	s.DeviceTag = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendRequest) SetFrom(v int64) *GetUserDeviceWorkloadTrendRequest {
	s.From = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendRequest) SetTo(v int64) *GetUserDeviceWorkloadTrendRequest {
	s.To = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendRequest) SetWorkloadType(v string) *GetUserDeviceWorkloadTrendRequest {
	s.WorkloadType = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendRequest) Validate() error {
	return dara.Validate(s)
}
