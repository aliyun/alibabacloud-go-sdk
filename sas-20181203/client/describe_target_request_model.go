// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeTargetRequest
	GetConfig() *string
	SetType(v string) *DescribeTargetRequest
	GetType() *string
}

type DescribeTargetRequest struct {
	// The type of the vulnerability. Valid values:
	//
	// 	- **cms**: Web CMS vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// example:
	//
	// {"vulType":"cms"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the query. Set the value to vul.
	//
	// example:
	//
	// vul
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTargetRequest) GoString() string {
	return s.String()
}

func (s *DescribeTargetRequest) GetConfig() *string {
	return s.Config
}

func (s *DescribeTargetRequest) GetType() *string {
	return s.Type
}

func (s *DescribeTargetRequest) SetConfig(v string) *DescribeTargetRequest {
	s.Config = &v
	return s
}

func (s *DescribeTargetRequest) SetType(v string) *DescribeTargetRequest {
	s.Type = &v
	return s
}

func (s *DescribeTargetRequest) Validate() error {
	return dara.Validate(s)
}
