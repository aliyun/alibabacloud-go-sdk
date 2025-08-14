// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartOrStopByPassShuntEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartOrStopByPassShuntEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *StartOrStopByPassShuntEventResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *StartOrStopByPassShuntEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartOrStopByPassShuntEventResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *StartOrStopByPassShuntEventResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *StartOrStopByPassShuntEventResponseBody
	GetSuccess() *bool
}

type StartOrStopByPassShuntEventResponseBody struct {
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
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartOrStopByPassShuntEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartOrStopByPassShuntEventResponseBody) GoString() string {
	return s.String()
}

func (s *StartOrStopByPassShuntEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartOrStopByPassShuntEventResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *StartOrStopByPassShuntEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartOrStopByPassShuntEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartOrStopByPassShuntEventResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *StartOrStopByPassShuntEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartOrStopByPassShuntEventResponseBody) SetCode(v string) *StartOrStopByPassShuntEventResponseBody {
	s.Code = &v
	return s
}

func (s *StartOrStopByPassShuntEventResponseBody) SetHttpStatusCode(v string) *StartOrStopByPassShuntEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartOrStopByPassShuntEventResponseBody) SetMessage(v string) *StartOrStopByPassShuntEventResponseBody {
	s.Message = &v
	return s
}

func (s *StartOrStopByPassShuntEventResponseBody) SetRequestId(v string) *StartOrStopByPassShuntEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartOrStopByPassShuntEventResponseBody) SetResultObject(v bool) *StartOrStopByPassShuntEventResponseBody {
	s.ResultObject = &v
	return s
}

func (s *StartOrStopByPassShuntEventResponseBody) SetSuccess(v bool) *StartOrStopByPassShuntEventResponseBody {
	s.Success = &v
	return s
}

func (s *StartOrStopByPassShuntEventResponseBody) Validate() error {
	return dara.Validate(s)
}
