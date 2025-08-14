// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteByPassShuntEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteByPassShuntEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DeleteByPassShuntEventResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DeleteByPassShuntEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteByPassShuntEventResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteByPassShuntEventResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *DeleteByPassShuntEventResponseBody
	GetSuccess() *bool
}

type DeleteByPassShuntEventResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Whether it was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteByPassShuntEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteByPassShuntEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteByPassShuntEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteByPassShuntEventResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteByPassShuntEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteByPassShuntEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteByPassShuntEventResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteByPassShuntEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteByPassShuntEventResponseBody) SetCode(v string) *DeleteByPassShuntEventResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteByPassShuntEventResponseBody) SetHttpStatusCode(v string) *DeleteByPassShuntEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteByPassShuntEventResponseBody) SetMessage(v string) *DeleteByPassShuntEventResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteByPassShuntEventResponseBody) SetRequestId(v string) *DeleteByPassShuntEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteByPassShuntEventResponseBody) SetResultObject(v bool) *DeleteByPassShuntEventResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteByPassShuntEventResponseBody) SetSuccess(v bool) *DeleteByPassShuntEventResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteByPassShuntEventResponseBody) Validate() error {
	return dara.Validate(s)
}
