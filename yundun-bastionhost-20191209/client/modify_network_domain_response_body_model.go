// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNetworkDomainResponseBody
	GetRequestId() *string
}

type ModifyNetworkDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ED49CD1E-3510-5E5C-9133-E2067B656501
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNetworkDomainResponseBody) SetRequestId(v string) *ModifyNetworkDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNetworkDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
