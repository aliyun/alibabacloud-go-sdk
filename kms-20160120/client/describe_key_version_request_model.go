// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DescribeKeyVersionRequest
	GetKeyId() *string
	SetKeyVersionId(v string) *DescribeKeyVersionRequest
	GetKeyVersionId() *string
}

type DescribeKeyVersionRequest struct {
	// The globally unique ID of the CMK.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// You can call the [ListKeyVersions](https://help.aliyun.com/document_detail/133966.html) operation to query the versions of the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s DescribeKeyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeKeyVersionRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *DescribeKeyVersionRequest) SetKeyId(v string) *DescribeKeyVersionRequest {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyVersionRequest) SetKeyVersionId(v string) *DescribeKeyVersionRequest {
	s.KeyVersionId = &v
	return s
}

func (s *DescribeKeyVersionRequest) Validate() error {
	return dara.Validate(s)
}
