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
	// Specifies whether to enable the built-in registry of SAE. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// If you set this parameter to true, a shared registry and related information are created for the namespace. After the shared registry is created, it cannot be disabled.
	//
	// example:
	//
	// true
	EnableMicroRegistration *bool `json:"EnableMicroRegistration,omitempty" xml:"EnableMicroRegistration,omitempty"`
	// The short namespace ID. You do not need to include the region ID. The value can contain only lowercase letters and digits and cannot exceed 20 characters in length. We recommend that you use this parameter.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The namespace description. The value cannot exceed 100 characters in length.
	//
	// example:
	//
	// desc
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	// The long namespace ID. If you specify this parameter, it takes effect and NameSpaceShortId is ignored. This parameter is retained for backward compatibility. We recommend that you use the short namespace ID. The format is `<RegionId>:<NamespaceId>`. `NamespaceId` can contain only lowercase letters and digits, such as `cn-beijing:test`. The value cannot exceed 32 characters in length. For more information about the **RegionId*	- values supported by SAE, see [DescribeRegions](https://help.aliyun.com/document_detail/2834842.html).
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The namespace name. The value cannot exceed 64 characters in length.
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
