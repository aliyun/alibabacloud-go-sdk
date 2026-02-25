// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateParameterSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateParameterSetResponseBody
	GetRequestId() *string
}

type AssociateParameterSetResponseBody struct {
	// example:
	//
	// BF75EF50-955D-5E1F-AB23-A657C2C6D3C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AssociateParameterSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateParameterSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateParameterSetResponseBody) SetRequestId(v string) *AssociateParameterSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateParameterSetResponseBody) Validate() error {
	return dara.Validate(s)
}
