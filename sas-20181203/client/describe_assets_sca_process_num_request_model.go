// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetsScaProcessNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeAssetsScaProcessNumRequest
	GetBizType() *string
	SetUuidList(v []*string) *DescribeAssetsScaProcessNumRequest
	GetUuidList() []*string
}

type DescribeAssetsScaProcessNumRequest struct {
	// The type of the application process. Default value: java. Valid values:
	//
	// 	- **java**
	//
	// 	- **php**
	//
	// example:
	//
	// java
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The UUIDs of the servers.
	//
	// This parameter is required.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s DescribeAssetsScaProcessNumRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsScaProcessNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetsScaProcessNumRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeAssetsScaProcessNumRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *DescribeAssetsScaProcessNumRequest) SetBizType(v string) *DescribeAssetsScaProcessNumRequest {
	s.BizType = &v
	return s
}

func (s *DescribeAssetsScaProcessNumRequest) SetUuidList(v []*string) *DescribeAssetsScaProcessNumRequest {
	s.UuidList = v
	return s
}

func (s *DescribeAssetsScaProcessNumRequest) Validate() error {
	return dara.Validate(s)
}
