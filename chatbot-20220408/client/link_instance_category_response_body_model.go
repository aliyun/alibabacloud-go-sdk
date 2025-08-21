// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkInstanceCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LinkInstanceCategoryResponseBody
	GetRequestId() *string
}

type LinkInstanceCategoryResponseBody struct {
	// example:
	//
	// D8C96601-E645-1BD7-99F3-04EADAB84E29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LinkInstanceCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LinkInstanceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *LinkInstanceCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LinkInstanceCategoryResponseBody) SetRequestId(v string) *LinkInstanceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *LinkInstanceCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
