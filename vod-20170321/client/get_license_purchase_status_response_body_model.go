// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicensePurchaseStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLicensePurchaseStatusResponseBody
	GetCode() *string
	SetData(v map[string]*DataValue) *GetLicensePurchaseStatusResponseBody
	GetData() map[string]*DataValue
	SetHttpStatusCode(v int32) *GetLicensePurchaseStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLicensePurchaseStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLicensePurchaseStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLicensePurchaseStatusResponseBody
	GetSuccess() *bool
}

type GetLicensePurchaseStatusResponseBody struct {
	Code           *string               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           map[string]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLicensePurchaseStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLicensePurchaseStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicensePurchaseStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLicensePurchaseStatusResponseBody) GetData() map[string]*DataValue {
	return s.Data
}

func (s *GetLicensePurchaseStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLicensePurchaseStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLicensePurchaseStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLicensePurchaseStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLicensePurchaseStatusResponseBody) SetCode(v string) *GetLicensePurchaseStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetLicensePurchaseStatusResponseBody) SetData(v map[string]*DataValue) *GetLicensePurchaseStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetLicensePurchaseStatusResponseBody) SetHttpStatusCode(v int32) *GetLicensePurchaseStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLicensePurchaseStatusResponseBody) SetMessage(v string) *GetLicensePurchaseStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetLicensePurchaseStatusResponseBody) SetRequestId(v string) *GetLicensePurchaseStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLicensePurchaseStatusResponseBody) SetSuccess(v bool) *GetLicensePurchaseStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetLicensePurchaseStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
