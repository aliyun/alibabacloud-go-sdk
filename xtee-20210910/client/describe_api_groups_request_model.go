// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeApiGroupsRequest
	GetLang() *string
	SetApiRegionId(v string) *DescribeApiGroupsRequest
	GetApiRegionId() *string
	SetRegId(v string) *DescribeApiGroupsRequest
	GetRegId() *string
}

type DescribeApiGroupsRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ApiRegionId *string `json:"apiRegionId,omitempty" xml:"apiRegionId,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeApiGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApiGroupsRequest) GetApiRegionId() *string {
	return s.ApiRegionId
}

func (s *DescribeApiGroupsRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeApiGroupsRequest) SetLang(v string) *DescribeApiGroupsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetApiRegionId(v string) *DescribeApiGroupsRequest {
	s.ApiRegionId = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetRegId(v string) *DescribeApiGroupsRequest {
	s.RegId = &v
	return s
}

func (s *DescribeApiGroupsRequest) Validate() error {
	return dara.Validate(s)
}
