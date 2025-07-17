// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBusinessResponseBodyData) *GetBusinessResponseBody
	GetData() *GetBusinessResponseBodyData
	SetErrorCode(v string) *GetBusinessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetBusinessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetBusinessResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetBusinessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBusinessResponseBody
	GetSuccess() *bool
}

type GetBusinessResponseBody struct {
	Data *GetBusinessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetBusinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusinessResponseBody) GetData() *GetBusinessResponseBodyData {
	return s.Data
}

func (s *GetBusinessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBusinessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBusinessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBusinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBusinessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBusinessResponseBody) SetData(v *GetBusinessResponseBodyData) *GetBusinessResponseBody {
	s.Data = v
	return s
}

func (s *GetBusinessResponseBody) SetErrorCode(v string) *GetBusinessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBusinessResponseBody) SetErrorMessage(v string) *GetBusinessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBusinessResponseBody) SetHttpStatusCode(v int32) *GetBusinessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBusinessResponseBody) SetRequestId(v string) *GetBusinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBusinessResponseBody) SetSuccess(v bool) *GetBusinessResponseBody {
	s.Success = &v
	return s
}

func (s *GetBusinessResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBusinessResponseBodyData struct {
	// example:
	//
	// 1000001
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// example:
	//
	// The first business process
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// This is my first business process.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 20000****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s GetBusinessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBusinessResponseBodyData) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetBusinessResponseBodyData) GetBusinessName() *string {
	return s.BusinessName
}

func (s *GetBusinessResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetBusinessResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetBusinessResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetBusinessResponseBodyData) GetUseType() *string {
	return s.UseType
}

func (s *GetBusinessResponseBodyData) SetBusinessId(v int64) *GetBusinessResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetBusinessName(v string) *GetBusinessResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetDescription(v string) *GetBusinessResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetOwner(v string) *GetBusinessResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetProjectId(v string) *GetBusinessResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetUseType(v string) *GetBusinessResponseBodyData {
	s.UseType = &v
	return s
}

func (s *GetBusinessResponseBodyData) Validate() error {
	return dara.Validate(s)
}
