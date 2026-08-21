// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOnlineHeatmapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetDeviceOnlineHeatmapRequest
	GetDate() *string
	SetDevTag(v string) *GetDeviceOnlineHeatmapRequest
	GetDevTag() *string
	SetSaseUserId(v string) *GetDeviceOnlineHeatmapRequest
	GetSaseUserId() *string
}

type GetDeviceOnlineHeatmapRequest struct {
	// The date to query, in the format yyyyMMdd with a fixed length of 8 characters. Online data is retained for only 8 days. Dates beyond the retention period return an empty list.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20260809
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The terminal device ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2987b3e0-8108-2f99-4d18-3b4f1c1****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The user ID. You can obtain this value from the following operations:
	//
	// - [ListUserDevices](~~ListUserDevices~~): Lists user terminal devices.
	//
	// - [GetUserDevice](~~GetUserDevice~~): Queries the details of a user terminal device.
	//
	// This parameter is required.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
}

func (s GetDeviceOnlineHeatmapRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOnlineHeatmapRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOnlineHeatmapRequest) GetDate() *string {
	return s.Date
}

func (s *GetDeviceOnlineHeatmapRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *GetDeviceOnlineHeatmapRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *GetDeviceOnlineHeatmapRequest) SetDate(v string) *GetDeviceOnlineHeatmapRequest {
	s.Date = &v
	return s
}

func (s *GetDeviceOnlineHeatmapRequest) SetDevTag(v string) *GetDeviceOnlineHeatmapRequest {
	s.DevTag = &v
	return s
}

func (s *GetDeviceOnlineHeatmapRequest) SetSaseUserId(v string) *GetDeviceOnlineHeatmapRequest {
	s.SaseUserId = &v
	return s
}

func (s *GetDeviceOnlineHeatmapRequest) Validate() error {
	return dara.Validate(s)
}
