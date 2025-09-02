// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableIntroWikiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableIntroWikiResponseBodyData) *GetMetaTableIntroWikiResponseBody
	GetData() *GetMetaTableIntroWikiResponseBodyData
	SetErrorCode(v string) *GetMetaTableIntroWikiResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableIntroWikiResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableIntroWikiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableIntroWikiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableIntroWikiResponseBody
	GetSuccess() *bool
}

type GetMetaTableIntroWikiResponseBody struct {
	// The business data.
	Data *GetMetaTableIntroWikiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
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
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableIntroWikiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableIntroWikiResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableIntroWikiResponseBody) GetData() *GetMetaTableIntroWikiResponseBodyData {
	return s.Data
}

func (s *GetMetaTableIntroWikiResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableIntroWikiResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableIntroWikiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableIntroWikiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableIntroWikiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableIntroWikiResponseBody) SetData(v *GetMetaTableIntroWikiResponseBodyData) *GetMetaTableIntroWikiResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableIntroWikiResponseBody) SetErrorCode(v string) *GetMetaTableIntroWikiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBody) SetErrorMessage(v string) *GetMetaTableIntroWikiResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBody) SetHttpStatusCode(v int32) *GetMetaTableIntroWikiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBody) SetRequestId(v string) *GetMetaTableIntroWikiResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBody) SetSuccess(v bool) *GetMetaTableIntroWikiResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableIntroWikiResponseBodyData struct {
	// The description of the metatable.
	//
	// example:
	//
	// \\# Business requirements\\n\\n\\&lt;a name=\\&quot;xiw5n\\&quot;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the metatable was created.
	//
	// example:
	//
	// 1584444247000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the metatable.
	//
	// example:
	//
	// abc
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The name of the user who creates the metatable.
	//
	// example:
	//
	// abc
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The time when the metatable was last modified.
	//
	// example:
	//
	// 1584444247000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetMetaTableIntroWikiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableIntroWikiResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableIntroWikiResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetMetaTableIntroWikiResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaTableIntroWikiResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetMetaTableIntroWikiResponseBodyData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetMetaTableIntroWikiResponseBodyData) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetMetaTableIntroWikiResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *GetMetaTableIntroWikiResponseBodyData) SetContent(v string) *GetMetaTableIntroWikiResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBodyData) SetCreateTime(v int64) *GetMetaTableIntroWikiResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBodyData) SetCreator(v string) *GetMetaTableIntroWikiResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBodyData) SetCreatorName(v string) *GetMetaTableIntroWikiResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBodyData) SetModifiedTime(v int64) *GetMetaTableIntroWikiResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBodyData) SetVersion(v int64) *GetMetaTableIntroWikiResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetMetaTableIntroWikiResponseBodyData) Validate() error {
	return dara.Validate(s)
}
