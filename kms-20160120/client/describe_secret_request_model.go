// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFetchTags(v string) *DescribeSecretRequest
	GetFetchTags() *string
	SetSecretName(v string) *DescribeSecretRequest
	GetSecretName() *string
}

type DescribeSecretRequest struct {
	// Specifies whether to return the resource tags of the secret. Valid values:
	//
	// 	- true: The resource tags are returned.
	//
	// 	- false: The resource tags are not returned. This is the default value.
	//
	// example:
	//
	// true
	FetchTags *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DescribeSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecretRequest) GetFetchTags() *string {
	return s.FetchTags
}

func (s *DescribeSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeSecretRequest) SetFetchTags(v string) *DescribeSecretRequest {
	s.FetchTags = &v
	return s
}

func (s *DescribeSecretRequest) SetSecretName(v string) *DescribeSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretRequest) Validate() error {
	return dara.Validate(s)
}
