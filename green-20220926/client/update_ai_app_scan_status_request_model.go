// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiAppScanStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *UpdateAiAppScanStatusRequest
	GetAppIds() *string
	SetRegionId(v string) *UpdateAiAppScanStatusRequest
	GetRegionId() *string
	SetStatus(v string) *UpdateAiAppScanStatusRequest
	GetStatus() *string
}

type UpdateAiAppScanStatusRequest struct {
	// The application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10d74f5b-6edf-4826-a989-de03463e479d
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status. Valid values:
	//
	// - enable: enabled.
	//
	// - disable: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAiAppScanStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiAppScanStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiAppScanStatusRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *UpdateAiAppScanStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAiAppScanStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAiAppScanStatusRequest) SetAppIds(v string) *UpdateAiAppScanStatusRequest {
	s.AppIds = &v
	return s
}

func (s *UpdateAiAppScanStatusRequest) SetRegionId(v string) *UpdateAiAppScanStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAiAppScanStatusRequest) SetStatus(v string) *UpdateAiAppScanStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateAiAppScanStatusRequest) Validate() error {
	return dara.Validate(s)
}
