// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMobileAgentPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPackageId(v string) *ModifyMobileAgentPackageRequest
	GetPackageId() *string
	SetTag(v []*ModifyMobileAgentPackageRequestTag) *ModifyMobileAgentPackageRequest
	GetTag() []*ModifyMobileAgentPackageRequestTag
}

type ModifyMobileAgentPackageRequest struct {
	// The package ID.
	//
	// example:
	//
	// cmag-xxxx
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The list of tags. You can specify 1 to 20 tags.
	Tag []*ModifyMobileAgentPackageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ModifyMobileAgentPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMobileAgentPackageRequest) GoString() string {
	return s.String()
}

func (s *ModifyMobileAgentPackageRequest) GetPackageId() *string {
	return s.PackageId
}

func (s *ModifyMobileAgentPackageRequest) GetTag() []*ModifyMobileAgentPackageRequestTag {
	return s.Tag
}

func (s *ModifyMobileAgentPackageRequest) SetPackageId(v string) *ModifyMobileAgentPackageRequest {
	s.PackageId = &v
	return s
}

func (s *ModifyMobileAgentPackageRequest) SetTag(v []*ModifyMobileAgentPackageRequestTag) *ModifyMobileAgentPackageRequest {
	s.Tag = v
	return s
}

func (s *ModifyMobileAgentPackageRequest) Validate() error {
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

type ModifyMobileAgentPackageRequestTag struct {
	// The tag key. If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// keyname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. It can be an empty string. The tag value can be up to 128 characters in length and cannot start with acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyMobileAgentPackageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ModifyMobileAgentPackageRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyMobileAgentPackageRequestTag) GetKey() *string {
	return s.Key
}

func (s *ModifyMobileAgentPackageRequestTag) GetValue() *string {
	return s.Value
}

func (s *ModifyMobileAgentPackageRequestTag) SetKey(v string) *ModifyMobileAgentPackageRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyMobileAgentPackageRequestTag) SetValue(v string) *ModifyMobileAgentPackageRequestTag {
	s.Value = &v
	return s
}

func (s *ModifyMobileAgentPackageRequestTag) Validate() error {
	return dara.Validate(s)
}
