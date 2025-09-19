// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceDirectoryAccountId(v string) *DescribeVersionConfigRequest
	GetResourceDirectoryAccountId() *string
	SetSourceIp(v string) *DescribeVersionConfigRequest
	GetSourceIp() *string
}

type DescribeVersionConfigRequest struct {
	// The ID of the Alibaba Cloud account that uses Security Center.
	//
	// >  You can call the [GetUser](https://help.aliyun.com/document_detail/28681.html) operation to query the IDs of Alibaba Cloud accounts.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeVersionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigRequest) GetResourceDirectoryAccountId() *string {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeVersionConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeVersionConfigRequest) SetResourceDirectoryAccountId(v string) *DescribeVersionConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeVersionConfigRequest) SetSourceIp(v string) *DescribeVersionConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVersionConfigRequest) Validate() error {
	return dara.Validate(s)
}
