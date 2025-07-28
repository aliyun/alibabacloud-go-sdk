// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitUrgentDemandPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SubmitUrgentDemandPlanHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *SubmitUrgentDemandPlanHeaders
	GetYunUserId() *string
}

type SubmitUrgentDemandPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s SubmitUrgentDemandPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s SubmitUrgentDemandPlanHeaders) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SubmitUrgentDemandPlanHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *SubmitUrgentDemandPlanHeaders) SetCommonHeaders(v map[string]*string) *SubmitUrgentDemandPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitUrgentDemandPlanHeaders) SetYunUserId(v string) *SubmitUrgentDemandPlanHeaders {
	s.YunUserId = &v
	return s
}

func (s *SubmitUrgentDemandPlanHeaders) Validate() error {
	return dara.Validate(s)
}
