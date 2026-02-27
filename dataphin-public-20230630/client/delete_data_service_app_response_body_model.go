// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDataServiceAppResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDataServiceAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDataServiceAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDataServiceAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataServiceAppResponseBody
	GetSuccess() *bool
}

type DeleteDataServiceAppResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s DeleteDataServiceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDataServiceAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDataServiceAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDataServiceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataServiceAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataServiceAppResponseBody) SetCode(v string) *DeleteDataServiceAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataServiceAppResponseBody) SetHttpStatusCode(v int32) *DeleteDataServiceAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataServiceAppResponseBody) SetMessage(v string) *DeleteDataServiceAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataServiceAppResponseBody) SetRequestId(v string) *DeleteDataServiceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataServiceAppResponseBody) SetSuccess(v bool) *DeleteDataServiceAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataServiceAppResponseBody) Validate() error {
	return dara.Validate(s)
}
