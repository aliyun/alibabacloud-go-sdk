// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateCompliancePackRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *GetAggregateCompliancePackRequest
	GetCompliancePackId() *string
	SetTag(v []*GetAggregateCompliancePackRequestTag) *GetAggregateCompliancePackRequest
	GetTag() []*GetAggregateCompliancePackRequestTag
}

type GetAggregateCompliancePackRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance pack.
	//
	// For more information about how to obtain the ID of a compliance pack, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// Deprecated
	//
	// The tags of the resource. This parameter is deprecated and no longer takes effect.
	//
	// You can attach up to 20 tags to a resource.
	Tag []*GetAggregateCompliancePackRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetAggregateCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateCompliancePackRequest) GetTag() []*GetAggregateCompliancePackRequestTag {
	return s.Tag
}

func (s *GetAggregateCompliancePackRequest) SetAggregatorId(v string) *GetAggregateCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateCompliancePackRequest) SetCompliancePackId(v string) *GetAggregateCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateCompliancePackRequest) SetTag(v []*GetAggregateCompliancePackRequestTag) *GetAggregateCompliancePackRequest {
	s.Tag = v
	return s
}

func (s *GetAggregateCompliancePackRequest) Validate() error {
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

type GetAggregateCompliancePackRequestTag struct {
	// The tag key of the resource.
	//
	// You can attach up to 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// You can attach up to 20 tag values.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAggregateCompliancePackRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackRequestTag) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetAggregateCompliancePackRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetAggregateCompliancePackRequestTag) SetKey(v string) *GetAggregateCompliancePackRequestTag {
	s.Key = &v
	return s
}

func (s *GetAggregateCompliancePackRequestTag) SetValue(v string) *GetAggregateCompliancePackRequestTag {
	s.Value = &v
	return s
}

func (s *GetAggregateCompliancePackRequestTag) Validate() error {
	return dara.Validate(s)
}
