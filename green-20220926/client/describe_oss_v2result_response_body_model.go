// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssV2ResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeOssV2ResultResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeOssV2ResultResponseBodyItems) *DescribeOssV2ResultResponseBody
	GetItems() []*DescribeOssV2ResultResponseBodyItems
	SetPageSize(v int32) *DescribeOssV2ResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeOssV2ResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeOssV2ResultResponseBody
	GetTotalCount() *int64
}

type DescribeOssV2ResultResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeOssV2ResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOssV2ResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssV2ResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssV2ResultResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOssV2ResultResponseBody) GetItems() []*DescribeOssV2ResultResponseBodyItems {
	return s.Items
}

func (s *DescribeOssV2ResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOssV2ResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssV2ResultResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeOssV2ResultResponseBody) SetCurrentPage(v int32) *DescribeOssV2ResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssV2ResultResponseBody) SetItems(v []*DescribeOssV2ResultResponseBodyItems) *DescribeOssV2ResultResponseBody {
	s.Items = v
	return s
}

func (s *DescribeOssV2ResultResponseBody) SetPageSize(v int32) *DescribeOssV2ResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOssV2ResultResponseBody) SetRequestId(v string) *DescribeOssV2ResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssV2ResultResponseBody) SetTotalCount(v int64) *DescribeOssV2ResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOssV2ResultResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOssV2ResultResponseBodyItems struct {
	// example:
	//
	// buckect_test
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// image
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// true
	Freeze *bool `json:"Freeze,omitempty" xml:"Freeze,omitempty"`
	// example:
	//
	// FREEZED
	FreezeStatus *string `json:"FreezeStatus,omitempty" xml:"FreezeStatus,omitempty"`
	// example:
	//
	// ACL
	FreezeType    *string                                              `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	LabelDetails  []*DescribeOssV2ResultResponseBodyItemsLabelDetails  `json:"LabelDetails,omitempty" xml:"LabelDetails,omitempty" type:"Repeated"`
	LabelDetails2 []*DescribeOssV2ResultResponseBodyItemsLabelDetails2 `json:"LabelDetails2,omitempty" xml:"LabelDetails2,omitempty" type:"Repeated"`
	Labels        []*string                                            `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Labels2       []*string                                            `json:"Labels2,omitempty" xml:"Labels2,omitempty" type:"Repeated"`
	// example:
	//
	// FREEZE
	ManualFreezeAction *string `json:"ManualFreezeAction,omitempty" xml:"ManualFreezeAction,omitempty"`
	// example:
	//
	// 2025-08-09 12:00:00
	ManualOperateTime *string `json:"ManualOperateTime,omitempty" xml:"ManualOperateTime,omitempty"`
	// example:
	//
	// xx
	ManualOperator *string `json:"ManualOperator,omitempty" xml:"ManualOperator,omitempty"`
	// example:
	//
	// 54416c9b159df4a60ae03c04ccb94cb5
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// 1713014531569_958.png
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// low
	RiskLevel0 *string `json:"RiskLevel0,omitempty" xml:"RiskLevel0,omitempty"`
	// example:
	//
	// none
	RiskLevel2 *string `json:"RiskLevel2,omitempty" xml:"RiskLevel2,omitempty"`
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// Service code。
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// COPY
	SysDisposalStatus *string `json:"SysDisposalStatus,omitempty" xml:"SysDisposalStatus,omitempty"`
	// example:
	//
	// P_BT3FHS
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeOssV2ResultResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssV2ResultResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeOssV2ResultResponseBodyItems) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeOssV2ResultResponseBodyItems) GetCode() *string {
	return s.Code
}

func (s *DescribeOssV2ResultResponseBodyItems) GetContentType() *string {
	return s.ContentType
}

func (s *DescribeOssV2ResultResponseBodyItems) GetFreeze() *bool {
	return s.Freeze
}

func (s *DescribeOssV2ResultResponseBodyItems) GetFreezeStatus() *string {
	return s.FreezeStatus
}

func (s *DescribeOssV2ResultResponseBodyItems) GetFreezeType() *string {
	return s.FreezeType
}

func (s *DescribeOssV2ResultResponseBodyItems) GetLabelDetails() []*DescribeOssV2ResultResponseBodyItemsLabelDetails {
	return s.LabelDetails
}

func (s *DescribeOssV2ResultResponseBodyItems) GetLabelDetails2() []*DescribeOssV2ResultResponseBodyItemsLabelDetails2 {
	return s.LabelDetails2
}

func (s *DescribeOssV2ResultResponseBodyItems) GetLabels() []*string {
	return s.Labels
}

func (s *DescribeOssV2ResultResponseBodyItems) GetLabels2() []*string {
	return s.Labels2
}

func (s *DescribeOssV2ResultResponseBodyItems) GetManualFreezeAction() *string {
	return s.ManualFreezeAction
}

func (s *DescribeOssV2ResultResponseBodyItems) GetManualOperateTime() *string {
	return s.ManualOperateTime
}

func (s *DescribeOssV2ResultResponseBodyItems) GetManualOperator() *string {
	return s.ManualOperator
}

func (s *DescribeOssV2ResultResponseBodyItems) GetMd5() *string {
	return s.Md5
}

func (s *DescribeOssV2ResultResponseBodyItems) GetObject() *string {
	return s.Object
}

