// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveInstanceWhiteListResponseBody
	GetCode() *string
	SetData(v interface{}) *RemoveInstanceWhiteListResponseBody
	GetData() interface{}
	SetMessage(v string) *RemoveInstanceWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveInstanceWhiteListResponseBody
	GetRequestId() *string
	SetStatusCode(v string) *RemoveInstanceWhiteListResponseBody
	GetStatusCode() *string
	SetSuccess(v bool) *RemoveInstanceWhiteListResponseBody
	GetSuccess() *bool
}

type RemoveInstanceWhiteListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// can not find resource: instanceId:amqp-cn-xxx, whiteListItemId:420, whiteListType:1
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 92385FD2-624A-48C9-8FB5-753F2AFA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveInstanceWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstanceWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveInstanceWhiteListResponseBody) GetData() interface{} {
	return s.Data
}

func (s *RemoveInstanceWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveInstanceWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveInstanceWhiteListResponseBody) GetStatusCode() *string {
	return s.StatusCode
}

func (s *RemoveInstanceWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveInstanceWhiteListResponseBody) SetCode(v string) *RemoveInstanceWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveInstanceWhiteListResponseBody) SetData(v interface{}) *RemoveInstanceWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *RemoveInstanceWhiteListResponseBody) SetMessage(v string) *RemoveInstanceWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveInstanceWhiteListResponseBody) SetRequestId(v string) *RemoveInstanceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstanceWhiteListResponseBody) SetStatusCode(v string) *RemoveInstanceWhiteListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstanceWhiteListResponseBody) SetSuccess(v bool) *RemoveInstanceWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveInstanceWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
