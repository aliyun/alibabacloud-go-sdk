// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityOverallDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeQualityOverallDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeQualityOverallDataRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeQualityOverallDataRequest
	GetStartDate() *int64
	SetTypes(v string) *DescribeQualityOverallDataRequest
	GetTypes() *string
}

type DescribeQualityOverallDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JOIN_CHANNEL_SUC_RATE
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeQualityOverallDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeQualityOverallDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeQualityOverallDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeQualityOverallDataRequest) GetTypes() *string {
	return s.Types
}

func (s *DescribeQualityOverallDataRequest) SetAppId(v string) *DescribeQualityOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetEndDate(v int64) *DescribeQualityOverallDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetStartDate(v int64) *DescribeQualityOverallDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetTypes(v string) *DescribeQualityOverallDataRequest {
	s.Types = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) Validate() error {
	return dara.Validate(s)
}
