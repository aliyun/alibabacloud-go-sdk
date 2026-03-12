// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplateId(v int64) *SaveContactTemplateResponseBody
	GetContactTemplateId() *int64
	SetRequestId(v string) *SaveContactTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveContactTemplateResponseBody
	GetSuccess() *bool
}

type SaveContactTemplateResponseBody struct {
	ContactTemplateId *int64  `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveContactTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveContactTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveContactTemplateResponseBody) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *SaveContactTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveContactTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveContactTemplateResponseBody) SetContactTemplateId(v int64) *SaveContactTemplateResponseBody {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveContactTemplateResponseBody) SetRequestId(v string) *SaveContactTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveContactTemplateResponseBody) SetSuccess(v bool) *SaveContactTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *SaveContactTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
