// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeVulDetailsRequest
	GetAliasName() *string
	SetLang(v string) *DescribeVulDetailsRequest
	GetLang() *string
	SetName(v string) *DescribeVulDetailsRequest
	GetName() *string
	SetResourceDirectoryAccountId(v int64) *DescribeVulDetailsRequest
	GetResourceDirectoryAccountId() *int64
	SetType(v string) *DescribeVulDetailsRequest
	GetType() *string
}

type DescribeVulDetailsRequest struct {
	// The vulnerability announcement.
	//
	// example:
	//
	// RHSA-2019:3197-Important: sudo security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the vulnerability.
	//
	// > You can call the [DescribeGroupedVul](~~DescribeGroupedVul~~) or [DescribeVulList](~~DescribeVulList~~) operation to query the names of vulnerabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// SCA:ACSV-2020-052801
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **app**: application vulnerabilitiy
	//
	// 	- **emg**: urgent vulnerability
	//
	// 	- **sca**: vulnerability that is detected based on software component analysis
	//
	// This parameter is required.
	//
	// example:
	//
	// sca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeVulDetailsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVulDetailsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeVulDetailsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeVulDetailsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVulDetailsRequest) SetAliasName(v string) *DescribeVulDetailsRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetLang(v string) *DescribeVulDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetName(v string) *DescribeVulDetailsRequest {
	s.Name = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetResourceDirectoryAccountId(v int64) *DescribeVulDetailsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetType(v string) *DescribeVulDetailsRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulDetailsRequest) Validate() error {
	return dara.Validate(s)
}
