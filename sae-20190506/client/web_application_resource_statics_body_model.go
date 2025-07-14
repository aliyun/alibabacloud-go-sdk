// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationResourceStaticsBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebApplicationResourceStaticsBody
	GetCode() *int32
	SetData(v *DescribeWebAppStaticsOutput) *WebApplicationResourceStaticsBody
	GetData() *DescribeWebAppStaticsOutput
	SetMessage(v string) *WebApplicationResourceStaticsBody
	GetMessage() *string
	SetRequestId(v string) *WebApplicationResourceStaticsBody
	GetRequestId() *string
	SetSuccess(v bool) *WebApplicationResourceStaticsBody
	GetSuccess() *bool
}

type WebApplicationResourceStaticsBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeWebAppStaticsOutput `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebApplicationResourceStaticsBody) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationResourceStaticsBody) GoString() string {
	return s.String()
}

func (s *WebApplicationResourceStaticsBody) GetCode() *int32 {
	return s.Code
}

func (s *WebApplicationResourceStaticsBody) GetData() *DescribeWebAppStaticsOutput {
	return s.Data
}

func (s *WebApplicationResourceStaticsBody) GetMessage() *string {
	return s.Message
}

func (s *WebApplicationResourceStaticsBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebApplicationResourceStaticsBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebApplicationResourceStaticsBody) SetCode(v int32) *WebApplicationResourceStaticsBody {
	s.Code = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetData(v *DescribeWebAppStaticsOutput) *WebApplicationResourceStaticsBody {
	s.Data = v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetMessage(v string) *WebApplicationResourceStaticsBody {
	s.Message = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetRequestId(v string) *WebApplicationResourceStaticsBody {
	s.RequestId = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) SetSuccess(v bool) *WebApplicationResourceStaticsBody {
	s.Success = &v
	return s
}

func (s *WebApplicationResourceStaticsBody) Validate() error {
	return dara.Validate(s)
}
