// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompanyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCompanyResponseBody
	GetRequestId() *string
}

type DeleteCompanyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 28627F67-2F87-55E6-B9C6-CE32FEC91315
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCompanyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCompanyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCompanyResponseBody) SetRequestId(v string) *DeleteCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCompanyResponseBody) Validate() error {
	return dara.Validate(s)
}
