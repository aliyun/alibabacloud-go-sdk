// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisAssociateTagOptionFromResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisAssociateTagOptionFromResourceResponseBody
	GetRequestId() *string
}

type DisAssociateTagOptionFromResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7E23D1F3-4333-587B-909C-39EA4626****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisAssociateTagOptionFromResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisAssociateTagOptionFromResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DisAssociateTagOptionFromResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisAssociateTagOptionFromResourceResponseBody) SetRequestId(v string) *DisAssociateTagOptionFromResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisAssociateTagOptionFromResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
