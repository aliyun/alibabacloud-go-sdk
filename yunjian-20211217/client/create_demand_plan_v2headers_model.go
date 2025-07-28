// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDemandPlanV2Headers
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *CreateDemandPlanV2Headers
	GetYunUserId() *string
}

type CreateDemandPlanV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s CreateDemandPlanV2Headers) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanV2Headers) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDemandPlanV2Headers) GetYunUserId() *string {
	return s.YunUserId
}

func (s *CreateDemandPlanV2Headers) SetCommonHeaders(v map[string]*string) *CreateDemandPlanV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *CreateDemandPlanV2Headers) SetYunUserId(v string) *CreateDemandPlanV2Headers {
	s.YunUserId = &v
	return s
}

func (s *CreateDemandPlanV2Headers) Validate() error {
	return dara.Validate(s)
}
