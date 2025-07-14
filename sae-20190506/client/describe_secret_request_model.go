// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DescribeSecretRequest
	GetNamespaceId() *string
	SetSecretId(v int64) *DescribeSecretRequest
	GetSecretId() *int64
}

type DescribeSecretRequest struct {
	// The ID of the namespace where the Secret resides. If the namespace is the default namespace, you need to only enter the region ID, such as `cn-beijing`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the Secret instance to be queried. You can call the [ListSecrets](https://help.aliyun.com/document_detail/466929.html) operation to view the IDs of Secrete instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DescribeSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecretRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeSecretRequest) GetSecretId() *int64 {
	return s.SecretId
}

func (s *DescribeSecretRequest) SetNamespaceId(v string) *DescribeSecretRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeSecretRequest) SetSecretId(v int64) *DescribeSecretRequest {
	s.SecretId = &v
	return s
}

func (s *DescribeSecretRequest) Validate() error {
	return dara.Validate(s)
}
