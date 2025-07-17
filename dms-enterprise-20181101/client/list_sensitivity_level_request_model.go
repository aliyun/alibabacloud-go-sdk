// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitivityLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *ListSensitivityLevelRequest
	GetTemplateId() *int64
	SetTemplateType(v string) *ListSensitivityLevelRequest
	GetTemplateType() *string
	SetTid(v int64) *ListSensitivityLevelRequest
	GetTid() *int64
}

type ListSensitivityLevelRequest struct {
	// The ID of the classification template. You can call the [ListClassificationTemplates](https://help.aliyun.com/document_detail/460613.html) operation to query the ID of the classification template.
	//
	// example:
	//
	// 15**
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the classification template. You can call the [ListClassificationTemplates](https://help.aliyun.com/document_detail/460613.html) operation to query the type of the classification template.
	//
	// Valid values:
	//
	// 	- USER_DEFINE: a custom template.
	//
	// 	- INNER: a built-in template.
	//
	// example:
	//
	// INNER
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) in the topic "Manage DMS tenants."
	//
	// example:
	//
	// 20***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSensitivityLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSensitivityLevelRequest) GoString() string {
	return s.String()
}

func (s *ListSensitivityLevelRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListSensitivityLevelRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListSensitivityLevelRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSensitivityLevelRequest) SetTemplateId(v int64) *ListSensitivityLevelRequest {
	s.TemplateId = &v
	return s
}

func (s *ListSensitivityLevelRequest) SetTemplateType(v string) *ListSensitivityLevelRequest {
	s.TemplateType = &v
	return s
}

func (s *ListSensitivityLevelRequest) SetTid(v int64) *ListSensitivityLevelRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitivityLevelRequest) Validate() error {
	return dara.Validate(s)
}
