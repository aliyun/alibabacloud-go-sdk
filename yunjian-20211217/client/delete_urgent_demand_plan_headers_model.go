// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteUrgentDemandPlanHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *DeleteUrgentDemandPlanHeaders
	GetYunUserId() *string
}

type DeleteUrgentDemandPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s DeleteUrgentDemandPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandPlanHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteUrgentDemandPlanHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *DeleteUrgentDemandPlanHeaders) SetCommonHeaders(v map[string]*string) *DeleteUrgentDemandPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUrgentDemandPlanHeaders) SetYunUserId(v string) *DeleteUrgentDemandPlanHeaders {
	s.YunUserId = &v
	return s
}

func (s *DeleteUrgentDemandPlanHeaders) Validate() error {
	return dara.Validate(s)
}
