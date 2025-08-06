// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignProductAccountIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductAccountId(v int64) *AssignProductAccountIdResponseBody
	GetProductAccountId() *int64
	SetRequestId(v string) *AssignProductAccountIdResponseBody
	GetRequestId() *string
}

type AssignProductAccountIdResponseBody struct {
	ProductAccountId *int64  `json:"ProductAccountId,omitempty" xml:"ProductAccountId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignProductAccountIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignProductAccountIdResponseBody) GoString() string {
	return s.String()
}

func (s *AssignProductAccountIdResponseBody) GetProductAccountId() *int64 {
	return s.ProductAccountId
}

func (s *AssignProductAccountIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignProductAccountIdResponseBody) SetProductAccountId(v int64) *AssignProductAccountIdResponseBody {
	s.ProductAccountId = &v
	return s
}

func (s *AssignProductAccountIdResponseBody) SetRequestId(v string) *AssignProductAccountIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignProductAccountIdResponseBody) Validate() error {
	return dara.Validate(s)
}
