// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeVulConfigRequest
	GetSourceIp() *string
	SetType(v string) *DescribeVulConfigRequest
	GetType() *string
}

type DescribeVulConfigRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 113.110.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of configuration. By default, all types of configurations are queried. Valid values:
	//
	// 	- **cve**: Linux software vulnerability.
	//
	// 	- **sys**: Windows system vulnerability.
	//
	// 	- **cms**: Web-CMS vulnerability.
	//
	// 	- **app**: application vulnerability that is detected by using web scanner.
	//
	// 	- **emg**: urgent vulnerability.
	//
	// 	- **scanMode**: displays easily exploitable vulnerability.
	//
	// 	- **imageVulClean**: vulnerability retention duration.
	//
	// 	- **yum**: preferentially uses YUM or APT sources of Alibaba Cloud to fix vulnerabilities.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeVulConfigRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVulConfigRequest) SetSourceIp(v string) *DescribeVulConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVulConfigRequest) SetType(v string) *DescribeVulConfigRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulConfigRequest) Validate() error {
	return dara.Validate(s)
}
