// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientConfStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v string) *ModifyClientConfStrategyRequest
	GetTag() *string
	SetTagExt(v string) *ModifyClientConfStrategyRequest
	GetTagExt() *string
	SetTagValue(v string) *ModifyClientConfStrategyRequest
	GetTagValue() *string
	SetUuid(v string) *ModifyClientConfStrategyRequest
	GetUuid() *string
	SetUuids(v []*string) *ModifyClientConfStrategyRequest
	GetUuids() []*string
}

type ModifyClientConfStrategyRequest struct {
	// The key of the tag that is added to the agent configuration policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// machineResource
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The extended tag of the agent configuration policy.
	//
	// example:
	//
	// auto
	TagExt *string `json:"TagExt,omitempty" xml:"TagExt,omitempty"`
	// The value of the tag that is added to the agent configuration policy.
	//
	// 	- major
	//
	// 	- advanced
	//
	// 	- basic
	//
	// This parameter is required.
	//
	// example:
	//
	// advanced
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The UUID of the server that you want to query.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The UUID of the asset. You can specify a maximum of 500 UUIDs at a time.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s ModifyClientConfStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientConfStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyClientConfStrategyRequest) GetTag() *string {
	return s.Tag
}

func (s *ModifyClientConfStrategyRequest) GetTagExt() *string {
	return s.TagExt
}

func (s *ModifyClientConfStrategyRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *ModifyClientConfStrategyRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyClientConfStrategyRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *ModifyClientConfStrategyRequest) SetTag(v string) *ModifyClientConfStrategyRequest {
	s.Tag = &v
	return s
}

func (s *ModifyClientConfStrategyRequest) SetTagExt(v string) *ModifyClientConfStrategyRequest {
	s.TagExt = &v
	return s
}

func (s *ModifyClientConfStrategyRequest) SetTagValue(v string) *ModifyClientConfStrategyRequest {
	s.TagValue = &v
	return s
}

func (s *ModifyClientConfStrategyRequest) SetUuid(v string) *ModifyClientConfStrategyRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyClientConfStrategyRequest) SetUuids(v []*string) *ModifyClientConfStrategyRequest {
	s.Uuids = v
	return s
}

func (s *ModifyClientConfStrategyRequest) Validate() error {
	return dara.Validate(s)
}
