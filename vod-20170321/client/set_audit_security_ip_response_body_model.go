// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAuditSecurityIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAuditSecurityIpResponseBody
	GetRequestId() *string
}

type SetAuditSecurityIpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAuditSecurityIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAuditSecurityIpResponseBody) GoString() string {
	return s.String()
}

func (s *SetAuditSecurityIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAuditSecurityIpResponseBody) SetRequestId(v string) *SetAuditSecurityIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAuditSecurityIpResponseBody) Validate() error {
	return dara.Validate(s)
}
