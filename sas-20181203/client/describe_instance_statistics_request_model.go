// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeInstanceStatisticsRequest
	GetFrom() *string
	SetLang(v string) *DescribeInstanceStatisticsRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeInstanceStatisticsRequest
	GetSourceIp() *string
	SetUuid(v string) *DescribeInstanceStatisticsRequest
	GetUuid() *string
}

type DescribeInstanceStatisticsRequest struct {
	// The source of the request. Set the value to **sas**, which indicates that the request is sent from Security Center.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
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
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUIDs of the assets to query. Separate multiple UUIDs with commas (,). You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of assets.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeInstanceStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInstanceStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstanceStatisticsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstanceStatisticsRequest) SetFrom(v string) *DescribeInstanceStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetLang(v string) *DescribeInstanceStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetUuid(v string) *DescribeInstanceStatisticsRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
