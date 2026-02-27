// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataServiceAppResponseBody
	GetCode() *string
	SetData(v int32) *CreateDataServiceAppResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *CreateDataServiceAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDataServiceAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataServiceAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataServiceAppResponseBody
	GetSuccess() *bool
}

type CreateDataServiceAppResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -809333788831
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataServiceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataServiceAppResponseBody) GetData() *int32 {
	return s.Data
}

func (s *CreateDataServiceAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataServiceAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataServiceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataServiceAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataServiceAppResponseBody) SetCode(v string) *CreateDataServiceAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataServiceAppResponseBody) SetData(v int32) *CreateDataServiceAppResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDataServiceAppResponseBody) SetHttpStatusCode(v int32) *CreateDataServiceAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataServiceAppResponseBody) SetMessage(v string) *CreateDataServiceAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataServiceAppResponseBody) SetRequestId(v string) *CreateDataServiceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataServiceAppResponseBody) SetSuccess(v bool) *CreateDataServiceAppResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataServiceAppResponseBody) Validate() error {
	return dara.Validate(s)
}
