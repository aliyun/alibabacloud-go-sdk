// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoRenewInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AutoRenewInstanceResponseBody
	GetData() *bool
	SetRequestId(v string) *AutoRenewInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AutoRenewInstanceResponseBody
	GetSuccess() *bool
}

type AutoRenewInstanceResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AutoRenewInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AutoRenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *AutoRenewInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AutoRenewInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AutoRenewInstanceResponseBody) SetData(v bool) *AutoRenewInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *AutoRenewInstanceResponseBody) SetRequestId(v string) *AutoRenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AutoRenewInstanceResponseBody) SetSuccess(v bool) *AutoRenewInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *AutoRenewInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
