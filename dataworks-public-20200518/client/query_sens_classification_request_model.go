// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensClassificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *QuerySensClassificationRequest
	GetTemplateId() *string
	SetTenantId(v string) *QuerySensClassificationRequest
	GetTenantId() *string
}

type QuerySensClassificationRequest struct {
	// The ID of the template defined by Data Security Guard. You can call the [QueryDefaultTemplate](https://help.aliyun.com/document_detail/2743948.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1970541-6cf5-4d23-b101-d5b66f6e1024
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QuerySensClassificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySensClassificationRequest) GoString() string {
	return s.String()
}

func (s *QuerySensClassificationRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QuerySensClassificationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySensClassificationRequest) SetTemplateId(v string) *QuerySensClassificationRequest {
	s.TemplateId = &v
	return s
}

func (s *QuerySensClassificationRequest) SetTenantId(v string) *QuerySensClassificationRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySensClassificationRequest) Validate() error {
	return dara.Validate(s)
}
