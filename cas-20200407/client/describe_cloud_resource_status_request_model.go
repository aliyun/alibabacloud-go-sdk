// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecretId(v string) *DescribeCloudResourceStatusRequest
	GetSecretId() *string
}

type DescribeCloudResourceStatusRequest struct {
	// The AccessKey secret used to access cloud resources.
	//
	// >  You can call the [ListCloudAccess](https://help.aliyun.com/document_detail/2712219.html) operation to obtain the ID.
	//
	// example:
	//
	// AKID9*******XX
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DescribeCloudResourceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *DescribeCloudResourceStatusRequest) SetSecretId(v string) *DescribeCloudResourceStatusRequest {
	s.SecretId = &v
	return s
}

func (s *DescribeCloudResourceStatusRequest) Validate() error {
	return dara.Validate(s)
}
