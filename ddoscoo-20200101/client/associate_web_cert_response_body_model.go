// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateWebCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateWebCertResponseBody
	GetRequestId() *string
}

type AssociateWebCertResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40F11005-A75C-4644-95F2-52A4E7D43E91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateWebCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateWebCertResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateWebCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateWebCertResponseBody) SetRequestId(v string) *AssociateWebCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateWebCertResponseBody) Validate() error {
	return dara.Validate(s)
}
