// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceWhiteListResponseBody
	GetCode() *string
	SetData(v interface{}) *ListInstanceWhiteListResponseBody
	GetData() interface{}
	SetMessage(v string) *ListInstanceWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceWhiteListResponseBody
	GetRequestId() *string
	SetStatusCode(v string) *ListInstanceWhiteListResponseBody
	GetStatusCode() *string
	SetSuccess(v bool) *ListInstanceWhiteListResponseBody
	GetSuccess() *bool
}

type ListInstanceWhiteListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// "Data": [
	//
	//     {
	//
	//       "id": 454,
	//
	//       "value": "10.1.2.1/30"
	//
	//     }
	//
	//   ]
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxx failed,xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6DC68EC9-0E76-5435-B8C0-FF9492B4****
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

func (s ListInstanceWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceWhiteListResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListInstanceWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceWhiteListResponseBody) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListInstanceWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceWhiteListResponseBody) SetCode(v string) *ListInstanceWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceWhiteListResponseBody) SetData(v interface{}) *ListInstanceWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceWhiteListResponseBody) SetMessage(v string) *ListInstanceWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceWhiteListResponseBody) SetRequestId(v string) *ListInstanceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceWhiteListResponseBody) SetStatusCode(v string) *ListInstanceWhiteListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceWhiteListResponseBody) SetSuccess(v bool) *ListInstanceWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
