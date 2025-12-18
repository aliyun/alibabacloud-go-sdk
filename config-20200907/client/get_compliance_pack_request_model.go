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
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
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
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
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
