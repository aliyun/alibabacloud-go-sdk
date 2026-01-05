// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclName(v string) *CreateAclRequest
	GetAclName() *string
	SetClientToken(v string) *CreateAclRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateAclRequest
	GetDryRun() *bool
	SetResourceGroupId(v string) *CreateAclRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateAclRequestTag) *CreateAclRequest
	GetTag() []*CreateAclRequestTag
}

type CreateAclRequest struct {
	// The name of the ACL. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test-acl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*CreateAclRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateAclRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) GetAclName() *string {
	return s.AclName
}

func (s *CreateAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAclRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAclRequest) GetTag() []*CreateAclRequestTag {
	return s.Tag
}

func (s *CreateAclRequest) SetAclName(v string) *CreateAclRequest {
	s.AclName = &v
	return s
}

func (s *CreateAclRequest) SetClientToken(v string) *CreateAclRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAclRequest) SetDryRun(v bool) *CreateAclRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAclRequest) SetResourceGroupId(v string) *CreateAclRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAclRequest) SetTag(v []*CreateAclRequestTag) *CreateAclRequest {
	s.Tag = v
	return s
}

func (s *CreateAclRequest) Validate() error {
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

type CreateAclRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAclRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAclRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAclRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAclRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAclRequestTag) SetKey(v string) *CreateAclRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAclRequestTag) SetValue(v string) *CreateAclRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAclRequestTag) Validate() error {
	return dara.Validate(s)
}
