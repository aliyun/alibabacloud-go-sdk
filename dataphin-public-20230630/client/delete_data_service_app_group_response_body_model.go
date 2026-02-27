// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDataServiceAppGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDataServiceAppGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDataServiceAppGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDataServiceAppGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataServiceAppGroupResponseBody
	GetSuccess() *bool
}

type DeleteDataServiceAppGroupResponseBody struct {
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

func (s DeleteDataServiceAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceAppGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDataServiceAppGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDataServiceAppGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDataServiceAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataServiceAppGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataServiceAppGroupResponseBody) SetCode(v string) *DeleteDataServiceAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataServiceAppGroupResponseBody) SetHttpStatusCode(v int32) *DeleteDataServiceAppGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataServiceAppGroupResponseBody) SetMessage(v string) *DeleteDataServiceAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataServiceAppGroupResponseBody) SetRequestId(v string) *DeleteDataServiceAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataServiceAppGroupResponseBody) SetSuccess(v bool) *DeleteDataServiceAppGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataServiceAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
