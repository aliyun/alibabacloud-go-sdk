// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenResponseBody
	GetCode() *string
	SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody
	GetData() *GetTokenResponseBodyData
	SetHttpStatusCode(v string) *GetTokenResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTokenResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetTokenResponseBody
	GetSuccess() *string
}

type GetTokenResponseBody struct {
	// example:
	//
	// success
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// B08AAA14-AD93-51F6-82AE-82AFAE9375B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenResponseBody) GetData() *GetTokenResponseBodyData {
	return s.Data
}

func (s *GetTokenResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetHttpStatusCode(v string) *GetTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v string) *GetTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenResponseBodyData struct {
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// b79d692c315d6bfb28312edf15
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// example:
	//
	// 127.0.0.1
	RequestIp *string `json:"requestIp,omitempty" xml:"requestIp,omitempty"`
	// example:
	//
	// 1748413248360
	ResponseTime *string `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	// example:
	//
	// N1faAjFhhaRNFaZNC8woRpQyAzEfBaIoWQEgDfds/Fwm7nIyEDLlSK3Ttx2OFebiHZ/MpHRr/3MnI/jpVWB/xNYIQxm6sccHJENHNAz6gaW+itU5wUrh+46EpqySABV8kc2pQ0HmYlbePfjjOK6lCfQjEGpekSAgQ6tDhG1lXWfKdtggq58Ut5bImMxMhk4R/PFUWrJe4CDuFu072C+foI0JlUV9TnGtVQ58oz8VRndrGXyauS/xqg8iGSZn6FyprUf5p+0ow20E
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetTokenResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetTokenResponseBodyData) GetNonce() *string {
	return s.Nonce
}

func (s *GetTokenResponseBodyData) GetRequestIp() *string {
	return s.RequestIp
}

func (s *GetTokenResponseBodyData) GetResponseTime() *string {
	return s.ResponseTime
}

func (s *GetTokenResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetTokenResponseBodyData) SetAppId(v string) *GetTokenResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetTokenResponseBodyData) SetDeviceName(v string) *GetTokenResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetTokenResponseBodyData) SetNonce(v string) *GetTokenResponseBodyData {
	s.Nonce = &v
	return s
}

func (s *GetTokenResponseBodyData) SetRequestIp(v string) *GetTokenResponseBodyData {
	s.RequestIp = &v
	return s
}

func (s *GetTokenResponseBodyData) SetResponseTime(v string) *GetTokenResponseBodyData {
	s.ResponseTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSignature(v string) *GetTokenResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