func (s *DescribeOssV2ResultResponseBodyItems) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssV2ResultResponseBodyItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeOssV2ResultResponseBodyItems) GetRiskLevel0() *string {
	return s.RiskLevel0
}

func (s *DescribeOssV2ResultResponseBodyItems) GetRiskLevel2() *string {
	return s.RiskLevel2
}

func (s *DescribeOssV2ResultResponseBodyItems) GetScanResult() *string {
	return s.ScanResult
}

func (s *DescribeOssV2ResultResponseBodyItems) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeOssV2ResultResponseBodyItems) GetSysDisposalStatus() *string {
	return s.SysDisposalStatus
}

func (s *DescribeOssV2ResultResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeOssV2ResultResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *DescribeOssV2ResultResponseBodyItems) SetBucket(v string) *DescribeOssV2ResultResponseBodyItems {
	s.Bucket = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetCode(v string) *DescribeOssV2ResultResponseBodyItems {
	s.Code = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetContentType(v string) *DescribeOssV2ResultResponseBodyItems {
	s.ContentType = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetFreeze(v bool) *DescribeOssV2ResultResponseBodyItems {
	s.Freeze = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetFreezeStatus(v string) *DescribeOssV2ResultResponseBodyItems {
	s.FreezeStatus = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetFreezeType(v string) *DescribeOssV2ResultResponseBodyItems {
	s.FreezeType = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetLabelDetails(v []*DescribeOssV2ResultResponseBodyItemsLabelDetails) *DescribeOssV2ResultResponseBodyItems {
	s.LabelDetails = v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetLabelDetails2(v []*DescribeOssV2ResultResponseBodyItemsLabelDetails2) *DescribeOssV2ResultResponseBodyItems {
	s.LabelDetails2 = v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetLabels(v []*string) *DescribeOssV2ResultResponseBodyItems {
	s.Labels = v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetLabels2(v []*string) *DescribeOssV2ResultResponseBodyItems {
	s.Labels2 = v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetManualFreezeAction(v string) *DescribeOssV2ResultResponseBodyItems {
	s.ManualFreezeAction = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetManualOperateTime(v string) *DescribeOssV2ResultResponseBodyItems {
	s.ManualOperateTime = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetManualOperator(v string) *DescribeOssV2ResultResponseBodyItems {
	s.ManualOperator = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetMd5(v string) *DescribeOssV2ResultResponseBodyItems {
	s.Md5 = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetObject(v string) *DescribeOssV2ResultResponseBodyItems {
	s.Object = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetRequestId(v string) *DescribeOssV2ResultResponseBodyItems {
	s.RequestId = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetRiskLevel(v string) *DescribeOssV2ResultResponseBodyItems {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetRiskLevel0(v string) *DescribeOssV2ResultResponseBodyItems {
	s.RiskLevel0 = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetRiskLevel2(v string) *DescribeOssV2ResultResponseBodyItems {
	s.RiskLevel2 = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetScanResult(v string) *DescribeOssV2ResultResponseBodyItems {
	s.ScanResult = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetServiceCode(v string) *DescribeOssV2ResultResponseBodyItems {
	s.ServiceCode = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetSysDisposalStatus(v string) *DescribeOssV2ResultResponseBodyItems {
	s.SysDisposalStatus = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetTaskId(v string) *DescribeOssV2ResultResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) SetUrl(v string) *DescribeOssV2ResultResponseBodyItems {
	s.Url = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItems) Validate() error {
	if s.LabelDetails != nil {
		for _, item := range s.LabelDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LabelDetails2 != nil {
		for _, item := range s.LabelDetails2 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOssV2ResultResponseBodyItemsLabelDetails struct {
	// example:
	//
	// 50.00
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// politics
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeOssV2ResultResponseBodyItemsLabelDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssV2ResultResponseBodyItemsLabelDetails) GoString() string {
	return s.String()
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails) GetConfidence() *float32 {
	return s.Confidence
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails) GetDescription() *string {
	return s.Description
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails) GetLabel() *string {
	return s.Label
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails) SetConfidence(v float32) *DescribeOssV2ResultResponseBodyItemsLabelDetails {
	s.Confidence = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails) SetDescription(v string) *DescribeOssV2ResultResponseBodyItemsLabelDetails {
	s.Description = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails) SetLabel(v string) *DescribeOssV2ResultResponseBodyItemsLabelDetails {
	s.Label = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeOssV2ResultResponseBodyItemsLabelDetails2 struct {
	// example:
	//
	// 90.00
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// politics
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeOssV2ResultResponseBodyItemsLabelDetails2) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssV2ResultResponseBodyItemsLabelDetails2) GoString() string {
	return s.String()
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails2) GetConfidence() *float32 {
	return s.Confidence
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails2) GetDescription() *string {
	return s.Description
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails2) GetLabel() *string {
	return s.Label
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails2) SetConfidence(v float32) *DescribeOssV2ResultResponseBodyItemsLabelDetails2 {
	s.Confidence = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails2) SetDescription(v string) *DescribeOssV2ResultResponseBodyItemsLabelDetails2 {
	s.Description = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails2) SetLabel(v string) *DescribeOssV2ResultResponseBodyItemsLabelDetails2 {
	s.Label = &v
	return s
}

func (s *DescribeOssV2ResultResponseBodyItemsLabelDetails2) Validate() error {
	return dara.Validate(s)
}
