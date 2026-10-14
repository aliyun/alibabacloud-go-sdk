// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotfixStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryHotfixStatusResponseBody
	GetCode() *string
	SetData(v *QueryHotfixStatusResponseBodyData) *QueryHotfixStatusResponseBody
	GetData() *QueryHotfixStatusResponseBodyData
	SetMessage(v string) *QueryHotfixStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryHotfixStatusResponseBody
	GetRequestId() *string
}

type QueryHotfixStatusResponseBody struct {
	// The status code.
	//
	// - `code == Success`: The authorization is successful.
	//
	// - Other status codes: The authorization failed. If the authorization fails, check the `message` field for the detailed fault message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *QueryHotfixStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryHotfixStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHotfixStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotfixStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryHotfixStatusResponseBody) GetData() *QueryHotfixStatusResponseBodyData {
	return s.Data
}

func (s *QueryHotfixStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryHotfixStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHotfixStatusResponseBody) SetCode(v string) *QueryHotfixStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotfixStatusResponseBody) SetData(v *QueryHotfixStatusResponseBodyData) *QueryHotfixStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryHotfixStatusResponseBody) SetMessage(v string) *QueryHotfixStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotfixStatusResponseBody) SetRequestId(v string) *QueryHotfixStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotfixStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHotfixStatusResponseBodyData struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// hotfix-da3dcc0b-5019-4b59-b12f-2e500da9083c
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryHotfixStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryHotfixStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryHotfixStatusResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryHotfixStatusResponseBodyData) SetTaskId(v string) *QueryHotfixStatusResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryHotfixStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
