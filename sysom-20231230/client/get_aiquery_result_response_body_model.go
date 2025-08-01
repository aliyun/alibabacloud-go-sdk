// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIQueryResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAIQueryResultResponseBody
	GetCode() *string
	SetData(v string) *GetAIQueryResultResponseBody
	GetData() *string
	SetMessage(v string) *GetAIQueryResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAIQueryResultResponseBody
	GetRequestId() *string
}

type GetAIQueryResultResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {\\"task_id\\": \\"y4ba8uRV\\"}
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAIQueryResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIQueryResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAIQueryResultResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAIQueryResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAIQueryResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIQueryResultResponseBody) SetCode(v string) *GetAIQueryResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAIQueryResultResponseBody) SetData(v string) *GetAIQueryResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetAIQueryResultResponseBody) SetMessage(v string) *GetAIQueryResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAIQueryResultResponseBody) SetRequestId(v string) *GetAIQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIQueryResultResponseBody) Validate() error {
	return dara.Validate(s)
}
