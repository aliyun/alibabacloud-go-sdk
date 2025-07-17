// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListBusinessResponseBodyData) *ListBusinessResponseBody
	GetData() *ListBusinessResponseBodyData
	SetErrorCode(v string) *ListBusinessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListBusinessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListBusinessResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListBusinessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBusinessResponseBody
	GetSuccess() *bool
}

type ListBusinessResponseBody struct {
	Data *ListBusinessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListBusinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusinessResponseBody) GetData() *ListBusinessResponseBodyData {
	return s.Data
}

func (s *ListBusinessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListBusinessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListBusinessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBusinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBusinessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBusinessResponseBody) SetData(v *ListBusinessResponseBodyData) *ListBusinessResponseBody {
	s.Data = v
	return s
}

func (s *ListBusinessResponseBody) SetErrorCode(v string) *ListBusinessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListBusinessResponseBody) SetErrorMessage(v string) *ListBusinessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListBusinessResponseBody) SetHttpStatusCode(v int32) *ListBusinessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBusinessResponseBody) SetRequestId(v string) *ListBusinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusinessResponseBody) SetSuccess(v bool) *ListBusinessResponseBody {
	s.Success = &v
	return s
}

func (s *ListBusinessResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBusinessResponseBodyData struct {
	Business []*ListBusinessResponseBodyDataBusiness `json:"Business,omitempty" xml:"Business,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBusinessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBusinessResponseBodyData) GetBusiness() []*ListBusinessResponseBodyDataBusiness {
	return s.Business
}

func (s *ListBusinessResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBusinessResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBusinessResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBusinessResponseBodyData) SetBusiness(v []*ListBusinessResponseBodyDataBusiness) *ListBusinessResponseBodyData {
	s.Business = v
	return s
}

func (s *ListBusinessResponseBodyData) SetPageNumber(v int32) *ListBusinessResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBusinessResponseBodyData) SetPageSize(v int32) *ListBusinessResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBusinessResponseBodyData) SetTotalCount(v int32) *ListBusinessResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBusinessResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListBusinessResponseBodyDataBusiness struct {
	// example:
	//
	// 3000001
	BusinessId   *int64  `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 34824327****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s ListBusinessResponseBodyDataBusiness) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessResponseBodyDataBusiness) GoString() string {
	return s.String()
}

func (s *ListBusinessResponseBodyDataBusiness) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListBusinessResponseBodyDataBusiness) GetBusinessName() *string {
	return s.BusinessName
}

func (s *ListBusinessResponseBodyDataBusiness) GetDescription() *string {
	return s.Description
}

func (s *ListBusinessResponseBodyDataBusiness) GetOwner() *string {
	return s.Owner
}

func (s *ListBusinessResponseBodyDataBusiness) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBusinessResponseBodyDataBusiness) GetUseType() *string {
	return s.UseType
}

func (s *ListBusinessResponseBodyDataBusiness) SetBusinessId(v int64) *ListBusinessResponseBodyDataBusiness {
	s.BusinessId = &v
	return s
}

func (s *ListBusinessResponseBodyDataBusiness) SetBusinessName(v string) *ListBusinessResponseBodyDataBusiness {
	s.BusinessName = &v
	return s
}

func (s *ListBusinessResponseBodyDataBusiness) SetDescription(v string) *ListBusinessResponseBodyDataBusiness {
	s.Description = &v
	return s
}

func (s *ListBusinessResponseBodyDataBusiness) SetOwner(v string) *ListBusinessResponseBodyDataBusiness {
	s.Owner = &v
	return s
}

func (s *ListBusinessResponseBodyDataBusiness) SetProjectId(v int64) *ListBusinessResponseBodyDataBusiness {
	s.ProjectId = &v
	return s
}

func (s *ListBusinessResponseBodyDataBusiness) SetUseType(v string) *ListBusinessResponseBodyDataBusiness {
	s.UseType = &v
	return s
}

func (s *ListBusinessResponseBodyDataBusiness) Validate() error {
	return dara.Validate(s)
}
