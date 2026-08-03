// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeScenesResponseBody
	GetRequestId() *string
	SetSceneList(v []*DescribeScenesResponseBodySceneList) *DescribeScenesResponseBody
	GetSceneList() []*DescribeScenesResponseBodySceneList
}

type DescribeScenesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7EC26DF0-35AC-5F37-82B3-F5545D0A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of scenarios.
	SceneList []*DescribeScenesResponseBodySceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
}

func (s DescribeScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScenesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScenesResponseBody) GetSceneList() []*DescribeScenesResponseBodySceneList {
	return s.SceneList
}

func (s *DescribeScenesResponseBody) SetRequestId(v string) *DescribeScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScenesResponseBody) SetSceneList(v []*DescribeScenesResponseBodySceneList) *DescribeScenesResponseBody {
	s.SceneList = v
	return s
}

func (s *DescribeScenesResponseBody) Validate() error {
	if s.SceneList != nil {
		for _, item := range s.SceneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScenesResponseBodySceneList struct {
	// The description of the scenario.
	//
	// example:
	//
	// Query access events for the primary and sub-accounts and access keys under various scenarios, such as access events occurrence, access without MFA authentication, and failed access attempts.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the scenario.
	//
	// example:
	//
	// Account-related or AccessKey Pair-related Events
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the scenario.
	//
	// example:
	//
	// sc-lpYrjKouRfy3MK-wteJW_Q
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The identifier for the scenario category.
	//
	// example:
	//
	// identity
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The type of the scenario.
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeScenesResponseBodySceneList) String() string {
	return dara.Prettify(s)
}

func (s DescribeScenesResponseBodySceneList) GoString() string {
	return s.String()
}

func (s *DescribeScenesResponseBodySceneList) GetDescription() *string {
	return s.Description
}

func (s *DescribeScenesResponseBodySceneList) GetName() *string {
	return s.Name
}

func (s *DescribeScenesResponseBodySceneList) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeScenesResponseBodySceneList) GetToken() *string {
	return s.Token
}

func (s *DescribeScenesResponseBodySceneList) GetType() *string {
	return s.Type
}

func (s *DescribeScenesResponseBodySceneList) SetDescription(v string) *DescribeScenesResponseBodySceneList {
	s.Description = &v
	return s
}

func (s *DescribeScenesResponseBodySceneList) SetName(v string) *DescribeScenesResponseBodySceneList {
	s.Name = &v
	return s
}

func (s *DescribeScenesResponseBodySceneList) SetSceneId(v string) *DescribeScenesResponseBodySceneList {
	s.SceneId = &v
	return s
}

func (s *DescribeScenesResponseBodySceneList) SetToken(v string) *DescribeScenesResponseBodySceneList {
	s.Token = &v
	return s
}

func (s *DescribeScenesResponseBodySceneList) SetType(v string) *DescribeScenesResponseBodySceneList {
	s.Type = &v
	return s
}

func (s *DescribeScenesResponseBodySceneList) Validate() error {
	return dara.Validate(s)
}
