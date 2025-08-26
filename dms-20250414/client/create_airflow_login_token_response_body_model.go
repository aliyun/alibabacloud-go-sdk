// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAirflowLoginTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAirflowLoginTokenResponseBody
	GetCode() *string
	SetData(v *CreateAirflowLoginTokenResponseBodyData) *CreateAirflowLoginTokenResponseBody
	GetData() *CreateAirflowLoginTokenResponseBodyData
	SetErrorCode(v string) *CreateAirflowLoginTokenResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *CreateAirflowLoginTokenResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAirflowLoginTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAirflowLoginTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAirflowLoginTokenResponseBody
	GetSuccess() *bool
}

type CreateAirflowLoginTokenResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the site monitoring task.
	Data *CreateAirflowLoginTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the call failed. Variable description:
	//
	// 	- If the request was successful, this parameter is not returned.
	//
	// 	- This parameter is returned only if the request failed.
	//
	// For more information, see the "Error codes" section in this topic.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The description of the error code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 4284D079-30F4-5B23-ADC4-28F291622C9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAirflowLoginTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAirflowLoginTokenResponseBody) GetData() *CreateAirflowLoginTokenResponseBodyData {
	return s.Data
}

func (s *CreateAirflowLoginTokenResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAirflowLoginTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAirflowLoginTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAirflowLoginTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAirflowLoginTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAirflowLoginTokenResponseBody) SetCode(v string) *CreateAirflowLoginTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetData(v *CreateAirflowLoginTokenResponseBodyData) *CreateAirflowLoginTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetErrorCode(v string) *CreateAirflowLoginTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetHttpStatusCode(v int32) *CreateAirflowLoginTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetMessage(v string) *CreateAirflowLoginTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetRequestId(v string) *CreateAirflowLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetSuccess(v bool) *CreateAirflowLoginTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAirflowLoginTokenResponseBodyData struct {
	// The endpoint that is used to access the Airflow instance.
	//
	// example:
	//
	// https://data-dms.aliyuncs.com/airflow/x/xxxx/af-ehrmszbxxxxxxx
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The generated token.
	//
	// example:
	//
	// f432d77de03b6b95fc24f91414e29c
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateAirflowLoginTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowLoginTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *CreateAirflowLoginTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *CreateAirflowLoginTokenResponseBodyData) SetHost(v string) *CreateAirflowLoginTokenResponseBodyData {
	s.Host = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBodyData) SetToken(v string) *CreateAirflowLoginTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
