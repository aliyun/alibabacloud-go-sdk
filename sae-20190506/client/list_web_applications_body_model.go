// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationsBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWebApplicationsBody
	GetCode() *int32
	SetData(v *ListWebApplicationsOutput) *ListWebApplicationsBody
	GetData() *ListWebApplicationsOutput
	SetMessage(v string) *ListWebApplicationsBody
	GetMessage() *string
	SetRequestId(v string) *ListWebApplicationsBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWebApplicationsBody
	GetSuccess() *bool
}

type ListWebApplicationsBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListWebApplicationsOutput `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWebApplicationsBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationsBody) GoString() string {
	return s.String()
}

func (s *ListWebApplicationsBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWebApplicationsBody) GetData() *ListWebApplicationsOutput {
	return s.Data
}

func (s *ListWebApplicationsBody) GetMessage() *string {
	return s.Message
}

func (s *ListWebApplicationsBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebApplicationsBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWebApplicationsBody) SetCode(v int32) *ListWebApplicationsBody {
	s.Code = &v
	return s
}

func (s *ListWebApplicationsBody) SetData(v *ListWebApplicationsOutput) *ListWebApplicationsBody {
	s.Data = v
	return s
}

func (s *ListWebApplicationsBody) SetMessage(v string) *ListWebApplicationsBody {
	s.Message = &v
	return s
}

func (s *ListWebApplicationsBody) SetRequestId(v string) *ListWebApplicationsBody {
	s.RequestId = &v
	return s
}

func (s *ListWebApplicationsBody) SetSuccess(v bool) *ListWebApplicationsBody {
	s.Success = &v
	return s
}

func (s *ListWebApplicationsBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
