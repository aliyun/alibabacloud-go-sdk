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
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
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
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be an empty string or up to 128 characters in length. The tag values cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
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
