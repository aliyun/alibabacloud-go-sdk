// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConfigByNameResponseBody
	GetCode() *string
	SetData(v string) *GetConfigByNameResponseBody
	GetData() *string
	SetMessage(v string) *GetConfigByNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConfigByNameResponseBody
	GetRequestId() *string
}

type GetConfigByNameResponseBody struct {
	// The status code.
	//
	// - `code == Success` indicates that the authorization is successful.
	//
	// - Other status codes indicate that the authorization has failed. Check the `message` field for the detailed error message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	//
	// example:
	//
	// 1
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error information.
	//
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

func (s GetConfigByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigByNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConfigByNameResponseBody) GetData() *string {
	return s.Data
}

func (s *GetConfigByNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConfigByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigByNameResponseBody) SetCode(v string) *GetConfigByNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetConfigByNameResponseBody) SetData(v string) *GetConfigByNameResponseBody {
	s.Data = &v
	return s
}

func (s *GetConfigByNameResponseBody) SetMessage(v string) *GetConfigByNameResponseBody {
	s.Message = &v
	return s
}

func (s *GetConfigByNameResponseBody) SetRequestId(v string) *GetConfigByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigByNameResponseBody) Validate() error {
	return dara.Validate(s)
}
