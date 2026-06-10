// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceInspectionResponseBody
	GetCode() *string
	SetData(v interface{}) *CreateInstanceInspectionResponseBody
	GetData() interface{}
	SetMessage(v string) *CreateInstanceInspectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceInspectionResponseBody
	GetRequestId() *string
}

type CreateInstanceInspectionResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// reportId
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateInstanceInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceInspectionResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateInstanceInspectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceInspectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceInspectionResponseBody) SetCode(v string) *CreateInstanceInspectionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetData(v interface{}) *CreateInstanceInspectionResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetMessage(v string) *CreateInstanceInspectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetRequestId(v string) *CreateInstanceInspectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) Validate() error {
	return dara.Validate(s)
}
