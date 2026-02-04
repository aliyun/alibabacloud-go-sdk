// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCenRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCenRequest
	GetDescription() *string
	SetName(v string) *CreateCenRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCenRequest
	GetOwnerId() *int64
	SetProtectionLevel(v string) *CreateCenRequest
	GetProtectionLevel() *string
	SetResourceOwnerAccount(v string) *CreateCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCenRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateCenRequestTag) *CreateCenRequest
	GetTag() []*CreateCenRequestTag
}

type CreateCenRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the CEN instance.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the CEN instance.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testname
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The level of CIDR block overlapping.
	//
	// Set the value to **REDUCED*	- (default). This value specifies that CIDR blocks can overlap but cannot be the same.
	//
	// example:
	//
	// REDUCED
	ProtectionLevel      *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	//
	// You can specify up to 20 tags in each call.
	Tag []*CreateCenRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateCenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenRequest) GoString() string {
	return s.String()
}

func (s *CreateCenRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCenRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCenRequest) GetName() *string {
	return s.Name
}

func (s *CreateCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCenRequest) GetProtectionLevel() *string {
	return s.ProtectionLevel
}

func (s *CreateCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCenRequest) GetTag() []*CreateCenRequestTag {
	return s.Tag
}

func (s *CreateCenRequest) SetClientToken(v string) *CreateCenRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenRequest) SetDescription(v string) *CreateCenRequest {
	s.Description = &v
	return s
}

func (s *CreateCenRequest) SetName(v string) *CreateCenRequest {
	s.Name = &v
	return s
}

func (s *CreateCenRequest) SetOwnerAccount(v string) *CreateCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenRequest) SetOwnerId(v int64) *CreateCenRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenRequest) SetProtectionLevel(v string) *CreateCenRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateCenRequest) SetResourceOwnerAccount(v string) *CreateCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenRequest) SetResourceOwnerId(v int64) *CreateCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenRequest) SetTag(v []*CreateCenRequestTag) *CreateCenRequest {
	s.Tag = v
	return s
}

func (s *CreateCenRequest) Validate() error {
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

type CreateCenRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify up to 20 tag values in each call.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCenRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCenRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCenRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCenRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCenRequestTag) SetKey(v string) *CreateCenRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCenRequestTag) SetValue(v string) *CreateCenRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCenRequestTag) Validate() error {
	return dara.Validate(s)
}
