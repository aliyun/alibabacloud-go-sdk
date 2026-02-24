// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregatorRequest
	GetAggregatorId() *string
	SetTag(v []*GetAggregatorRequestTag) *GetAggregatorRequest
	GetTag() []*GetAggregatorRequestTag
}

type GetAggregatorRequest struct {
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-88ea626622af0055****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// Deprecated
	//
	// The tags of the resource. This parameter is deprecated and is not in use.
	//
	// A maximum of 20 tags can be attached.
	Tag []*GetAggregatorRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetAggregatorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorRequest) GoString() string {
	return s.String()
}

func (s *GetAggregatorRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregatorRequest) GetTag() []*GetAggregatorRequestTag {
	return s.Tag
}

func (s *GetAggregatorRequest) SetAggregatorId(v string) *GetAggregatorRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregatorRequest) SetTag(v []*GetAggregatorRequestTag) *GetAggregatorRequest {
	s.Tag = v
	return s
}

func (s *GetAggregatorRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregatorRequestTag struct {
	// The tag key of the resource.
	//
	// A maximum of 20 tag keys can be attached.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// A maximum of 20 tag values can be attached.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAggregatorRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorRequestTag) GoString() string {
	return s.String()
}

func (s *GetAggregatorRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetAggregatorRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetAggregatorRequestTag) SetKey(v string) *GetAggregatorRequestTag {
	s.Key = &v
	return s
}

func (s *GetAggregatorRequestTag) SetValue(v string) *GetAggregatorRequestTag {
	s.Value = &v
	return s
}

func (s *GetAggregatorRequestTag) Validate() error {
	return dara.Validate(s)
}
