// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetOverviewResponseBody
	GetCode() *int32
	SetData(v string) *GetOverviewResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetOverviewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetOverviewResponseBody
	GetSuccess() *string
}

type GetOverviewResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	//
	// example:
	//
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 52BA6DA6-A702-4362-A32F-DFF79655****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetOverviewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetOverviewResponseBody) GetData() *string {
	return s.Data
}

func (s *GetOverviewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOverviewResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetOverviewResponseBody) SetCode(v int32) *GetOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetOverviewResponseBody) SetData(v string) *GetOverviewResponseBody {
	s.Data = &v
	return s
}

func (s *GetOverviewResponseBody) SetHttpStatusCode(v int32) *GetOverviewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOverviewResponseBody) SetMessage(v string) *GetOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetOverviewResponseBody) SetRequestId(v string) *GetOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOverviewResponseBody) SetSuccess(v string) *GetOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}
