// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupAgentModelSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetResourceGroupAgentModelSettingsResponseBodyData) *GetResourceGroupAgentModelSettingsResponseBody
	GetData() []*GetResourceGroupAgentModelSettingsResponseBodyData
	SetRequestId(v string) *GetResourceGroupAgentModelSettingsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetResourceGroupAgentModelSettingsResponseBody
	GetTotalCount() *int32
}

type GetResourceGroupAgentModelSettingsResponseBody struct {
	// The list of model settings for resource groups. One record is returned for each resource group ID specified in the request. If a resource group is not associated with a model group under the specified agent platform and provider, ModelTemplateId and ModelTemplateName are empty.
	Data []*GetResourceGroupAgentModelSettingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total count.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourceGroupAgentModelSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupAgentModelSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupAgentModelSettingsResponseBody) GetData() []*GetResourceGroupAgentModelSettingsResponseBodyData {
	return s.Data
}

func (s *GetResourceGroupAgentModelSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupAgentModelSettingsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetResourceGroupAgentModelSettingsResponseBody) SetData(v []*GetResourceGroupAgentModelSettingsResponseBodyData) *GetResourceGroupAgentModelSettingsResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponseBody) SetRequestId(v string) *GetResourceGroupAgentModelSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponseBody) SetTotalCount(v int32) *GetResourceGroupAgentModelSettingsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceGroupAgentModelSettingsResponseBodyData struct {
	// The model group ID.
	//
	// example:
	//
	// mt-xxxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The model group name.
	//
	// example:
	//
	// Default model group
	ModelTemplateName *string `json:"ModelTemplateName,omitempty" xml:"ModelTemplateName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetResourceGroupAgentModelSettingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupAgentModelSettingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceGroupAgentModelSettingsResponseBodyData) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *GetResourceGroupAgentModelSettingsResponseBodyData) GetModelTemplateName() *string {
	return s.ModelTemplateName
}

func (s *GetResourceGroupAgentModelSettingsResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetResourceGroupAgentModelSettingsResponseBodyData) SetModelTemplateId(v string) *GetResourceGroupAgentModelSettingsResponseBodyData {
	s.ModelTemplateId = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponseBodyData) SetModelTemplateName(v string) *GetResourceGroupAgentModelSettingsResponseBodyData {
	s.ModelTemplateName = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponseBodyData) SetResourceGroupId(v string) *GetResourceGroupAgentModelSettingsResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceGroupAgentModelSettingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
