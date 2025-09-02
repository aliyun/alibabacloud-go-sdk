// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *QuerySensLevelRequest
	GetTemplateId() *string
	SetTenantId(v string) *QuerySensLevelRequest
	GetTenantId() *string
}

type QuerySensLevelRequest struct {
	// The ID of the template defined by Data Security Guard. You can call the [QueryDefaultTemplate](https://help.aliyun.com/document_detail/2743948.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1970541-6cf5-4d23-b101-d5b66f6e10af
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QuerySensLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySensLevelRequest) GoString() string {
	return s.String()
}

func (s *QuerySensLevelRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QuerySensLevelRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySensLevelRequest) SetTemplateId(v string) *QuerySensLevelRequest {
	s.TemplateId = &v
	return s
}

func (s *QuerySensLevelRequest) SetTenantId(v string) *QuerySensLevelRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySensLevelRequest) Validate() error {
	return dara.Validate(s)
}
