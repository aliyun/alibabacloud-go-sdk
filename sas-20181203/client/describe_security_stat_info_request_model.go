// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityStatInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSecurityStatInfoRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v string) *DescribeSecurityStatInfoRequest
	GetResourceDirectoryAccountId() *string
	SetSourceIp(v string) *DescribeSecurityStatInfoRequest
	GetSourceIp() *string
}

type DescribeSecurityStatInfoRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the account that is added to the resource directory as a member for multi-account control. You can use this parameter to query the security status of the account.
	//
	// example:
	//
	// 12345
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSecurityStatInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSecurityStatInfoRequest) GetResourceDirectoryAccountId() *string {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeSecurityStatInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSecurityStatInfoRequest) SetLang(v string) *DescribeSecurityStatInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecurityStatInfoRequest) SetResourceDirectoryAccountId(v string) *DescribeSecurityStatInfoRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeSecurityStatInfoRequest) SetSourceIp(v string) *DescribeSecurityStatInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityStatInfoRequest) Validate() error {
	return dara.Validate(s)
}
