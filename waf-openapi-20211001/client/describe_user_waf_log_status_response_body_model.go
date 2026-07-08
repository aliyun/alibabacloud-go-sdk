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
	// The ID of the region where logs are stored. Valid values:
	//
	// - **cn-hangzhou**: indicates China East 1 (Hangzhou).
	//
	// - **cn-beijing**: indicates China North 2 (Beijing).
	//
	// - **cn-hongkong**: indicates China (Hong Kong).
	//
	// - **ap-southeast-1**: indicates Singapore.
	//
	// - **ap-southeast-3**: indicates Malaysia (Kuala Lumpur).
	//
	// - **ap-southeast-5**: indicates Indonesia (Jakarta).
	//
	// - **ap-southeast-6**: indicates Philippines (Manila).
	//
	// - **ap-southeast-7**: indicates Thailand (Bangkok).
	//
	// - **me-east-1**: indicates UAE (Dubai).
	//
	// - **eu-central-1**: indicates Germany (Frankfurt).
	//
	// - **us-east-1**: indicates US (Virginia).
	//
	// - **us-west-1**: indicates US (Silicon Valley).
	//
	// - **ap-northeast-1**: indicates Japan (Tokyo).
	//
	// - **ap-northeast-2**: indicates South Korea (Seoul).
	//
	// - **eu-west-1**: indicates UK (London).
	//
	// - **cn-hangzhou-finance**: indicates China East 1 Hangzhou Finance Cloud.
	//
	// - **cn-shanghai-finance-1**: indicates China East 2 Shanghai Finance Cloud.
	//
	// - **cn-shenzhen-finance**: indicates China South 1 Shenzhen Finance Cloud.
	//
	// > The Finance Cloud regions are available only to Finance Cloud users, and Finance Cloud users can obtain only these regions.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The status of WAF logs.
	//
	// - **initializing**: The logs are being initialized.
	//
	// - **initialize_failed**: The initialization failed.
	//
	// - **normal**: The logs are running properly.
	//
	// - **releasing**: The logs are being released.
	//
	// - **release_failed**: The release failed.
	//
	// example:
	//
	// normal
	LogStatus *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
