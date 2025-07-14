// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebCustomDomainBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebCustomDomainBody
	GetCode() *int32
	SetData(v *WebCustomDomain) *WebCustomDomainBody
	GetData() *WebCustomDomain
	SetMessage(v string) *WebCustomDomainBody
	GetMessage() *string
	SetRequestId(v string) *WebCustomDomainBody
	GetRequestId() *string
	SetSuccess(v bool) *WebCustomDomainBody
	GetSuccess() *bool
}

type WebCustomDomainBody struct {
	Code      *int32           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *WebCustomDomain `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebCustomDomainBody) String() string {
	return dara.Prettify(s)
}

func (s WebCustomDomainBody) GoString() string {
	return s.String()
}

func (s *WebCustomDomainBody) GetCode() *int32 {
	return s.Code
}

func (s *WebCustomDomainBody) GetData() *WebCustomDomain {
	return s.Data
}

func (s *WebCustomDomainBody) GetMessage() *string {
	return s.Message
}

func (s *WebCustomDomainBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebCustomDomainBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebCustomDomainBody) SetCode(v int32) *WebCustomDomainBody {
	s.Code = &v
	return s
}

func (s *WebCustomDomainBody) SetData(v *WebCustomDomain) *WebCustomDomainBody {
	s.Data = v
	return s
}

func (s *WebCustomDomainBody) SetMessage(v string) *WebCustomDomainBody {
	s.Message = &v
	return s
}

func (s *WebCustomDomainBody) SetRequestId(v string) *WebCustomDomainBody {
	s.RequestId = &v
	return s
}

func (s *WebCustomDomainBody) SetSuccess(v bool) *WebCustomDomainBody {
	s.Success = &v
	return s
}

func (s *WebCustomDomainBody) Validate() error {
	return dara.Validate(s)
}
