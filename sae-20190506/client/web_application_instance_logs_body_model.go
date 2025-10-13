// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationInstanceLogsBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationInstanceLogsBody
	GetCode() *int32
	SetData(v *DescribeInstanceLogsOutput) *WebApplicationInstanceLogsBody
	GetData() *DescribeInstanceLogsOutput
	SetMessage(v string) *WebApplicationInstanceLogsBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationInstanceLogsBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationInstanceLogsBody
	GetSuccess() *bool
}

type WebApplicationInstanceLogsBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeInstanceLogsOutput `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebApplicationInstanceLogsBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationInstanceLogsBody) GoString() string {
	return s.String()
}

func (s *WebApplicationInstanceLogsBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationInstanceLogsBody) GetData() *DescribeInstanceLogsOutput {
	return s.Data
}

func (s *WebApplicationInstanceLogsBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationInstanceLogsBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationInstanceLogsBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationInstanceLogsBody) SetCode(v int32) *WebApplicationInstanceLogsBody {
	s.Code = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetData(v *DescribeInstanceLogsOutput) *WebApplicationInstanceLogsBody {
	s.Data = v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetMessage(v string) *WebApplicationInstanceLogsBody {
	s.Message = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetRequestId(v string) *WebApplicationInstanceLogsBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) SetSuccess(v bool) *WebApplicationInstanceLogsBody {
	s.Success = &v
	return s
}

func (s *WebApplicationInstanceLogsBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
