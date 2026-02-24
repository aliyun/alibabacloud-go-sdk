// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *GetCompliancePackRequest
	GetCompliancePackId() *string
	SetTag(v []*GetCompliancePackRequestTag) *GetCompliancePackRequest
	GetTag() []*GetCompliancePackRequestTag
}

type GetCompliancePackRequest struct {
	// The compliance package ID.
	//
	// For more information about how to obtain the compliance package ID, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// Deprecated
	//
	// The tags of the resource. This parameter is deprecated and takes no effect if it is specified.
	//
	// You can add up to 20 tags.
	Tag []*GetCompliancePackRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *GetCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetCompliancePackRequest) GetTag() []*GetCompliancePackRequestTag {
	return s.Tag
}

func (s *GetCompliancePackRequest) SetCompliancePackId(v string) *GetCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetCompliancePackRequest) SetTag(v []*GetCompliancePackRequestTag) *GetCompliancePackRequest {
	s.Tag = v
	return s
}

func (s *GetCompliancePackRequest) Validate() error {
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

type GetCompliancePackRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value can be an empty string or a string of up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each tag key must have a corresponding tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCompliancePackRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackRequestTag) GoString() string {
	return s.String()
}

func (s *GetCompliancePackRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetCompliancePackRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetCompliancePackRequestTag) SetKey(v string) *GetCompliancePackRequestTag {
	s.Key = &v
	return s
}

func (s *GetCompliancePackRequestTag) SetValue(v string) *GetCompliancePackRequestTag {
	s.Value = &v
	return s
}

func (s *GetCompliancePackRequestTag) Validate() error {
	return dara.Validate(s)
}
