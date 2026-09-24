// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestLargeModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *TestLargeModelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *TestLargeModelResponseBody
	GetCode() *string
	SetData(v []*TestLargeModelResponseBodyData) *TestLargeModelResponseBody
	GetData() []*TestLargeModelResponseBodyData
	SetMessage(v string) *TestLargeModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *TestLargeModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestLargeModelResponseBody
	GetSuccess() *bool
}

type TestLargeModelResponseBody struct {
	// The authentication error details.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code. A value of OK indicates that the call was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data list.
	Data []*TestLargeModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2051E18-FF3F-5C08-8D24-6F150D2AF757
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestLargeModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestLargeModelResponseBody) GoString() string {
	return s.String()
}

func (s *TestLargeModelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *TestLargeModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *TestLargeModelResponseBody) GetData() []*TestLargeModelResponseBodyData {
	return s.Data
}

func (s *TestLargeModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TestLargeModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestLargeModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestLargeModelResponseBody) SetAccessDeniedDetail(v string) *TestLargeModelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TestLargeModelResponseBody) SetCode(v string) *TestLargeModelResponseBody {
	s.Code = &v
	return s
}

func (s *TestLargeModelResponseBody) SetData(v []*TestLargeModelResponseBodyData) *TestLargeModelResponseBody {
	s.Data = v
	return s
}

func (s *TestLargeModelResponseBody) SetMessage(v string) *TestLargeModelResponseBody {
	s.Message = &v
	return s
}

func (s *TestLargeModelResponseBody) SetRequestId(v string) *TestLargeModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestLargeModelResponseBody) SetSuccess(v bool) *TestLargeModelResponseBody {
	s.Success = &v
	return s
}

func (s *TestLargeModelResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TestLargeModelResponseBodyData struct {
	// The error message.
	//
	// example:
	//
	// {\\"code\\":\\"InvalidApiKey\\",\\"message\\":\\"Invalid API-key provided.\\",\\"request_id\\":\\"dd14065e-3dd5-90a1-b8ee-d6c80891defe\\"}
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The model version number.
	//
	// example:
	//
	// 1
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// The output result.
	//
	// example:
	//
	// ⌈Hello! I am the AI assistant on the Alibaba Cloud website.\\n\\nYour input seems incomplete. What specific information about Alibaba Cloud would you like to know?\\n\\nFor example:\\n	- What are the core products of Alibaba Cloud (such as ECS, RDS, OSS, etc.)?\\n	- How do I purchase or renew a cloud server?\\n	- What security protection capabilities does Alibaba Cloud offer?\\n	- How do I get technical support or view documentation?\\n\\nFeel free to ask your specific questions, and I will answer them for you!⌋
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestLargeModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestLargeModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestLargeModelResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TestLargeModelResponseBodyData) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *TestLargeModelResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *TestLargeModelResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *TestLargeModelResponseBodyData) SetErrorMsg(v string) *TestLargeModelResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *TestLargeModelResponseBodyData) SetModelVersion(v string) *TestLargeModelResponseBodyData {
	s.ModelVersion = &v
	return s
}

func (s *TestLargeModelResponseBodyData) SetOutput(v string) *TestLargeModelResponseBodyData {
	s.Output = &v
	return s
}

func (s *TestLargeModelResponseBodyData) SetSuccess(v bool) *TestLargeModelResponseBodyData {
	s.Success = &v
	return s
}

func (s *TestLargeModelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
