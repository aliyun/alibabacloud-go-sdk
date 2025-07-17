// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFolderResponseBodyData) *GetFolderResponseBody
	GetData() *GetFolderResponseBodyData
	SetErrorCode(v string) *GetFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetFolderResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetFolderResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFolderResponseBody
	GetSuccess() *bool
}

type GetFolderResponseBody struct {
	Data *GetFolderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBody) GetData() *GetFolderResponseBodyData {
	return s.Data
}

func (s *GetFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetFolderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFolderResponseBody) SetData(v *GetFolderResponseBodyData) *GetFolderResponseBody {
	s.Data = v
	return s
}

func (s *GetFolderResponseBody) SetErrorCode(v string) *GetFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFolderResponseBody) SetErrorMessage(v string) *GetFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFolderResponseBody) SetHttpStatusCode(v int32) *GetFolderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFolderResponseBody) SetRequestId(v string) *GetFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFolderResponseBody) SetSuccess(v bool) *GetFolderResponseBody {
	s.Success = &v
	return s
}

func (s *GetFolderResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFolderResponseBodyData struct {
	// example:
	//
	// 2735****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FolderPath *string `json:"FolderPath,omitempty" xml:"FolderPath,omitempty"`
}

func (s GetFolderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFolderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBodyData) GetFolderId() *string {
	return s.FolderId
}

func (s *GetFolderResponseBodyData) GetFolderPath() *string {
	return s.FolderPath
}

func (s *GetFolderResponseBodyData) SetFolderId(v string) *GetFolderResponseBodyData {
	s.FolderId = &v
	return s
}

func (s *GetFolderResponseBodyData) SetFolderPath(v string) *GetFolderResponseBodyData {
	s.FolderPath = &v
	return s
}

func (s *GetFolderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
