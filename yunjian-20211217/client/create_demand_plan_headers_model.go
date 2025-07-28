// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDemandPlanHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *CreateDemandPlanHeaders
	GetYunUserId() *string
}

type CreateDemandPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s CreateDemandPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanHeaders) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDemandPlanHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *CreateDemandPlanHeaders) SetCommonHeaders(v map[string]*string) *CreateDemandPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDemandPlanHeaders) SetYunUserId(v string) *CreateDemandPlanHeaders {
	s.YunUserId = &v
	return s
}

func (s *CreateDemandPlanHeaders) Validate() error {
	return dara.Validate(s)
}
