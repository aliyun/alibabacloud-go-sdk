// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenGovernanceServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenGovernanceServiceResponseBody
	GetRequestId() *string
}

type OpenGovernanceServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 019FEF9F-6442-16F6-B041-9013B97987AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenGovernanceServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenGovernanceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenGovernanceServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenGovernanceServiceResponseBody) SetRequestId(v string) *OpenGovernanceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenGovernanceServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
