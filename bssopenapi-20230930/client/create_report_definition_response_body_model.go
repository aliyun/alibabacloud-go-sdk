// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginBillingCycle(v string) *CreateReportDefinitionResponseBody
	GetBeginBillingCycle() *string
	SetMetadata(v interface{}) *CreateReportDefinitionResponseBody
	GetMetadata() interface{}
	SetOssBucketName(v string) *CreateReportDefinitionResponseBody
	GetOssBucketName() *string
	SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionResponseBody
	GetOssBucketOwnerAccountId() *int64
	SetOssBucketPath(v string) *CreateReportDefinitionResponseBody
	GetOssBucketPath() *string
	SetReportSourceName(v string) *CreateReportDefinitionResponseBody
	GetReportSourceName() *string
	SetReportSourceType(v string) *CreateReportDefinitionResponseBody
	GetReportSourceType() *string
	SetReportTaskId(v int64) *CreateReportDefinitionResponseBody
	GetReportTaskId() *int64
	SetReportType(v string) *CreateReportDefinitionResponseBody
	GetReportType() *string
	SetRequestId(v string) *CreateReportDefinitionResponseBody
	GetRequestId() *string
	SetSubscribeCreateTime(v string) *CreateReportDefinitionResponseBody
	GetSubscribeCreateTime() *string
}

type CreateReportDefinitionResponseBody struct {
	// The start billing cycle for push. After successful subscription, the system automatically pushes data from the start billing cycle to the current time. This parameter is invalid for monthly bill PDF subscriptions and does not re-push historical data. Data within the last year can be pushed.
	//
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The response struct metadata.
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the OSS bucket for file storage.
	//
	// example:
	//
	// sh-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The UID of the OSS owner that stores the files. If this is a Bid/Reseller subscription and you need to push to a sub-account\\"s OSS, specify this parameter. The account must be a sub-account of the calling account, and the AliyunConsumeDump2OSSRole permission must be granted to this account. Regular users do not need to specify this parameter. The default value is the calling account.
	//
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// The OSS bucket storage path.
	//
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// The subscription source name.
	//
	// example:
	//
	// OSS
	ReportSourceName *string `json:"ReportSourceName,omitempty" xml:"ReportSourceName,omitempty"`
	// The subscription source. Valid values: OSS or MC.
	//
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// The bill subscription task ID.
	//
	// example:
	//
	// 123123
	ReportTaskId *int64 `json:"ReportTaskId,omitempty" xml:"ReportTaskId,omitempty"`
	// The subscription type. Valid values:
	//
	// - consumeDetailBillV2: consumption details (supported only for OSS/MC subscriptions).
	//
	// - splitDetailBillV2: split details (supported only for OSS/MC subscriptions).
	//
	// - costDetailBillV2: cost details (supported only for OSS/MC subscriptions).
	//
	// - monthBillOverview: monthly bill summary (supported only for OSS/MSC_EMAIL subscriptions).
	//
	// - focus: FOCUS bill (supported only for OSS/MC subscriptions).
	//
	// example:
	//
	// consumeDetailBillV2
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 340CAB45-0637-5875-9BE4-EFD5750F6BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The subscription creation time.
	//
	// example:
	//
	// 2025-05-21 10:36:31
	SubscribeCreateTime *string `json:"SubscribeCreateTime,omitempty" xml:"SubscribeCreateTime,omitempty"`
}

func (s CreateReportDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReportDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionResponseBody) GetBeginBillingCycle() *string {
	return s.BeginBillingCycle
}

func (s *CreateReportDefinitionResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateReportDefinitionResponseBody) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateReportDefinitionResponseBody) GetOssBucketOwnerAccountId() *int64 {
	return s.OssBucketOwnerAccountId
}

func (s *CreateReportDefinitionResponseBody) GetOssBucketPath() *string {
	return s.OssBucketPath
}

func (s *CreateReportDefinitionResponseBody) GetReportSourceName() *string {
	return s.ReportSourceName
}

func (s *CreateReportDefinitionResponseBody) GetReportSourceType() *string {
	return s.ReportSourceType
}

func (s *CreateReportDefinitionResponseBody) GetReportTaskId() *int64 {
	return s.ReportTaskId
}

func (s *CreateReportDefinitionResponseBody) GetReportType() *string {
	return s.ReportType
}

func (s *CreateReportDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReportDefinitionResponseBody) GetSubscribeCreateTime() *string {
	return s.SubscribeCreateTime
}

func (s *CreateReportDefinitionResponseBody) SetBeginBillingCycle(v string) *CreateReportDefinitionResponseBody {
	s.BeginBillingCycle = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetMetadata(v interface{}) *CreateReportDefinitionResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketName(v string) *CreateReportDefinitionResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionResponseBody {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketPath(v string) *CreateReportDefinitionResponseBody {
	s.OssBucketPath = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportSourceName(v string) *CreateReportDefinitionResponseBody {
	s.ReportSourceName = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportSourceType(v string) *CreateReportDefinitionResponseBody {
	s.ReportSourceType = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportTaskId(v int64) *CreateReportDefinitionResponseBody {
	s.ReportTaskId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportType(v string) *CreateReportDefinitionResponseBody {
	s.ReportType = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetRequestId(v string) *CreateReportDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetSubscribeCreateTime(v string) *CreateReportDefinitionResponseBody {
	s.SubscribeCreateTime = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
