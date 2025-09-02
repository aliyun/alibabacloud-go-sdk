// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnabledExtensionsForProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtensions(v []*ListEnabledExtensionsForProjectResponseBodyExtensions) *ListEnabledExtensionsForProjectResponseBody
	GetExtensions() []*ListEnabledExtensionsForProjectResponseBodyExtensions
	SetRequestId(v string) *ListEnabledExtensionsForProjectResponseBody
	GetRequestId() *string
}

type ListEnabledExtensionsForProjectResponseBody struct {
	// The extensions.
	Extensions []*ListEnabledExtensionsForProjectResponseBodyExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4F2CA7ED-27E5-59EA-A8C4-F1F7A1FF0B22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnabledExtensionsForProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnabledExtensionsForProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnabledExtensionsForProjectResponseBody) GetExtensions() []*ListEnabledExtensionsForProjectResponseBodyExtensions {
	return s.Extensions
}

func (s *ListEnabledExtensionsForProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnabledExtensionsForProjectResponseBody) SetExtensions(v []*ListEnabledExtensionsForProjectResponseBodyExtensions) *ListEnabledExtensionsForProjectResponseBody {
	s.Extensions = v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBody) SetRequestId(v string) *ListEnabledExtensionsForProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEnabledExtensionsForProjectResponseBodyExtensions struct {
	// The creator of the extension.
	//
	// example:
	//
	// 3444434343555
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The unique code of the extension.
	//
	// example:
	//
	// a94a8e23bc0b4dfab9a5e4d2f374d645
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The description of the extension.
	//
	// example:
	//
	// ODPS SQL compatible Spark engine detection
	ExtensionDesc *string `json:"ExtensionDesc,omitempty" xml:"ExtensionDesc,omitempty"`
	// The name of the extension.
	//
	// example:
	//
	// max_pt function is not allowed.
	ExtensionName *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
	// The modifier of the extension.
	//
	// example:
	//
	// 34452335611988
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 21323672*******55500
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameter settings of the extension. For more information, see [Configure extension parameters](https://help.aliyun.com/document_detail/405354.html).
	//
	// example:
	//
	// extension.fileType.23.deploy-file.enabled=true
	ParameterSetting *string `json:"ParameterSetting,omitempty" xml:"ParameterSetting,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 529889518659842
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListEnabledExtensionsForProjectResponseBodyExtensions) String() string {
	return dara.Prettify(s)
}

func (s ListEnabledExtensionsForProjectResponseBodyExtensions) GoString() string {
	return s.String()
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetExtensionDesc() *string {
	return s.ExtensionDesc
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetExtensionName() *string {
	return s.ExtensionName
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetOwner() *string {
	return s.Owner
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetParameterSetting() *string {
	return s.ParameterSetting
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetCreateUser(v string) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.CreateUser = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetExtensionCode(v string) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.ExtensionCode = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetExtensionDesc(v string) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.ExtensionDesc = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetExtensionName(v string) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.ExtensionName = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetModifyUser(v string) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.ModifyUser = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetOwner(v string) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.Owner = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetParameterSetting(v string) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.ParameterSetting = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) SetTenantId(v int64) *ListEnabledExtensionsForProjectResponseBodyExtensions {
	s.TenantId = &v
	return s
}

func (s *ListEnabledExtensionsForProjectResponseBodyExtensions) Validate() error {
	return dara.Validate(s)
}
