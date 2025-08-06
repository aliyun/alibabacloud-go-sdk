// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateOrderForLicenseResponseBody
	GetData() *string
	SetRequestId(v string) *CreateOrderForLicenseResponseBody
	GetRequestId() *string
}

type CreateOrderForLicenseResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrderForLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderForLicenseResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateOrderForLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderForLicenseResponseBody) SetData(v string) *CreateOrderForLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOrderForLicenseResponseBody) SetRequestId(v string) *CreateOrderForLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderForLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
