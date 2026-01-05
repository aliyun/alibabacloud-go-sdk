// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v []*ListComponentsResponseBodyComponents) *ListComponentsResponseBody
	GetComponents() []*ListComponentsResponseBodyComponents
	SetRequestId(v string) *ListComponentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListComponentsResponseBody
	GetTotalCount() *int32
}

type ListComponentsResponseBody struct {
	Components []*ListComponentsResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) GetComponents() []*ListComponentsResponseBodyComponents {
	return s.Components
}

func (s *ListComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComponentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListComponentsResponseBody) SetComponents(v []*ListComponentsResponseBodyComponents) *ListComponentsResponseBody {
	s.Components = v
	return s
}

func (s *ListComponentsResponseBody) SetRequestId(v string) *ListComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentsResponseBody) SetTotalCount(v int32) *ListComponentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListComponentsResponseBody) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComponentsResponseBodyComponents struct {
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// dataset-accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId  *string                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Template *ListComponentsResponseBodyComponentsTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// cmpt-5zk866779me51jgu3w
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// v1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListComponentsResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponents) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListComponentsResponseBodyComponents) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListComponentsResponseBodyComponents) GetName() *string {
	return s.Name
}

func (s *ListComponentsResponseBodyComponents) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListComponentsResponseBodyComponents) GetTemplate() *ListComponentsResponseBodyComponentsTemplate {
	return s.Template
}

func (s *ListComponentsResponseBodyComponents) GetUserId() *string {
	return s.UserId
}

func (s *ListComponentsResponseBodyComponents) GetUuid() *string {
	return s.Uuid
}

func (s *ListComponentsResponseBodyComponents) GetVersion() *string {
	return s.Version
}

func (s *ListComponentsResponseBodyComponents) SetGmtCreateTime(v string) *ListComponentsResponseBodyComponents {
	s.GmtCreateTime = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetGmtModifiedTime(v string) *ListComponentsResponseBodyComponents {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetName(v string) *ListComponentsResponseBodyComponents {
	s.Name = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetOwnerId(v string) *ListComponentsResponseBodyComponents {
	s.OwnerId = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetTemplate(v *ListComponentsResponseBodyComponentsTemplate) *ListComponentsResponseBodyComponents {
	s.Template = v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetUserId(v string) *ListComponentsResponseBodyComponents {
	s.UserId = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetUuid(v string) *ListComponentsResponseBodyComponents {
	s.Uuid = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetVersion(v string) *ListComponentsResponseBodyComponents {
	s.Version = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComponentsResponseBodyComponentsTemplate struct {
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// data/VOCdevkit/VOC2007/ImageSets/Main/val.txt
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListComponentsResponseBodyComponentsTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsResponseBodyComponentsTemplate) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponentsTemplate) GetType() *string {
	return s.Type
}

func (s *ListComponentsResponseBodyComponentsTemplate) GetUri() *string {
	return s.Uri
}

func (s *ListComponentsResponseBodyComponentsTemplate) SetType(v string) *ListComponentsResponseBodyComponentsTemplate {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyComponentsTemplate) SetUri(v string) *ListComponentsResponseBodyComponentsTemplate {
	s.Uri = &v
	return s
}

func (s *ListComponentsResponseBodyComponentsTemplate) Validate() error {
	return dara.Validate(s)
}
