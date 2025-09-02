// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataServiceApplicationResponseBodyData) *GetDataServiceApplicationResponseBody
	GetData() *GetDataServiceApplicationResponseBodyData
	SetErrorCode(v string) *GetDataServiceApplicationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataServiceApplicationResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetDataServiceApplicationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetDataServiceApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceApplicationResponseBody
	GetSuccess() *bool
}

type GetDataServiceApplicationResponseBody struct {
	// The details of the application.
	Data *GetDataServiceApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Normal
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApplicationResponseBody) GetData() *GetDataServiceApplicationResponseBodyData {
	return s.Data
}

func (s *GetDataServiceApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataServiceApplicationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataServiceApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceApplicationResponseBody) SetData(v *GetDataServiceApplicationResponseBodyData) *GetDataServiceApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceApplicationResponseBody) SetErrorCode(v string) *GetDataServiceApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataServiceApplicationResponseBody) SetErrorMessage(v string) *GetDataServiceApplicationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataServiceApplicationResponseBody) SetHttpStatusCode(v int32) *GetDataServiceApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceApplicationResponseBody) SetRequestId(v string) *GetDataServiceApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceApplicationResponseBody) SetSuccess(v bool) *GetDataServiceApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApplicationResponseBodyData struct {
	// The AppCode for simple authentication. You can select simple authentication or signature authentication when you call an API operation.
	//
	// example:
	//
	// CODE123
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 10000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The AppKey for signature authentication. You can select simple authentication or signature authentication when you call an API operation.
	//
	// example:
	//
	// KEY123
	ApplicationKey *string `json:"ApplicationKey,omitempty" xml:"ApplicationKey,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// Test application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The AppSecret for signature authentication. You can select simple authentication or signature authentication when you call an API operation.
	//
	// example:
	//
	// SECRET123
	ApplicationSecret *string `json:"ApplicationSecret,omitempty" xml:"ApplicationSecret,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDataServiceApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceApplicationResponseBodyData) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *GetDataServiceApplicationResponseBodyData) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *GetDataServiceApplicationResponseBodyData) GetApplicationKey() *string {
	return s.ApplicationKey
}

func (s *GetDataServiceApplicationResponseBodyData) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetDataServiceApplicationResponseBodyData) GetApplicationSecret() *string {
	return s.ApplicationSecret
}

func (s *GetDataServiceApplicationResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServiceApplicationResponseBodyData) SetApplicationCode(v string) *GetDataServiceApplicationResponseBodyData {
	s.ApplicationCode = &v
	return s
}

func (s *GetDataServiceApplicationResponseBodyData) SetApplicationId(v int64) *GetDataServiceApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *GetDataServiceApplicationResponseBodyData) SetApplicationKey(v string) *GetDataServiceApplicationResponseBodyData {
	s.ApplicationKey = &v
	return s
}

func (s *GetDataServiceApplicationResponseBodyData) SetApplicationName(v string) *GetDataServiceApplicationResponseBodyData {
	s.ApplicationName = &v
	return s
}

func (s *GetDataServiceApplicationResponseBodyData) SetApplicationSecret(v string) *GetDataServiceApplicationResponseBodyData {
	s.ApplicationSecret = &v
	return s
}

func (s *GetDataServiceApplicationResponseBodyData) SetProjectId(v int64) *GetDataServiceApplicationResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
