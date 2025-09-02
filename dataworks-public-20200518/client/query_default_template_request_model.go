// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDefaultTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *QueryDefaultTemplateRequest
	GetTenantId() *string
}

type QueryDefaultTemplateRequest struct {
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryDefaultTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDefaultTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryDefaultTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryDefaultTemplateRequest) SetTenantId(v string) *QueryDefaultTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *QueryDefaultTemplateRequest) Validate() error {
	return dara.Validate(s)
}
