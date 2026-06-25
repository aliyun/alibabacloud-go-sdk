// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableMicroRegistration(v bool) *CreateNamespaceRequest
	GetEnableMicroRegistration() *bool
	SetNameSpaceShortId(v string) *CreateNamespaceRequest
	GetNameSpaceShortId() *string
	SetNamespaceDescription(v string) *CreateNamespaceRequest
	GetNamespaceDescription() *string
	SetNamespaceId(v string) *CreateNamespaceRequest
	GetNamespaceId() *string
	SetNamespaceName(v string) *CreateNamespaceRequest
	GetNamespaceName() *string
}

type CreateNamespaceRequest struct {
	// Specifies whether to enable the built-in service registry of SAE.
	//
	// - **true**
	//
	// - **false**
	//
	// The default value is true. If you do not use the built-in service registry, set this parameter to false to speed up namespace creation.
	//
	// example:
	//
	// true
	EnableMicroRegistration *bool `json:"EnableMicroRegistration,omitempty" xml:"EnableMicroRegistration,omitempty"`
	// The short-format namespace ID. You do not need to specify a region ID. This parameter is recommended. The ID cannot exceed 20 characters in length and can contain only lowercase letters and digits.
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
	// The long-format namespace ID. If you specify this parameter, NameSpaceShortId is ignored. This parameter is retained for compatibility. Use the short-format namespace ID instead. The format is `<RegionId>:<NamespaceId>`. The `NamespaceId` can contain only lowercase letters and digits and cannot exceed 32 characters in length. Example: `cn-beijing:test`. For information about the regions that SAE supports, see [DescribeRegions](https://help.aliyun.com/document_detail/126213.html).
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

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetEnableMicroRegistration() *bool {
	return s.EnableMicroRegistration
}

func (s *CreateNamespaceRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *CreateNamespaceRequest) GetNamespaceDescription() *string {
	return s.NamespaceDescription
}

func (s *CreateNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateNamespaceRequest) SetEnableMicroRegistration(v bool) *CreateNamespaceRequest {
	s.EnableMicroRegistration = &v
	return s
}

func (s *CreateNamespaceRequest) SetNameSpaceShortId(v string) *CreateNamespaceRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceDescription(v string) *CreateNamespaceRequest {
	s.NamespaceDescription = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceId(v string) *CreateNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceName(v string) *CreateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
