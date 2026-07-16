// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetaName(v string) *DescribeRuleMetadataRequest
	GetMetaName() *string
}

type DescribeRuleMetadataRequest struct {
	// The name of the metadata.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa_condition_region_config_en
	MetaName *string `json:"MetaName,omitempty" xml:"MetaName,omitempty"`
}

func (s DescribeRuleMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleMetadataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleMetadataRequest) GetMetaName() *string {
	return s.MetaName
}

func (s *DescribeRuleMetadataRequest) SetMetaName(v string) *DescribeRuleMetadataRequest {
	s.MetaName = &v
	return s
}

func (s *DescribeRuleMetadataRequest) Validate() error {
	return dara.Validate(s)
}
