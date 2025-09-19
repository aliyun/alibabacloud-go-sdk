// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertItemsResponseBody
	GetCode() *string
	SetData(v interface{}) *ListAlertItemsResponseBody
	GetData() interface{}
	SetMessage(v string) *ListAlertItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlertItemsResponseBody
	GetRequestId() *string
}

type ListAlertItemsResponseBody struct {
	// example:
	//
	// Success
	Code *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAlertItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertItemsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListAlertItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertItemsResponseBody) SetCode(v string) *ListAlertItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertItemsResponseBody) SetData(v interface{}) *ListAlertItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertItemsResponseBody) SetMessage(v string) *ListAlertItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertItemsResponseBody) SetRequestId(v string) *ListAlertItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
