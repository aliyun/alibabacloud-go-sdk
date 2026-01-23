// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualitySchedulesByWatchIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualitySchedulesByWatchIdRequest
	GetOpTenantId() *int64
	SetWatchId(v int64) *GetQualitySchedulesByWatchIdRequest
	GetWatchId() *int64
}

type GetQualitySchedulesByWatchIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s GetQualitySchedulesByWatchIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualitySchedulesByWatchIdRequest) GoString() string {
	return s.String()
}

func (s *GetQualitySchedulesByWatchIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualitySchedulesByWatchIdRequest) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualitySchedulesByWatchIdRequest) SetOpTenantId(v int64) *GetQualitySchedulesByWatchIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdRequest) SetWatchId(v int64) *GetQualitySchedulesByWatchIdRequest {
	s.WatchId = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdRequest) Validate() error {
	return dara.Validate(s)
}
