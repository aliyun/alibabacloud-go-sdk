// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonOverallConfigListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeCommonOverallConfigListRequest
	GetSourceIp() *string
	SetTypeList(v []*string) *DescribeCommonOverallConfigListRequest
	GetTypeList() []*string
}

type DescribeCommonOverallConfigListRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 119.136.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The types of the configuration items.
	//
	// >  You can query up to 50 types at a time.
	//
	// This parameter is required.
	TypeList []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribeCommonOverallConfigListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonOverallConfigListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommonOverallConfigListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCommonOverallConfigListRequest) GetTypeList() []*string {
	return s.TypeList
}

func (s *DescribeCommonOverallConfigListRequest) SetSourceIp(v string) *DescribeCommonOverallConfigListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCommonOverallConfigListRequest) SetTypeList(v []*string) *DescribeCommonOverallConfigListRequest {
	s.TypeList = v
	return s
}

func (s *DescribeCommonOverallConfigListRequest) Validate() error {
	return dara.Validate(s)
}
