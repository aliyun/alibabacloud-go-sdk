// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDomainResourceResponseBody
	GetRequestId() *string
}

type ModifyDomainResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDomainResourceResponseBody) SetRequestId(v string) *ModifyDomainResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDomainResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
