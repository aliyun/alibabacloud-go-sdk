// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationInstancesBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWebApplicationInstancesBody
	GetCode() *int32
	SetData(v *ListWebInstancesOutput) *ListWebApplicationInstancesBody
	GetData() *ListWebInstancesOutput
	SetMessage(v string) *ListWebApplicationInstancesBody
	GetMessage() *string
	SetRequestId(v string) *ListWebApplicationInstancesBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWebApplicationInstancesBody
	GetSuccess() *bool
}

type ListWebApplicationInstancesBody struct {
	Code      *int32                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListWebInstancesOutput `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWebApplicationInstancesBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationInstancesBody) GoString() string {
	return s.String()
}

func (s *ListWebApplicationInstancesBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWebApplicationInstancesBody) GetData() *ListWebInstancesOutput {
	return s.Data
}

func (s *ListWebApplicationInstancesBody) GetMessage() *string {
	return s.Message
}

func (s *ListWebApplicationInstancesBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebApplicationInstancesBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWebApplicationInstancesBody) SetCode(v int32) *ListWebApplicationInstancesBody {
	s.Code = &v
	return s
}

func (s *ListWebApplicationInstancesBody) SetData(v *ListWebInstancesOutput) *ListWebApplicationInstancesBody {
	s.Data = v
	return s
}

func (s *ListWebApplicationInstancesBody) SetMessage(v string) *ListWebApplicationInstancesBody {
	s.Message = &v
	return s
}

func (s *ListWebApplicationInstancesBody) SetRequestId(v string) *ListWebApplicationInstancesBody {
	s.RequestId = &v
	return s
}

func (s *ListWebApplicationInstancesBody) SetSuccess(v bool) *ListWebApplicationInstancesBody {
	s.Success = &v
	return s
}

func (s *ListWebApplicationInstancesBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
