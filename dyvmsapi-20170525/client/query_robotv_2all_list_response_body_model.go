// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotv2AllListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRobotv2AllListResponseBody
	GetCode() *string
	SetData(v string) *QueryRobotv2AllListResponseBody
	GetData() *string
	SetMessage(v string) *QueryRobotv2AllListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRobotv2AllListResponseBody
	GetRequestId() *string
}

type QueryRobotv2AllListResponseBody struct {
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
	// The information about the robot communication script, in the JSON format.
	//
	// 	- **id**: the ID of the robot communication script.
	//
	// 	- **robotName**: the name of the robot communication script.
	//
	// 	- **robotType**: the type of the robot communication script.
	//
	// example:
	//
	// {"id":100000007****,"robotName":"Robot","robotType":"CUSTOM"}
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

func (s QueryRobotv2AllListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotv2AllListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotv2AllListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRobotv2AllListResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryRobotv2AllListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRobotv2AllListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRobotv2AllListResponseBody) SetCode(v string) *QueryRobotv2AllListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetData(v string) *QueryRobotv2AllListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetMessage(v string) *QueryRobotv2AllListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetRequestId(v string) *QueryRobotv2AllListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) Validate() error {
	return dara.Validate(s)
}
