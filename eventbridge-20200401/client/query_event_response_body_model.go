// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEventResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryEventResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEventResponseBody
	GetSuccess() *bool
}

type QueryEventResponseBody struct {
	// The response code. Valid values:
	//
	// 200: The request was successful.
	//
	// Other values indicate that the request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The content of the event.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 580A938B-6107-586C-8EC7-F22EEBEDA9E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEventResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEventResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEventResponseBody) SetCode(v string) *QueryEventResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventResponseBody) SetData(v map[string]interface{}) *QueryEventResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventResponseBody) SetMessage(v string) *QueryEventResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEventResponseBody) SetRequestId(v string) *QueryEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventResponseBody) SetSuccess(v bool) *QueryEventResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEventResponseBody) Validate() error {
	return dara.Validate(s)
}
