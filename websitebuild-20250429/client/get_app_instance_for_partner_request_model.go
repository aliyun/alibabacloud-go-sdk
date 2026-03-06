// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppInstanceForPartnerRequest
	GetBizId() *string
}

type GetAppInstanceForPartnerRequest struct {
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetAppInstanceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForPartnerRequest) SetBizId(v string) *GetAppInstanceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
