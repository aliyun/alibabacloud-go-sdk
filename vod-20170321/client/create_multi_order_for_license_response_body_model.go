// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderForLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateMultiOrderForLicenseResponseBody
	GetData() *string
	SetRequestId(v string) *CreateMultiOrderForLicenseResponseBody
	GetRequestId() *string
}

type CreateMultiOrderForLicenseResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiOrderForLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderForLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderForLicenseResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateMultiOrderForLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultiOrderForLicenseResponseBody) SetData(v string) *CreateMultiOrderForLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateMultiOrderForLicenseResponseBody) SetRequestId(v string) *CreateMultiOrderForLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiOrderForLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
