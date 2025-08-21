// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDomainResourceResponseBody
	GetRequestId() *string
}

type DeleteDomainResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainResourceResponseBody) SetRequestId(v string) *DeleteDomainResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
