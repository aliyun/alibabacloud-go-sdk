// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProcessDefinitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListProcessDefinitionsResponseBodyPagingInfo) *ListProcessDefinitionsResponseBody
	GetPagingInfo() *ListProcessDefinitionsResponseBodyPagingInfo
	SetRequestId(v string) *ListProcessDefinitionsResponseBody
	GetRequestId() *string
}

type ListProcessDefinitionsResponseBody struct {
	PagingInfo *ListProcessDefinitionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProcessDefinitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProcessDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProcessDefinitionsResponseBody) GetPagingInfo() *ListProcessDefinitionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListProcessDefinitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProcessDefinitionsResponseBody) SetPagingInfo(v *ListProcessDefinitionsResponseBodyPagingInfo) *ListProcessDefinitionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProcessDefinitionsResponseBody) SetRequestId(v string) *ListProcessDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProcessDefinitionsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProcessDefinitionsResponseBodyPagingInfo struct {
	ProcessDefinitions []*ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions `json:"ProcessDefinitions,omitempty" xml:"ProcessDefinitions,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProcessDefinitionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListProcessDefinitionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProcessDefinitionsResponseBodyPagingInfo) GetProcessDefinitions() []*ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	return s.ProcessDefinitions
}

func (s *ListProcessDefinitionsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProcessDefinitionsResponseBodyPagingInfo) SetProcessDefinitions(v []*ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) *ListProcessDefinitionsResponseBodyPagingInfo {
	s.ProcessDefinitions = v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListProcessDefinitionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfo) Validate() error {
	if s.ProcessDefinitions != nil {
		for _, item := range s.ProcessDefinitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 10354346
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsSystem *bool   `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Priority *string   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Scopes   []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// example:
	//
	// Table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) String() string {
	return dara.Prettify(s)
}

func (s ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GoString() string {
	return s.String()
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetDescription() *string {
	return s.Description
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetId() *string {
	return s.Id
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetIsSystem() *bool {
	return s.IsSystem
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetName() *string {
	return s.Name
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetPriority() *string {
	return s.Priority
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetScopes() []*string {
	return s.Scopes
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetSubType() *string {
	return s.SubType
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) GetType() *string {
	return s.Type
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetDescription(v string) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.Description = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetEnabled(v bool) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.Enabled = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetId(v string) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.Id = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetIsSystem(v bool) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.IsSystem = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetName(v string) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.Name = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetPriority(v string) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.Priority = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetScopes(v []*string) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.Scopes = v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetSubType(v string) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.SubType = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) SetType(v string) *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions {
	s.Type = &v
	return s
}

func (s *ListProcessDefinitionsResponseBodyPagingInfoProcessDefinitions) Validate() error {
	return dara.Validate(s)
}
