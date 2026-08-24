// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportDefinitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v interface{}) *ListReportDefinitionsResponseBody
	GetMetadata() interface{}
	SetReportDefinitions(v []*ListReportDefinitionsResponseBodyReportDefinitions) *ListReportDefinitionsResponseBody
	GetReportDefinitions() []*ListReportDefinitionsResponseBodyReportDefinitions
	SetRequestId(v string) *ListReportDefinitionsResponseBody
	GetRequestId() *string
}

type ListReportDefinitionsResponseBody struct {
	// The metadata of the response struct.
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The data list.
	ReportDefinitions []*ListReportDefinitionsResponseBodyReportDefinitions `json:"ReportDefinitions,omitempty" xml:"ReportDefinitions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListReportDefinitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReportDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListReportDefinitionsResponseBody) GetReportDefinitions() []*ListReportDefinitionsResponseBodyReportDefinitions {
	return s.ReportDefinitions
}

func (s *ListReportDefinitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListReportDefinitionsResponseBody) SetMetadata(v interface{}) *ListReportDefinitionsResponseBody {
	s.Metadata = v
	return s
}

func (s *ListReportDefinitionsResponseBody) SetReportDefinitions(v []*ListReportDefinitionsResponseBodyReportDefinitions) *ListReportDefinitionsResponseBody {
	s.ReportDefinitions = v
	return s
}

func (s *ListReportDefinitionsResponseBody) SetRequestId(v string) *ListReportDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListReportDefinitionsResponseBody) Validate() error {
	if s.ReportDefinitions != nil {
		for _, item := range s.ReportDefinitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReportDefinitionsResponseBodyReportDefinitions struct {
	// The start billing cycle for push. After the subscription is created, the system automatically pushes data from the start billing cycle to the current time. This parameter does not take effect for monthly bill PDF subscriptions, and historical data is not re-pushed. Data within the last year can be pushed.
	//
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The name of the OSS bucket that stores the files.
	//
	// example:
	//
	// oss-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The UID of the OSS bucket owner that stores the files. Specify this parameter when a Bid/Reseller subscription needs to push data to the OSS bucket of a sub-account. The specified account must be a sub-account of the calling account and must be granted the AliyunConsumeDump2OSSRole permission. Regular users do not need to specify this parameter. The calling account is used by default.
	//
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// The storage path of the OSS bucket.
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
	// The subscription source. Valid values: OSS and MC.
	//
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// The bill subscription task ID.
	//
	// example:
	//
	// 123321
	ReportTaskId *int64 `json:"ReportTaskId,omitempty" xml:"ReportTaskId,omitempty"`
	// The subscription type. Valid values:
	//
	// - BillingItemDetailForBillingPeriod: billable item consumption details.
	//
	// - InstanceDetailForBillingPeriod: instance consumption details.
	//
	// - BillingItemDetailMonthly: billable item consumption summary by billing cycle.
	//
	// example:
	//
	// BillingItemDetailForBillingPeriod
	ReportType     *string   `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	SelectedFields []*string `json:"SelectedFields,omitempty" xml:"SelectedFields,omitempty" type:"Repeated"`
	// The time when the subscription was created.
	//
	// example:
	//
	// 2025-05-21 10:36:31
	SubscribeCreateTime *string `json:"SubscribeCreateTime,omitempty" xml:"SubscribeCreateTime,omitempty"`
}

func (s ListReportDefinitionsResponseBodyReportDefinitions) String() string {
	return dara.Prettify(s)
}

func (s ListReportDefinitionsResponseBodyReportDefinitions) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetBeginBillingCycle() *string {
	return s.BeginBillingCycle
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetOssBucketOwnerAccountId() *int64 {
	return s.OssBucketOwnerAccountId
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetOssBucketPath() *string {
	return s.OssBucketPath
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetReportSourceName() *string {
	return s.ReportSourceName
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetReportSourceType() *string {
	return s.ReportSourceType
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetReportTaskId() *int64 {
	return s.ReportTaskId
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetReportType() *string {
	return s.ReportType
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetSelectedFields() []*string {
	return s.SelectedFields
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) GetSubscribeCreateTime() *string {
	return s.SubscribeCreateTime
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetBeginBillingCycle(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.BeginBillingCycle = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetOssBucketName(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.OssBucketName = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetOssBucketOwnerAccountId(v int64) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetOssBucketPath(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.OssBucketPath = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportSourceName(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportSourceName = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportSourceType(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportSourceType = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportTaskId(v int64) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportTaskId = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportType(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportType = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetSelectedFields(v []*string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.SelectedFields = v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetSubscribeCreateTime(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.SubscribeCreateTime = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) Validate() error {
	return dara.Validate(s)
}
