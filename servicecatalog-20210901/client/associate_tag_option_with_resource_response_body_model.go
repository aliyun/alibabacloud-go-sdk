// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTagOptionWithResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateTagOptionWithResourceResponseBody
	GetRequestId() *string
}

type AssociateTagOptionWithResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DC833558-AFF4-5935-9AB6-8A12EE7D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateTagOptionWithResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateTagOptionWithResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateTagOptionWithResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateTagOptionWithResourceResponseBody) SetRequestId(v string) *AssociateTagOptionWithResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateTagOptionWithResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
