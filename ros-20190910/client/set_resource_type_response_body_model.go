// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetResourceTypeResponseBody
	GetRequestId() *string
}

type SetResourceTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SetResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetResourceTypeResponseBody) SetRequestId(v string) *SetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetResourceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
