// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserWafLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogRegionId(v string) *DescribeUserWafLogStatusResponseBody
	GetLogRegionId() *string
	SetLogStatus(v string) *DescribeUserWafLogStatusResponseBody
	GetLogStatus() *string
	SetRequestId(v string) *DescribeUserWafLogStatusResponseBody
	GetRequestId() *string
	SetStatusUpdateTime(v int64) *DescribeUserWafLogStatusResponseBody
	GetStatusUpdateTime() *int64
}

type DescribeUserWafLogStatusResponseBody struct {
	// The ID of the region where WAF logs are stored. Valid values:
	//
	// 	- **cn-hangzhou**: China (Hangzhou).
	//
	// 	- **cn-beijing**: China (Beijing).
	//
	// 	- **cn-hongkong**: China (Hong Kong).
	//
	// 	- **ap-southeast-1**: Singapore.
	//
	// 	- **ap-southeast-3**: Malaysia (Kuala Lumpur).
	//
	// 	- **ap-southeast-5**: Indonesia (Jakarta).
	//
	// 	- **ap-southeast-6**: Philippines (Manila).
	//
	// 	- **ap-southeast-7**: Thailand (Bangkok).
	//
	// 	- **me-east-1**: UAE (Dubai).
	//
	// 	- **eu-central-1**: Germany (Frankfurt).
	//
	// 	- **us-east-1**: US (Virginia).
	//
	// 	- **us-west-1**: US (Silicon Valley).
	//
	// 	- **ap-northeast-1**: Japan (Tokyo).
	//
	// 	- **ap-northeast-2**: South Korea (Seoul).
	//
	// 	- **eu-west-1**: UK (London).
	//
	// 	- **cn-hangzhou-finance**: China East 1 Finance.
	//
	// 	- **cn-shanghai-finance-1**: China East 2 Finance.
	//
	// 	- **cn-shenzhen-finance**: China South 1 Finance.
	//
	// >  The China East 1 Finance, China East 2 Finance, and China South 1 Finance regions are available only for Alibaba Finance Cloud users. Alibaba Finance Cloud users are also limited to storing logs within these specific regions.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The status of WAF logs.
	//
	// 	- **initializing**
	//
	// 	- **initialize_failed**
	//
	// 	- **normal**
	//
	// 	- **releasing**
	//
	// 	- **release_failed**
	//
	// example:
	//
	// normal
	LogStatus *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the log status was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1706771796859
	StatusUpdateTime *int64 `json:"StatusUpdateTime,omitempty" xml:"StatusUpdateTime,omitempty"`
}

func (s DescribeUserWafLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserWafLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserWafLogStatusResponseBody) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *DescribeUserWafLogStatusResponseBody) GetLogStatus() *string {
	return s.LogStatus
}

func (s *DescribeUserWafLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserWafLogStatusResponseBody) GetStatusUpdateTime() *int64 {
	return s.StatusUpdateTime
}

func (s *DescribeUserWafLogStatusResponseBody) SetLogRegionId(v string) *DescribeUserWafLogStatusResponseBody {
	s.LogRegionId = &v
	return s
}

func (s *DescribeUserWafLogStatusResponseBody) SetLogStatus(v string) *DescribeUserWafLogStatusResponseBody {
	s.LogStatus = &v
	return s
}

func (s *DescribeUserWafLogStatusResponseBody) SetRequestId(v string) *DescribeUserWafLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserWafLogStatusResponseBody) SetStatusUpdateTime(v int64) *DescribeUserWafLogStatusResponseBody {
	s.StatusUpdateTime = &v
	return s
}

func (s *DescribeUserWafLogStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
