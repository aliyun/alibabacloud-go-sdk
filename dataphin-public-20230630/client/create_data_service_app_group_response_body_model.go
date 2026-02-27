// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataServiceAppGroupResponseBody
	GetCode() *string
	SetData(v int32) *CreateDataServiceAppGroupResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *CreateDataServiceAppGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDataServiceAppGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataServiceAppGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataServiceAppGroupResponseBody
	GetSuccess() *bool
}

type CreateDataServiceAppGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 538893915703
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataServiceAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataServiceAppGroupResponseBody) GetData() *int32 {
	return s.Data
}

func (s *CreateDataServiceAppGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataServiceAppGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataServiceAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataServiceAppGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataServiceAppGroupResponseBody) SetCode(v string) *CreateDataServiceAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataServiceAppGroupResponseBody) SetData(v int32) *CreateDataServiceAppGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDataServiceAppGroupResponseBody) SetHttpStatusCode(v int32) *CreateDataServiceAppGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataServiceAppGroupResponseBody) SetMessage(v string) *CreateDataServiceAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataServiceAppGroupResponseBody) SetRequestId(v string) *CreateDataServiceAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataServiceAppGroupResponseBody) SetSuccess(v bool) *CreateDataServiceAppGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataServiceAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
