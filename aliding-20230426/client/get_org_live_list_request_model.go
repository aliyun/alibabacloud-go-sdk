// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgLiveListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpId(v string) *GetOrgLiveListRequest
	GetCorpId() *string
	SetEndTime(v int64) *GetOrgLiveListRequest
	GetEndTime() *int64
	SetPageNumber(v int64) *GetOrgLiveListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetOrgLiveListRequest
	GetPageSize() *int64
	SetStartTime(v int64) *GetOrgLiveListRequest
	GetStartTime() *int64
	SetTenantContext(v *GetOrgLiveListRequestTenantContext) *GetOrgLiveListRequest
	GetTenantContext() *GetOrgLiveListRequestTenantContext
	SetUserId(v string) *GetOrgLiveListRequest
	GetUserId() *string
}

type GetOrgLiveListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// db4d318xxxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 1720211800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1719211800000
	StartTime     *int64                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *GetOrgLiveListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetOrgLiveListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListRequest) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GetOrgLiveListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetOrgLiveListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetOrgLiveListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetOrgLiveListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetOrgLiveListRequest) GetTenantContext() *GetOrgLiveListRequestTenantContext {
	return s.TenantContext
}

func (s *GetOrgLiveListRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetOrgLiveListRequest) SetCorpId(v string) *GetOrgLiveListRequest {
	s.CorpId = &v
	return s
}

func (s *GetOrgLiveListRequest) SetEndTime(v int64) *GetOrgLiveListRequest {
	s.EndTime = &v
	return s
}

func (s *GetOrgLiveListRequest) SetPageNumber(v int64) *GetOrgLiveListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOrgLiveListRequest) SetPageSize(v int64) *GetOrgLiveListRequest {
	s.PageSize = &v
	return s
}

func (s *GetOrgLiveListRequest) SetStartTime(v int64) *GetOrgLiveListRequest {
	s.StartTime = &v
	return s
}

func (s *GetOrgLiveListRequest) SetTenantContext(v *GetOrgLiveListRequestTenantContext) *GetOrgLiveListRequest {
	s.TenantContext = v
	return s
}

func (s *GetOrgLiveListRequest) SetUserId(v string) *GetOrgLiveListRequest {
	s.UserId = &v
	return s
}

func (s *GetOrgLiveListRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrgLiveListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetOrgLiveListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetOrgLiveListRequestTenantContext) SetTenantId(v string) *GetOrgLiveListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetOrgLiveListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
