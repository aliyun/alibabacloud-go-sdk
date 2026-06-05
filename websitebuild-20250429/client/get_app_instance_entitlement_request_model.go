// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceEntitlementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppInstanceEntitlementRequest
	GetBizId() *string
}

type GetAppInstanceEntitlementRequest struct {
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetAppInstanceEntitlementRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceEntitlementRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceEntitlementRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceEntitlementRequest) SetBizId(v string) *GetAppInstanceEntitlementRequest {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceEntitlementRequest) Validate() error {
	return dara.Validate(s)
}
