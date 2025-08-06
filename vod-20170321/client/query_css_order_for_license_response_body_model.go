// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCssOrderForLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *QueryCssOrderForLicenseResponseBody
	GetData() *string
	SetRequestId(v string) *QueryCssOrderForLicenseResponseBody
	GetRequestId() *string
}

type QueryCssOrderForLicenseResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCssOrderForLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCssOrderForLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCssOrderForLicenseResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryCssOrderForLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCssOrderForLicenseResponseBody) SetData(v string) *QueryCssOrderForLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCssOrderForLicenseResponseBody) SetRequestId(v string) *QueryCssOrderForLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCssOrderForLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
