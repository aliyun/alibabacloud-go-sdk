// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeployApplicationResponseBody
	GetCode() *string
	SetData(v int64) *DeployApplicationResponseBody
	GetData() *int64
	SetMessage(v string) *DeployApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeployApplicationResponseBody
	GetRequestId() *string
}

type DeployApplicationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the application.
	//
	// example:
	//
	// 123
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeployApplicationResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeployApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeployApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployApplicationResponseBody) SetCode(v string) *DeployApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployApplicationResponseBody) SetData(v int64) *DeployApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *DeployApplicationResponseBody) SetMessage(v string) *DeployApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployApplicationResponseBody) SetRequestId(v string) *DeployApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
