// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkDomainResponseBody
	GetRequestId() *string
}

type DeleteNetworkDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 66B9D942-E3C8-5068-A479-5A7B7BF3DE35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkDomainResponseBody) SetRequestId(v string) *DeleteNetworkDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
