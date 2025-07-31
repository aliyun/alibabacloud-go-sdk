// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfByVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUdfByVersionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUdfByVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUdfByVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUdfByVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUdfByVersionResponseBody
	GetSuccess() *bool
	SetUdfInfo(v *GetUdfByVersionResponseBodyUdfInfo) *GetUdfByVersionResponseBody
	GetUdfInfo() *GetUdfByVersionResponseBodyUdfInfo
}

type GetUdfByVersionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	UdfInfo   *GetUdfByVersionResponseBodyUdfInfo `json:"UdfInfo,omitempty" xml:"UdfInfo,omitempty" type:"Struct"`
}

func (s GetUdfByVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUdfByVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetUdfByVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUdfByVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUdfByVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUdfByVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUdfByVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUdfByVersionResponseBody) GetUdfInfo() *GetUdfByVersionResponseBodyUdfInfo {
	return s.UdfInfo
}

func (s *GetUdfByVersionResponseBody) SetCode(v string) *GetUdfByVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetUdfByVersionResponseBody) SetHttpStatusCode(v int32) *GetUdfByVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUdfByVersionResponseBody) SetMessage(v string) *GetUdfByVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetUdfByVersionResponseBody) SetRequestId(v string) *GetUdfByVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUdfByVersionResponseBody) SetSuccess(v bool) *GetUdfByVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetUdfByVersionResponseBody) SetUdfInfo(v *GetUdfByVersionResponseBodyUdfInfo) *GetUdfByVersionResponseBody {
	s.UdfInfo = v
	return s
}

func (s *GetUdfByVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUdfByVersionResponseBodyUdfInfo struct {
	// example:
	//
	// 10
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// com.lydaas.dataphin.UdfTest
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// example:
	//
	// udf_to_lower(char x)
	CommandHelp *string `json:"CommandHelp,omitempty" xml:"CommandHelp,omitempty"`
	// example:
	//
	// HADOOP
	ComputeEngineType *string `json:"ComputeEngineType,omitempty" xml:"ComputeEngineType,omitempty"`
	// example:
	//
	// 30012110
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30012110
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// udf_to_lower
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUdfByVersionResponseBodyUdfInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUdfByVersionResponseBodyUdfInfo) GoString() string {
	return s.String()
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetCategory() *int32 {
	return s.Category
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetClassName() *string {
	return s.ClassName
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetCommandHelp() *string {
	return s.CommandHelp
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetComputeEngineType() *string {
	return s.ComputeEngineType
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetDescription() *string {
	return s.Description
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetDirectory() *string {
	return s.Directory
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetId() *int64 {
	return s.Id
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetUdfByVersionResponseBodyUdfInfo) GetName() *string {
	return s.Name
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetCategory(v int32) *GetUdfByVersionResponseBodyUdfInfo {
	s.Category = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetClassName(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.ClassName = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetCommandHelp(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.CommandHelp = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetComputeEngineType(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.ComputeEngineType = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetCreator(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.Creator = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetDescription(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.Description = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetDirectory(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.Directory = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetGmtCreate(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetGmtModified(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.GmtModified = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetId(v int64) *GetUdfByVersionResponseBodyUdfInfo {
	s.Id = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetLastModifier(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.LastModifier = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) SetName(v string) *GetUdfByVersionResponseBodyUdfInfo {
	s.Name = &v
	return s
}

func (s *GetUdfByVersionResponseBodyUdfInfo) Validate() error {
	return dara.Validate(s)
}
