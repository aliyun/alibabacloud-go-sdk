// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachNetworkInterfaceResponseBody
	GetRequestId() *string
}

type DetachNetworkInterfaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BE1B8ECF-9507-4C78-B197-5DE9FED344AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachNetworkInterfaceResponseBody) SetRequestId(v string) *DetachNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachNetworkInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
