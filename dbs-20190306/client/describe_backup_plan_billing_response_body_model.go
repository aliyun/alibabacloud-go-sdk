// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanBillingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeBackupPlanBillingResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupPlanBillingResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeBackupPlanBillingResponseBody
	GetHttpStatusCode() *int32
	SetItem(v *DescribeBackupPlanBillingResponseBodyItem) *DescribeBackupPlanBillingResponseBody
	GetItem() *DescribeBackupPlanBillingResponseBodyItem
	SetRequestId(v string) *DescribeBackupPlanBillingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupPlanBillingResponseBody
	GetSuccess() *bool
}

type DescribeBackupPlanBillingResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This backupPlan can\\"t support this action
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The billing information of the backup schedule.
	Item *DescribeBackupPlanBillingResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EB4DFD5E-3618-498D-BE35-4DBEA0072122
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPlanBillingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanBillingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupPlanBillingResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupPlanBillingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBackupPlanBillingResponseBody) GetItem() *DescribeBackupPlanBillingResponseBodyItem {
	return s.Item
}

func (s *DescribeBackupPlanBillingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPlanBillingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupPlanBillingResponseBody) SetErrCode(v string) *DescribeBackupPlanBillingResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetErrMessage(v string) *DescribeBackupPlanBillingResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetHttpStatusCode(v int32) *DescribeBackupPlanBillingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetItem(v *DescribeBackupPlanBillingResponseBodyItem) *DescribeBackupPlanBillingResponseBody {
	s.Item = v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetRequestId(v string) *DescribeBackupPlanBillingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetSuccess(v bool) *DescribeBackupPlanBillingResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) Validate() error {
	if s.Item != nil {
		if err := s.Item.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPlanBillingResponseBodyItem struct {
	// The billing method. Valid values:
	//
	// 	- PREPAY
	//
	// 	- POSTPAY
	//
	// example:
	//
	// PREPAY
	BuyChargeType *string `json:"BuyChargeType,omitempty" xml:"BuyChargeType,omitempty"`
	// The timestamp that indicates when the instance was purchased.
	//
	// example:
	//
	// 1554560477000
	BuyCreateTimestamp *int64 `json:"BuyCreateTimestamp,omitempty" xml:"BuyCreateTimestamp,omitempty"`
	// The timestamp that indicates when the instance expires.
	//
	// > This parameter is available only if the value of the BuyChargeType parameter is PREPAY.
	//
	// example:
	//
	// 1554560477000
	BuyExpiredTimestamp *int64 `json:"BuyExpiredTimestamp,omitempty" xml:"BuyExpiredTimestamp,omitempty"`
	// The specifications of the instance.
	//
	// example:
	//
	// micro
	BuySpec *string `json:"BuySpec,omitempty" xml:"BuySpec,omitempty"`
	// The size of the built-in storage for storing incremental backup data.
	//
	// example:
	//
	// 12134
	ContStorageSize *int64 `json:"ContStorageSize,omitempty" xml:"ContStorageSize,omitempty"`
	// The size of the built-in storage for storing full backup data.
	//
	// example:
	//
	// 13123
	FullStorageSize *int64 `json:"FullStorageSize,omitempty" xml:"FullStorageSize,omitempty"`
	// Indicates whether the instance expired.
	//
	// > This parameter is available only if the value of the BuyChargeType parameter is PREPAY.
	//
	// example:
	//
	// true
	IsExpired *bool `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// Indicates whether the instance has no backup traffic limit.
	//
	// example:
	//
	// true
	IsFreeBytesUnlimited *bool `json:"IsFreeBytesUnlimited,omitempty" xml:"IsFreeBytesUnlimited,omitempty"`
	// The total paid backup traffic in the current month.
	//
	// example:
	//
	// 213213
	PaiedBytes *int64 `json:"PaiedBytes,omitempty" xml:"PaiedBytes,omitempty"`
	// The timestamp that indicates when the billing cycle of free backup traffic ends.
	//
	// example:
	//
	// 1554560477000
	QuotaEndTimestamp *int64 `json:"QuotaEndTimestamp,omitempty" xml:"QuotaEndTimestamp,omitempty"`
	// The timestamp that indicates when the billing cycle of free backup traffic starts.
	//
	// example:
	//
	// 1554560477000
	QuotaStartTimestamp *int64 `json:"QuotaStartTimestamp,omitempty" xml:"QuotaStartTimestamp,omitempty"`
	// The total free backup traffic in the current month.
	//
	// > This parameter is available only if the value of the BuyChargeType parameter is PREPAY and the value of the IsFreeBytesUnlimited parameter is false.
	//
	// example:
	//
	// 13123
	TotalFreeBytes *int64 `json:"TotalFreeBytes,omitempty" xml:"TotalFreeBytes,omitempty"`
	// The paid full backup traffic in the current month.
	//
	// example:
	//
	// 1312313
	UsedFullBytes *int64 `json:"UsedFullBytes,omitempty" xml:"UsedFullBytes,omitempty"`
	// The paid incremental backup traffic in the current month.
	//
	// example:
	//
	// 131231
	UsedIncrementBytes *int64 `json:"UsedIncrementBytes,omitempty" xml:"UsedIncrementBytes,omitempty"`
}

func (s DescribeBackupPlanBillingResponseBodyItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanBillingResponseBodyItem) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetBuyChargeType() *string {
	return s.BuyChargeType
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetBuyCreateTimestamp() *int64 {
	return s.BuyCreateTimestamp
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetBuyExpiredTimestamp() *int64 {
	return s.BuyExpiredTimestamp
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetBuySpec() *string {
	return s.BuySpec
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetContStorageSize() *int64 {
	return s.ContStorageSize
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetFullStorageSize() *int64 {
	return s.FullStorageSize
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetIsFreeBytesUnlimited() *bool {
	return s.IsFreeBytesUnlimited
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetPaiedBytes() *int64 {
	return s.PaiedBytes
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetQuotaEndTimestamp() *int64 {
	return s.QuotaEndTimestamp
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetQuotaStartTimestamp() *int64 {
	return s.QuotaStartTimestamp
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetTotalFreeBytes() *int64 {
	return s.TotalFreeBytes
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetUsedFullBytes() *int64 {
	return s.UsedFullBytes
}

func (s *DescribeBackupPlanBillingResponseBodyItem) GetUsedIncrementBytes() *int64 {
	return s.UsedIncrementBytes
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuyChargeType(v string) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuyChargeType = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuyCreateTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuyCreateTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuyExpiredTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuyExpiredTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuySpec(v string) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuySpec = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetContStorageSize(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.ContStorageSize = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetFullStorageSize(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.FullStorageSize = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetIsExpired(v bool) *DescribeBackupPlanBillingResponseBodyItem {
	s.IsExpired = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetIsFreeBytesUnlimited(v bool) *DescribeBackupPlanBillingResponseBodyItem {
	s.IsFreeBytesUnlimited = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetPaiedBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.PaiedBytes = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetQuotaEndTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.QuotaEndTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetQuotaStartTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.QuotaStartTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetTotalFreeBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.TotalFreeBytes = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetUsedFullBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.UsedFullBytes = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetUsedIncrementBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.UsedIncrementBytes = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) Validate() error {
	return dara.Validate(s)
}
