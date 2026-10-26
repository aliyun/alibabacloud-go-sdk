// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTemplateResponseBody
	GetCode() *string
	SetCreatedTime(v string) *GetTemplateResponseBody
	GetCreatedTime() *string
	SetMessage(v string) *GetTemplateResponseBody
	GetMessage() *string
	SetName(v string) *GetTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetResourceGroupID(v string) *GetTemplateResponseBody
	GetResourceGroupID() *string
	SetResourceTags(v []*GetTemplateResponseBodyResourceTags) *GetTemplateResponseBody
	GetResourceTags() []*GetTemplateResponseBodyResourceTags
	SetRuntimeConfig(v *PublicTemplateRuntimeConfig) *GetTemplateResponseBody
	GetRuntimeConfig() *PublicTemplateRuntimeConfig
	SetStatus(v *PublicTemplateStatus) *GetTemplateResponseBody
	GetStatus() *PublicTemplateStatus
	SetTeamID(v string) *GetTemplateResponseBody
	GetTeamID() *string
	SetTeamName(v string) *GetTemplateResponseBody
	GetTeamName() *string
	SetTemplateID(v string) *GetTemplateResponseBody
	GetTemplateID() *string
}

type GetTemplateResponseBody struct {
	// The error code. The value is 200 if the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The creation time in UTC.
	//
	// example:
	//
	// 2026-09-19T10:00:00Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// my-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// The resource tags.
	ResourceTags []*GetTemplateResponseBodyResourceTags `json:"resourceTags" xml:"resourceTags" type:"Repeated"`
	// The runtime configuration of the template.
	RuntimeConfig *PublicTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	// The status of the template.
	Status *PublicTemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// The team ID.
	//
	// example:
	//
	// 13b721e6-8cc8-5df2-af13-80316f7508af
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The team name.
	//
	// example:
	//
	// my-team
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	// The unique identifier of the template.
	//
	// example:
	//
	// us7dxqaezw5uu7aa2cm5
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTemplateResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *GetTemplateResponseBody) GetResourceTags() []*GetTemplateResponseBodyResourceTags {
	return s.ResourceTags
}

func (s *GetTemplateResponseBody) GetRuntimeConfig() *PublicTemplateRuntimeConfig {
	return s.RuntimeConfig
}

func (s *GetTemplateResponseBody) GetStatus() *PublicTemplateStatus {
	return s.Status
}

func (s *GetTemplateResponseBody) GetTeamID() *string {
	return s.TeamID
}

func (s *GetTemplateResponseBody) GetTeamName() *string {
	return s.TeamName
}

func (s *GetTemplateResponseBody) GetTemplateID() *string {
	return s.TemplateID
}

func (s *GetTemplateResponseBody) SetCode(v string) *GetTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateResponseBody) SetCreatedTime(v string) *GetTemplateResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetMessage(v string) *GetTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateResponseBody) SetName(v string) *GetTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetResourceGroupID(v string) *GetTemplateResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *GetTemplateResponseBody) SetResourceTags(v []*GetTemplateResponseBodyResourceTags) *GetTemplateResponseBody {
	s.ResourceTags = v
	return s
}

func (s *GetTemplateResponseBody) SetRuntimeConfig(v *PublicTemplateRuntimeConfig) *GetTemplateResponseBody {
	s.RuntimeConfig = v
	return s
}

func (s *GetTemplateResponseBody) SetStatus(v *PublicTemplateStatus) *GetTemplateResponseBody {
	s.Status = v
	return s
}

func (s *GetTemplateResponseBody) SetTeamID(v string) *GetTemplateResponseBody {
	s.TeamID = &v
	return s
}

func (s *GetTemplateResponseBody) SetTeamName(v string) *GetTemplateResponseBody {
	s.TeamName = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateID(v string) *GetTemplateResponseBody {
	s.TemplateID = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuntimeConfig != nil {
		if err := s.RuntimeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTemplateResponseBodyResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// project
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// demo
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s GetTemplateResponseBodyResourceTags) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyResourceTags) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyResourceTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetTemplateResponseBodyResourceTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetTemplateResponseBodyResourceTags) SetTagKey(v string) *GetTemplateResponseBodyResourceTags {
	s.TagKey = &v
	return s
}

func (s *GetTemplateResponseBodyResourceTags) SetTagValue(v string) *GetTemplateResponseBodyResourceTags {
	s.TagValue = &v
	return s
}

func (s *GetTemplateResponseBodyResourceTags) Validate() error {
	return dara.Validate(s)
}
