// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DescribeKeyRequest
	GetKeyId() *string
}

type DescribeKeyRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 05754286-3ba2-4fa6-8d41-4323aca6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DescribeKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeKeyRequest) SetKeyId(v string) *DescribeKeyRequest {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyRequest) Validate() error {
	return dara.Validate(s)
}
