// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRobotInfoListResponseBody
	GetCode() *string
	SetData(v string) *QueryRobotInfoListResponseBody
	GetData() *string
	SetMessage(v string) *QueryRobotInfoListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRobotInfoListResponseBody
	GetRequestId() *string
}

type QueryRobotInfoListResponseBody struct {
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
	// The basic information about the robot, in the JSON format. The basic information contains the following parameters:
	//
	// 	- **id**: the robot ID.
	//
	// 	- **robotName**: the robot name.
	//
	// 	- **robotType**: the robot type.
	//
	// 	- **auditStatus**: the review state.
	//
	// 	- **gmtCreate**: the time when the task was created.
	//
	// 	- **gmtModified**: the time when the task was modified.
	//
	// 	- **partnerId**: the partner ID.
	//
	// 	- **asrId**: the ID of the automatic speech recognition (ASR) model.
	//
	// 	- **asrType**: the ASR type. Valid values: **Public*	- and **Private**.
	//
	// 	- **remark**: the additional information.
	//
	// example:
	//
	// {["id":1000010920004, "gmtModified":"Thu Mar 21 15:38:55 CST 2019", "auditStatus":"AUDITPASS","gmtCreate":"Thu Mar 21 12:00:51 CST 2019","remark":"tset","partnerId":100000022670001,"asrId":"a9a1d69081fd4266ad788346bf5e1b6c","robotType":"CUSTOM","asrType":"1","robotName":"Collection scenario"},{"id":1000010920003,"gmtModified":"Thu Mar 21 11:51:10 CST 2019","auditStatus":"AUDITPASS","gmtCreate":"Thu Mar 21 11:44:57 CST 2019","remark":"test","partnerId":100000022670001,"asrId":"a9a1d69081fd4266ad788346bf5e1b6c","robotType":"CUSTOM","asrType":"1","robotName":"Collection scenario"]}
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
	// F59AF338-655D-48E8-9471-5EB07692B1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRobotInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRobotInfoListResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryRobotInfoListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRobotInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRobotInfoListResponseBody) SetCode(v string) *QueryRobotInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetData(v string) *QueryRobotInfoListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetMessage(v string) *QueryRobotInfoListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetRequestId(v string) *QueryRobotInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) Validate() error {
	return dara.Validate(s)
}
