// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetQualityScheduleRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetQualityScheduleRequest
	GetOpTenantId() *int64
}

type GetQualityScheduleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetQualityScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetQualityScheduleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetQualityScheduleRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityScheduleRequest) SetId(v int64) *GetQualityScheduleRequest {
	s.Id = &v
	return s
}

func (s *GetQualityScheduleRequest) SetOpTenantId(v int64) *GetQualityScheduleRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityScheduleRequest) Validate() error {
	return dara.Validate(s)
}
