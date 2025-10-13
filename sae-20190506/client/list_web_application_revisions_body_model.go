// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationRevisionsBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWebApplicationRevisionsBody
	GetCode() *int32
	SetData(v *ListWebApplicationRevisionsOutput) *ListWebApplicationRevisionsBody
	GetData() *ListWebApplicationRevisionsOutput
	SetMessage(v string) *ListWebApplicationRevisionsBody
	GetMessage() *string
	SetRequestId(v string) *ListWebApplicationRevisionsBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWebApplicationRevisionsBody
	GetSuccess() *bool
}

type ListWebApplicationRevisionsBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListWebApplicationRevisionsOutput `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWebApplicationRevisionsBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationRevisionsBody) GoString() string {
	return s.String()
}

func (s *ListWebApplicationRevisionsBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWebApplicationRevisionsBody) GetData() *ListWebApplicationRevisionsOutput {
	return s.Data
}

func (s *ListWebApplicationRevisionsBody) GetMessage() *string {
	return s.Message
}

func (s *ListWebApplicationRevisionsBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebApplicationRevisionsBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWebApplicationRevisionsBody) SetCode(v int32) *ListWebApplicationRevisionsBody {
	s.Code = &v
	return s
}

func (s *ListWebApplicationRevisionsBody) SetData(v *ListWebApplicationRevisionsOutput) *ListWebApplicationRevisionsBody {
	s.Data = v
	return s
}

func (s *ListWebApplicationRevisionsBody) SetMessage(v string) *ListWebApplicationRevisionsBody {
	s.Message = &v
	return s
}

func (s *ListWebApplicationRevisionsBody) SetRequestId(v string) *ListWebApplicationRevisionsBody {
	s.RequestId = &v
	return s
}

func (s *ListWebApplicationRevisionsBody) SetSuccess(v bool) *ListWebApplicationRevisionsBody {
	s.Success = &v
	return s
}

func (s *ListWebApplicationRevisionsBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
