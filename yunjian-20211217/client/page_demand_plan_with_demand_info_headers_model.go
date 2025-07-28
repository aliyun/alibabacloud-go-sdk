// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageDemandPlanWithDemandInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PageDemandPlanWithDemandInfoHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *PageDemandPlanWithDemandInfoHeaders
	GetYunUserId() *string
}

type PageDemandPlanWithDemandInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s PageDemandPlanWithDemandInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s PageDemandPlanWithDemandInfoHeaders) GoString() string {
	return s.String()
}

func (s *PageDemandPlanWithDemandInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PageDemandPlanWithDemandInfoHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *PageDemandPlanWithDemandInfoHeaders) SetCommonHeaders(v map[string]*string) *PageDemandPlanWithDemandInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageDemandPlanWithDemandInfoHeaders) SetYunUserId(v string) *PageDemandPlanWithDemandInfoHeaders {
	s.YunUserId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoHeaders) Validate() error {
	return dara.Validate(s)
}
