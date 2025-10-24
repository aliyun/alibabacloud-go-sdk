// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceRegisterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeviceRegisterResponseBody
	GetCode() *string
	SetData(v *DeviceRegisterResponseBodyData) *DeviceRegisterResponseBody
	GetData() *DeviceRegisterResponseBodyData
	SetHttpStatusCode(v int32) *DeviceRegisterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeviceRegisterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeviceRegisterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeviceRegisterResponseBody
	GetSuccess() *bool
}

type DeviceRegisterResponseBody struct {
	// example:
	//
	// success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeviceRegisterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 52548360-B3AA-55EA-893F-48C16470F64A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeviceRegisterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeviceRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeviceRegisterResponseBody) GetData() *DeviceRegisterResponseBodyData {
	return s.Data
}

func (s *DeviceRegisterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeviceRegisterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeviceRegisterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeviceRegisterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeviceRegisterResponseBody) SetCode(v string) *DeviceRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetData(v *DeviceRegisterResponseBodyData) *DeviceRegisterResponseBody {
	s.Data = v
	return s
}

func (s *DeviceRegisterResponseBody) SetHttpStatusCode(v int32) *DeviceRegisterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetMessage(v string) *DeviceRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetRequestId(v string) *DeviceRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetSuccess(v bool) *DeviceRegisterResponseBody {
	s.Success = &v
	return s
}

func (s *DeviceRegisterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeviceRegisterResponseBodyData struct {
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 991fa52b7935aaa33536e05d4f4b5003
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// e2e928e8244f45ab30ec6ba9f9
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// example:
	//
	// 1748312544852
	ResponseTime *string `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	// example:
	//
	// s8wPO/w79jP7sz6OaHkixAje2GmgzmZiCuCZZ+J8w77ICTyqD30lL6rUhnXwwx4MyGF62DRPFnpXtJ6c5mlmt6QdML3FkjGn+i/wR5T6QMkVDW6YRPWsx3jkIWRQ2sDnmVNBtixo2s9w3RJrnddRzVCaR/WeLOfiVLWcrLcJditzO/1YXBZ9vuRKQ4iperfhZvw372N/m8/1qtjJl+FUe2+wxK6RMxr03K7R
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s DeviceRegisterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeviceRegisterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DeviceRegisterResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeviceRegisterResponseBodyData) GetNonce() *string {
	return s.Nonce
}

func (s *DeviceRegisterResponseBodyData) GetResponseTime() *string {
	return s.ResponseTime
}

func (s *DeviceRegisterResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *DeviceRegisterResponseBodyData) SetAppId(v string) *DeviceRegisterResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetDeviceName(v string) *DeviceRegisterResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetNonce(v string) *DeviceRegisterResponseBodyData {
	s.Nonce = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetResponseTime(v string) *DeviceRegisterResponseBodyData {
	s.ResponseTime = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetSignature(v string) *DeviceRegisterResponseBodyData {
	s.Signature = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
