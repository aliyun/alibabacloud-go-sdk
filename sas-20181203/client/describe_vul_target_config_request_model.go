// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulTargetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeVulTargetConfigRequest
	GetType() *string
	SetUuid(v string) *DescribeVulTargetConfigRequest
	GetUuid() *string
}

type DescribeVulTargetConfigRequest struct {
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// ae1527a9-2308-46ab-b10a-48ae7ff7****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeVulTargetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetConfigRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVulTargetConfigRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeVulTargetConfigRequest) SetType(v string) *DescribeVulTargetConfigRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulTargetConfigRequest) SetUuid(v string) *DescribeVulTargetConfigRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeVulTargetConfigRequest) Validate() error {
	return dara.Validate(s)
}
