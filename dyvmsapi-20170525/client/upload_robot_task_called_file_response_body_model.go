// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRobotTaskCalledFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadRobotTaskCalledFileResponseBody
	GetCode() *string
	SetData(v string) *UploadRobotTaskCalledFileResponseBody
	GetData() *string
	SetMessage(v string) *UploadRobotTaskCalledFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadRobotTaskCalledFileResponseBody
	GetRequestId() *string
}

type UploadRobotTaskCalledFileResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique ID of the robocall task.
	//
	// You can call the [QueryRobotTaskDetail](~~QueryRobotTaskDetail~~) operation to query the details of the robocall task based on the task ID.
	//
	// example:
	//
	// 10450****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadRobotTaskCalledFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadRobotTaskCalledFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadRobotTaskCalledFileResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadRobotTaskCalledFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadRobotTaskCalledFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadRobotTaskCalledFileResponseBody) SetCode(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetData(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Data = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetMessage(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetRequestId(v string) *UploadRobotTaskCalledFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) Validate() error {
	return dara.Validate(s)
}
