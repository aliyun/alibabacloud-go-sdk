// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDomainResourceResponseBody
	GetRequestId() *string
}

type CreateDomainResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDomainResourceResponseBody) SetRequestId(v string) *CreateDomainResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
