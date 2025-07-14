// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableMicroRegistration(v bool) *UpdateNamespaceRequest
	GetEnableMicroRegistration() *bool
	SetNameSpaceShortId(v string) *UpdateNamespaceRequest
	GetNameSpaceShortId() *string
	SetNamespaceDescription(v string) *UpdateNamespaceRequest
	GetNamespaceDescription() *string
	SetNamespaceId(v string) *UpdateNamespaceRequest
	GetNamespaceId() *string
	SetNamespaceName(v string) *UpdateNamespaceRequest
	GetNamespaceName() *string
}

type UpdateNamespaceRequest struct {
	// Indicates whether to enable SAE built-in registry:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// If you set this parameter to true, a shared registry is created for the namespace. The registry cannot be disabled after it is created.
	//
	// example:
	//
	// true
	EnableMicroRegistration *bool `json:"EnableMicroRegistration,omitempty" xml:"EnableMicroRegistration,omitempty"`
	// The short ID of the namespace. You do not need to specify a region ID. We recommend that you configure this parameter. The value of this parameter can be up to 20 characters in length and can contain only lowercase letters and digits.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The description of the namespace. The description cannot exceed 100 characters in length.
	//
	// example:
	//
	// desc
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	// The long ID of the namespace. If you configure this parameter, the long ID take effects and the value of the NameSpaceShortId parameter is ignored. To ensure compatibility, we recommend that you specify a short namespace ID. A long namespace ID follows the `<RegionId>:<NamespaceId>` format. The `NamespaceId` variable can contain only lowercase letters and digits. Example: `cn-beijing:test`. The value of the Namespaceid variable cannot exceed 32 characters in length. For more information about **RegionId**, you can call the [DescribeRegions](https://help.aliyun.com/document_detail/2834842.html) operation to obtain the IDs of regions supported by SAE.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace. The name cannot exceed 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) GetEnableMicroRegistration() *bool {
	return s.EnableMicroRegistration
}

func (s *UpdateNamespaceRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *UpdateNamespaceRequest) GetNamespaceDescription() *string {
	return s.NamespaceDescription
}

func (s *UpdateNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateNamespaceRequest) SetEnableMicroRegistration(v bool) *UpdateNamespaceRequest {
	s.EnableMicroRegistration = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNameSpaceShortId(v string) *UpdateNamespaceRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceDescription(v string) *UpdateNamespaceRequest {
	s.NamespaceDescription = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceId(v string) *UpdateNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceName(v string) *UpdateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
