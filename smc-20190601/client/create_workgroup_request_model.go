// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkgroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateWorkgroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateWorkgroupRequest
	GetDescription() *string
	SetName(v string) *CreateWorkgroupRequest
	GetName() *string
	SetOwnerId(v int64) *CreateWorkgroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateWorkgroupRequest
	GetResourceOwnerAccount() *string
	SetTag(v []*CreateWorkgroupRequestTag) *CreateWorkgroupRequest
	GetTag() []*CreateWorkgroupRequestTag
}

type CreateWorkgroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the workgroup. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the workgroup. The name must meet the following requirements:
	//
	// 	- The name must be unique.
	//
	// 	- The name must be 2 to 64 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testWorkgroupName
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tags of the reserved instance. You can specify up to 20 tags. If you specify multiple tags, the tag keys cannot be duplicated.
	Tag []*CreateWorkgroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateWorkgroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkgroupRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkgroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkgroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkgroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkgroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateWorkgroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateWorkgroupRequest) GetTag() []*CreateWorkgroupRequestTag {
	return s.Tag
}

func (s *CreateWorkgroupRequest) SetClientToken(v string) *CreateWorkgroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkgroupRequest) SetDescription(v string) *CreateWorkgroupRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkgroupRequest) SetName(v string) *CreateWorkgroupRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkgroupRequest) SetOwnerId(v int64) *CreateWorkgroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateWorkgroupRequest) SetResourceOwnerAccount(v string) *CreateWorkgroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateWorkgroupRequest) SetTag(v []*CreateWorkgroupRequestTag) *CreateWorkgroupRequest {
	s.Tag = v
	return s
}

func (s *CreateWorkgroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateWorkgroupRequestTag struct {
	// The tag key of the workgroup. You cannot specify an empty string as a tag key. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the workgroup. The tag value can be up to 128 characters in length, cannot be an empty string, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateWorkgroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkgroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateWorkgroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateWorkgroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateWorkgroupRequestTag) SetKey(v string) *CreateWorkgroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateWorkgroupRequestTag) SetValue(v string) *CreateWorkgroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateWorkgroupRequestTag) Validate() error {
	return dara.Validate(s)
}
