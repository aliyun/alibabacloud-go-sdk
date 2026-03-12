// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactTemplateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveContactTemplateCredentialResponseBody
	GetRequestId() *string
}

type SaveContactTemplateCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveContactTemplateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveContactTemplateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *SaveContactTemplateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveContactTemplateCredentialResponseBody) SetRequestId(v string) *SaveContactTemplateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveContactTemplateCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
