// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebCustomDomainBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWebCustomDomainBody
	GetCode() *int32
	SetData(v *ListWebCustomDomainOutput) *ListWebCustomDomainBody
	GetData() *ListWebCustomDomainOutput
	SetMessage(v string) *ListWebCustomDomainBody
	GetMessage() *string
	SetRequestId(v string) *ListWebCustomDomainBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWebCustomDomainBody
	GetSuccess() *bool
}

type ListWebCustomDomainBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListWebCustomDomainOutput `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWebCustomDomainBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebCustomDomainBody) GoString() string {
	return s.String()
}

func (s *ListWebCustomDomainBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWebCustomDomainBody) GetData() *ListWebCustomDomainOutput {
	return s.Data
}

func (s *ListWebCustomDomainBody) GetMessage() *string {
	return s.Message
}

func (s *ListWebCustomDomainBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebCustomDomainBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWebCustomDomainBody) SetCode(v int32) *ListWebCustomDomainBody {
	s.Code = &v
	return s
}

func (s *ListWebCustomDomainBody) SetData(v *ListWebCustomDomainOutput) *ListWebCustomDomainBody {
	s.Data = v
	return s
}

func (s *ListWebCustomDomainBody) SetMessage(v string) *ListWebCustomDomainBody {
	s.Message = &v
	return s
}

func (s *ListWebCustomDomainBody) SetRequestId(v string) *ListWebCustomDomainBody {
	s.RequestId = &v
	return s
}

func (s *ListWebCustomDomainBody) SetSuccess(v bool) *ListWebCustomDomainBody {
	s.Success = &v
	return s
}

func (s *ListWebCustomDomainBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
