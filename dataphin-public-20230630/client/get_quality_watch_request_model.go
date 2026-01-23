// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetQualityWatchRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetQualityWatchRequest
	GetOpTenantId() *int64
}

type GetQualityWatchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetQualityWatchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchRequest) GoString() string {
	return s.String()
}

func (s *GetQualityWatchRequest) GetId() *int64 {
	return s.Id
}

func (s *GetQualityWatchRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityWatchRequest) SetId(v int64) *GetQualityWatchRequest {
	s.Id = &v
	return s
}

func (s *GetQualityWatchRequest) SetOpTenantId(v int64) *GetQualityWatchRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityWatchRequest) Validate() error {
	return dara.Validate(s)
}
