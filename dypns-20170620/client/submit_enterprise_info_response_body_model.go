// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitEnterpriseInfoResponseBody
	GetCode() *string
	SetData(v bool) *SubmitEnterpriseInfoResponseBody
	GetData() *bool
	SetMessage(v string) *SubmitEnterpriseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitEnterpriseInfoResponseBody
	GetRequestId() *string
}

type SubmitEnterpriseInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitEnterpriseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitEnterpriseInfoResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitEnterpriseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitEnterpriseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitEnterpriseInfoResponseBody) SetCode(v string) *SubmitEnterpriseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitEnterpriseInfoResponseBody) SetData(v bool) *SubmitEnterpriseInfoResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitEnterpriseInfoResponseBody) SetMessage(v string) *SubmitEnterpriseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitEnterpriseInfoResponseBody) SetRequestId(v string) *SubmitEnterpriseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitEnterpriseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
