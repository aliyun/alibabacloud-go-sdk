// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClassificationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *GetClassificationTemplateRequest
	GetInstanceId() *int64
	SetTid(v int64) *GetClassificationTemplateRequest
	GetTid() *int64
}

type GetClassificationTemplateRequest struct {
	// The ID of the instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 23***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetClassificationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClassificationTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetClassificationTemplateRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetClassificationTemplateRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetClassificationTemplateRequest) SetInstanceId(v int64) *GetClassificationTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetClassificationTemplateRequest) SetTid(v int64) *GetClassificationTemplateRequest {
	s.Tid = &v
	return s
}

func (s *GetClassificationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
