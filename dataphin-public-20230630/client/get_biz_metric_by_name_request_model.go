// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizMetricByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizMetricByNameQuery(v *GetBizMetricByNameRequestBizMetricByNameQuery) *GetBizMetricByNameRequest
	GetBizMetricByNameQuery() *GetBizMetricByNameRequestBizMetricByNameQuery
	SetOpTenantId(v int64) *GetBizMetricByNameRequest
	GetOpTenantId() *int64
}

type GetBizMetricByNameRequest struct {
	// This parameter is required.
	BizMetricByNameQuery *GetBizMetricByNameRequestBizMetricByNameQuery `json:"BizMetricByNameQuery,omitempty" xml:"BizMetricByNameQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetBizMetricByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameRequest) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameRequest) GetBizMetricByNameQuery() *GetBizMetricByNameRequestBizMetricByNameQuery {
	return s.BizMetricByNameQuery
}

func (s *GetBizMetricByNameRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBizMetricByNameRequest) SetBizMetricByNameQuery(v *GetBizMetricByNameRequestBizMetricByNameQuery) *GetBizMetricByNameRequest {
	s.BizMetricByNameQuery = v
	return s
}

func (s *GetBizMetricByNameRequest) SetOpTenantId(v int64) *GetBizMetricByNameRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBizMetricByNameRequest) Validate() error {
	if s.BizMetricByNameQuery != nil {
		if err := s.BizMetricByNameQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBizMetricByNameRequestBizMetricByNameQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// True
	Draft *bool `json:"Draft,omitempty" xml:"Draft,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Metric1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetBizMetricByNameRequestBizMetricByNameQuery) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameRequestBizMetricByNameQuery) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameRequestBizMetricByNameQuery) GetDraft() *bool {
	return s.Draft
}

func (s *GetBizMetricByNameRequestBizMetricByNameQuery) GetName() *string {
	return s.Name
}

func (s *GetBizMetricByNameRequestBizMetricByNameQuery) SetDraft(v bool) *GetBizMetricByNameRequestBizMetricByNameQuery {
	s.Draft = &v
	return s
}

func (s *GetBizMetricByNameRequestBizMetricByNameQuery) SetName(v string) *GetBizMetricByNameRequestBizMetricByNameQuery {
	s.Name = &v
	return s
}

func (s *GetBizMetricByNameRequestBizMetricByNameQuery) Validate() error {
	return dara.Validate(s)
}
