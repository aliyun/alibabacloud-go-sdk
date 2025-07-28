// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUrgentDemandPlanListHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *GetUrgentDemandPlanListHeaders
	GetYunUserId() *string
}

type GetUrgentDemandPlanListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s GetUrgentDemandPlanListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanListHeaders) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUrgentDemandPlanListHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *GetUrgentDemandPlanListHeaders) SetCommonHeaders(v map[string]*string) *GetUrgentDemandPlanListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUrgentDemandPlanListHeaders) SetYunUserId(v string) *GetUrgentDemandPlanListHeaders {
	s.YunUserId = &v
	return s
}

func (s *GetUrgentDemandPlanListHeaders) Validate() error {
	return dara.Validate(s)
}
