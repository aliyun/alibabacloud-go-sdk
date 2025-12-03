// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainExtensionAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDomainExtensionAttributeResponseBody
	GetRequestId() *string
}

type SetDomainExtensionAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 149A2470-F010-4437-BF68-343D5099C19D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainExtensionAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDomainExtensionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainExtensionAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDomainExtensionAttributeResponseBody) SetRequestId(v string) *SetDomainExtensionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainExtensionAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
