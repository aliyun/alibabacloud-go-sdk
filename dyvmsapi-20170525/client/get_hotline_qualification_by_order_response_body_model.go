// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineQualificationByOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineQualificationByOrderResponseBody
	GetCode() *string
	SetData(v *GetHotlineQualificationByOrderResponseBodyData) *GetHotlineQualificationByOrderResponseBody
	GetData() *GetHotlineQualificationByOrderResponseBodyData
	SetMessage(v string) *GetHotlineQualificationByOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineQualificationByOrderResponseBody
	GetRequestId() *string
}

type GetHotlineQualificationByOrderResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetHotlineQualificationByOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHotlineQualificationByOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineQualificationByOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineQualificationByOrderResponseBody) GetData() *GetHotlineQualificationByOrderResponseBodyData {
	return s.Data
}

func (s *GetHotlineQualificationByOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineQualificationByOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineQualificationByOrderResponseBody) SetCode(v string) *GetHotlineQualificationByOrderResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBody) SetData(v *GetHotlineQualificationByOrderResponseBodyData) *GetHotlineQualificationByOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBody) SetMessage(v string) *GetHotlineQualificationByOrderResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBody) SetRequestId(v string) *GetHotlineQualificationByOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHotlineQualificationByOrderResponseBodyData struct {
	// The ID of the qualification application ticket.
	//
	// example:
	//
	// 22456****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The qualification ID.
	//
	// example:
	//
	// 1478*****
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The qualification state. Valid values:
	//
	// 	- **NORMAL**
	//
	// 	- **OTHER**
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetHotlineQualificationByOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineQualificationByOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *GetHotlineQualificationByOrderResponseBodyData) GetQualificationId() *string {
	return s.QualificationId
}

func (s *GetHotlineQualificationByOrderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetHotlineQualificationByOrderResponseBodyData) SetOrderId(v string) *GetHotlineQualificationByOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBodyData) SetQualificationId(v string) *GetHotlineQualificationByOrderResponseBodyData {
	s.QualificationId = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBodyData) SetStatus(v string) *GetHotlineQualificationByOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
