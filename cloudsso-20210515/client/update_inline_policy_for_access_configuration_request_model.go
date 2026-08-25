// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInlinePolicyForAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *UpdateInlinePolicyForAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *UpdateInlinePolicyForAccessConfigurationRequest
	GetDirectoryId() *string
	SetInlinePolicyName(v string) *UpdateInlinePolicyForAccessConfigurationRequest
	GetInlinePolicyName() *string
	SetNewInlinePolicyDocument(v string) *UpdateInlinePolicyForAccessConfigurationRequest
	GetNewInlinePolicyDocument() *string
}

type UpdateInlinePolicyForAccessConfigurationRequest struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the inline policy.
	//
	// example:
	//
	// InlinePolicy
	InlinePolicyName *string `json:"InlinePolicyName,omitempty" xml:"InlinePolicyName,omitempty"`
	// The new configurations of the inline policy.
	//
	// The value can be up to 4,096 characters in length.
	//
	// For more information about the syntax and structure of RAM policies, see [Policy syntax and structure](https://help.aliyun.com/document_detail/93739.html).
	//
	// example:
	//
	// {"Statement": [{"Action": "*","Effect": "Allow","Resource": "*"}],"Version": "1"}
	NewInlinePolicyDocument *string `json:"NewInlinePolicyDocument,omitempty" xml:"NewInlinePolicyDocument,omitempty"`
}

func (s UpdateInlinePolicyForAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInlinePolicyForAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) GetInlinePolicyName() *string {
	return s.InlinePolicyName
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) GetNewInlinePolicyDocument() *string {
	return s.NewInlinePolicyDocument
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetAccessConfigurationId(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetDirectoryId(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetInlinePolicyName(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.InlinePolicyName = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) SetNewInlinePolicyDocument(v string) *UpdateInlinePolicyForAccessConfigurationRequest {
	s.NewInlinePolicyDocument = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
