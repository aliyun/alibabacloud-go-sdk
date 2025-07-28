// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUrgentDemandPlanDetailHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *GetUrgentDemandPlanDetailHeaders
	GetYunUserId() *string
}

type GetUrgentDemandPlanDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s GetUrgentDemandPlanDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUrgentDemandPlanDetailHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *GetUrgentDemandPlanDetailHeaders) SetCommonHeaders(v map[string]*string) *GetUrgentDemandPlanDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUrgentDemandPlanDetailHeaders) SetYunUserId(v string) *GetUrgentDemandPlanDetailHeaders {
	s.YunUserId = &v
	return s
}

func (s *GetUrgentDemandPlanDetailHeaders) Validate() error {
	return dara.Validate(s)
}
