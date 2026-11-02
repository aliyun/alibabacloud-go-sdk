// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceName(v string) *GetServiceLinkedRoleTemplateRequest
	GetServiceName() *string
}

type GetServiceLinkedRoleTemplateRequest struct {
	// The cloud service name.
	//
	// For more information, see the **Cloud service identity*	- column in [Cloud services that support service-linked roles](https://help.aliyun.com/document_detail/461722.html).
	//
	// example:
	//
	// polardb.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetServiceLinkedRoleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceLinkedRoleTemplateRequest) SetServiceName(v string) *GetServiceLinkedRoleTemplateRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateRequest) Validate() error {
	return dara.Validate(s)
}
