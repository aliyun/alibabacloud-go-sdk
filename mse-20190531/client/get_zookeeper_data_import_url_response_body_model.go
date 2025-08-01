// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetZookeeperDataImportUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetZookeeperDataImportUrlResponseBody
	GetCode() *int32
	SetData(v *GetZookeeperDataImportUrlResponseBodyData) *GetZookeeperDataImportUrlResponseBody
	GetData() *GetZookeeperDataImportUrlResponseBodyData
	SetDynamicCode(v string) *GetZookeeperDataImportUrlResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetZookeeperDataImportUrlResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetZookeeperDataImportUrlResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetZookeeperDataImportUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetZookeeperDataImportUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetZookeeperDataImportUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetZookeeperDataImportUrlResponseBody
	GetSuccess() *bool
}

type GetZookeeperDataImportUrlResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The URL that is used to upload the configuration file.
	Data *GetZookeeperDataImportUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The returned data.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request was successfully processed.
	//
	// example:
	//
	// The dynamic part in the error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// E4E2058F-C524-5C29-9BC7-5874EA8D7CE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. A value of 200 is returned if the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetZookeeperDataImportUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetZookeeperDataImportUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetZookeeperDataImportUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetZookeeperDataImportUrlResponseBody) GetData() *GetZookeeperDataImportUrlResponseBodyData {
	return s.Data
}

func (s *GetZookeeperDataImportUrlResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetZookeeperDataImportUrlResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetZookeeperDataImportUrlResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetZookeeperDataImportUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetZookeeperDataImportUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetZookeeperDataImportUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetZookeeperDataImportUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetZookeeperDataImportUrlResponseBody) SetCode(v int32) *GetZookeeperDataImportUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetData(v *GetZookeeperDataImportUrlResponseBodyData) *GetZookeeperDataImportUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetDynamicCode(v string) *GetZookeeperDataImportUrlResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetDynamicMessage(v string) *GetZookeeperDataImportUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetErrorCode(v string) *GetZookeeperDataImportUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetHttpStatusCode(v int32) *GetZookeeperDataImportUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetMessage(v string) *GetZookeeperDataImportUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetRequestId(v string) *GetZookeeperDataImportUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) SetSuccess(v bool) *GetZookeeperDataImportUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetZookeeperDataImportUrlResponseBodyData struct {
	// code
	//
	// example:
	//
	// 250000
	MaxSize *string `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The maximum size of a file that can be uploaded each time. Unit: MB.
	//
	// example:
	//
	// http://xxxxxxxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetZookeeperDataImportUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetZookeeperDataImportUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetZookeeperDataImportUrlResponseBodyData) GetMaxSize() *string {
	return s.MaxSize
}

func (s *GetZookeeperDataImportUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetZookeeperDataImportUrlResponseBodyData) SetMaxSize(v string) *GetZookeeperDataImportUrlResponseBodyData {
	s.MaxSize = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBodyData) SetUrl(v string) *GetZookeeperDataImportUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
