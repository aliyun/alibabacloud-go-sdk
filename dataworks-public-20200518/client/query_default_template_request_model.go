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
	// The tenant ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console), go to the DataStudio page, click your username in the upper-right corner, and choose Menu > User Info to obtain the tenant ID.
	//
	// You can also obtain the tenant ID from Data.TenantId in the response of the GetProject operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024102
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
