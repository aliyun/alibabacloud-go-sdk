// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUdfResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUdfResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUdfResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUdfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUdfResponseBody
	GetSuccess() *bool
	SetUdfInfo(v *GetUdfResponseBodyUdfInfo) *GetUdfResponseBody
	GetUdfInfo() *GetUdfResponseBodyUdfInfo
}

type GetUdfResponseBody struct {
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
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	UdfInfo   *GetUdfResponseBodyUdfInfo `json:"UdfInfo,omitempty" xml:"UdfInfo,omitempty" type:"Struct"`
}

func (s GetUdfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUdfResponseBody) GoString() string {
	return s.String()
}

func (s *GetUdfResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUdfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUdfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUdfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUdfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUdfResponseBody) GetUdfInfo() *GetUdfResponseBodyUdfInfo {
	return s.UdfInfo
}

func (s *GetUdfResponseBody) SetCode(v string) *GetUdfResponseBody {
	s.Code = &v
	return s
}

func (s *GetUdfResponseBody) SetHttpStatusCode(v int32) *GetUdfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUdfResponseBody) SetMessage(v string) *GetUdfResponseBody {
	s.Message = &v
	return s
}

func (s *GetUdfResponseBody) SetRequestId(v string) *GetUdfResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUdfResponseBody) SetSuccess(v bool) *GetUdfResponseBody {
	s.Success = &v
	return s
}

func (s *GetUdfResponseBody) SetUdfInfo(v *GetUdfResponseBodyUdfInfo) *GetUdfResponseBody {
	s.UdfInfo = v
	return s
}

func (s *GetUdfResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUdfResponseBodyUdfInfo struct {
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

func (s GetUdfResponseBodyUdfInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUdfResponseBodyUdfInfo) GoString() string {
	return s.String()
}

func (s *GetUdfResponseBodyUdfInfo) GetCategory() *int32 {
	return s.Category
}

func (s *GetUdfResponseBodyUdfInfo) GetClassName() *string {
	return s.ClassName
}

func (s *GetUdfResponseBodyUdfInfo) GetCommandHelp() *string {
	return s.CommandHelp
}

func (s *GetUdfResponseBodyUdfInfo) GetComputeEngineType() *string {
	return s.ComputeEngineType
}

func (s *GetUdfResponseBodyUdfInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetUdfResponseBodyUdfInfo) GetDescription() *string {
	return s.Description
}

func (s *GetUdfResponseBodyUdfInfo) GetDirectory() *string {
	return s.Directory
}

func (s *GetUdfResponseBodyUdfInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetUdfResponseBodyUdfInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetUdfResponseBodyUdfInfo) GetId() *int64 {
	return s.Id
}

func (s *GetUdfResponseBodyUdfInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetUdfResponseBodyUdfInfo) GetName() *string {
	return s.Name
}

func (s *GetUdfResponseBodyUdfInfo) SetCategory(v int32) *GetUdfResponseBodyUdfInfo {
	s.Category = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetClassName(v string) *GetUdfResponseBodyUdfInfo {
	s.ClassName = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetCommandHelp(v string) *GetUdfResponseBodyUdfInfo {
	s.CommandHelp = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetComputeEngineType(v string) *GetUdfResponseBodyUdfInfo {
	s.ComputeEngineType = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetCreator(v string) *GetUdfResponseBodyUdfInfo {
	s.Creator = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetDescription(v string) *GetUdfResponseBodyUdfInfo {
	s.Description = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetDirectory(v string) *GetUdfResponseBodyUdfInfo {
	s.Directory = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetGmtCreate(v string) *GetUdfResponseBodyUdfInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetGmtModified(v string) *GetUdfResponseBodyUdfInfo {
	s.GmtModified = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetId(v int64) *GetUdfResponseBodyUdfInfo {
	s.Id = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetLastModifier(v string) *GetUdfResponseBodyUdfInfo {
	s.LastModifier = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) SetName(v string) *GetUdfResponseBodyUdfInfo {
	s.Name = &v
	return s
}

func (s *GetUdfResponseBodyUdfInfo) Validate() error {
	return dara.Validate(s)
}
