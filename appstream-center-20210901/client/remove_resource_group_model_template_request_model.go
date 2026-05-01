// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveResourceGroupModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelTemplateId(v string) *RemoveResourceGroupModelTemplateRequest
	GetModelTemplateId() *string
	SetResourceGroupIds(v []*string) *RemoveResourceGroupModelTemplateRequest
	GetResourceGroupIds() []*string
}

type RemoveResourceGroupModelTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// This parameter is required.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s RemoveResourceGroupModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveResourceGroupModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *RemoveResourceGroupModelTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *RemoveResourceGroupModelTemplateRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *RemoveResourceGroupModelTemplateRequest) SetModelTemplateId(v string) *RemoveResourceGroupModelTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *RemoveResourceGroupModelTemplateRequest) SetResourceGroupIds(v []*string) *RemoveResourceGroupModelTemplateRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *RemoveResourceGroupModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
