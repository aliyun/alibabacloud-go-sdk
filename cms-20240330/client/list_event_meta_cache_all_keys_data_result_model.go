// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventMetaCacheAllKeysDataResult interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotation(v []*string) *ListEventMetaCacheAllKeysDataResult
	GetAnnotation() []*string
	SetLabel(v []*string) *ListEventMetaCacheAllKeysDataResult
	GetLabel() []*string
	SetResourceTag(v []*string) *ListEventMetaCacheAllKeysDataResult
	GetResourceTag() []*string
}

type ListEventMetaCacheAllKeysDataResult struct {
	// annotation类型的Key列表
	//
	// example:
	//
	// ["message","current_value"]
	Annotation []*string `json:"annotation,omitempty" xml:"annotation,omitempty" type:"Repeated"`
	// label类型的Key列表
	//
	// example:
	//
	// ["_cms_rule_id","_cms_rule_name"]
	Label []*string `json:"label,omitempty" xml:"label,omitempty" type:"Repeated"`
	// resource.tag类型的Key列表
	//
	// example:
	//
	// ["arn","callType"]
	ResourceTag []*string `json:"resourceTag,omitempty" xml:"resourceTag,omitempty" type:"Repeated"`
}

func (s ListEventMetaCacheAllKeysDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListEventMetaCacheAllKeysDataResult) GoString() string {
	return s.String()
}

func (s *ListEventMetaCacheAllKeysDataResult) GetAnnotation() []*string {
	return s.Annotation
}

func (s *ListEventMetaCacheAllKeysDataResult) GetLabel() []*string {
	return s.Label
}

func (s *ListEventMetaCacheAllKeysDataResult) GetResourceTag() []*string {
	return s.ResourceTag
}

func (s *ListEventMetaCacheAllKeysDataResult) SetAnnotation(v []*string) *ListEventMetaCacheAllKeysDataResult {
	s.Annotation = v
	return s
}

func (s *ListEventMetaCacheAllKeysDataResult) SetLabel(v []*string) *ListEventMetaCacheAllKeysDataResult {
	s.Label = v
	return s
}

func (s *ListEventMetaCacheAllKeysDataResult) SetResourceTag(v []*string) *ListEventMetaCacheAllKeysDataResult {
	s.ResourceTag = v
	return s
}

func (s *ListEventMetaCacheAllKeysDataResult) Validate() error {
	return dara.Validate(s)
}
