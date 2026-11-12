// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApplicationSSLResponseBody
	GetRequestId() *string
}

type ModifyApplicationSSLResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationSSLResponseBody) SetRequestId(v string) *ModifyApplicationSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
