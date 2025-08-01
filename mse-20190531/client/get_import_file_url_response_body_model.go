// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImportFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetImportFileUrlResponseBody
	GetCode() *int32
	SetData(v *GetImportFileUrlResponseBodyData) *GetImportFileUrlResponseBody
	GetData() *GetImportFileUrlResponseBodyData
	SetDynamicMessage(v string) *GetImportFileUrlResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetImportFileUrlResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetImportFileUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetImportFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImportFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImportFileUrlResponseBody
	GetSuccess() *bool
}

type GetImportFileUrlResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data *GetImportFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic part in the error message. This parameter is used to replace the \\\\*\\\\*%s\\\\*\\\\	- variable in the **ErrMessage*	- parameter.\\n\\n>  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 19488A00-4AF5-55E1-A689-98BA9F5E7441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImportFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImportFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetImportFileUrlResponseBody) GetData() *GetImportFileUrlResponseBodyData {
	return s.Data
}

func (s *GetImportFileUrlResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetImportFileUrlResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetImportFileUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImportFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImportFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImportFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImportFileUrlResponseBody) SetCode(v int32) *GetImportFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetData(v *GetImportFileUrlResponseBodyData) *GetImportFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetImportFileUrlResponseBody) SetDynamicMessage(v string) *GetImportFileUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetErrorCode(v string) *GetImportFileUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetHttpStatusCode(v int32) *GetImportFileUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetMessage(v string) *GetImportFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetRequestId(v string) *GetImportFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetSuccess(v bool) *GetImportFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetImportFileUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImportFileUrlResponseBodyData struct {
	// The URL that is used to upload the configuration file.
	//
	// example:
	//
	// http://xxxxxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImportFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImportFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetImportFileUrlResponseBodyData) SetUrl(v string) *GetImportFileUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetImportFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
