// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppWarningByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAiAppWarningByPageResponseBody
	GetCurrentPage() *int32
	SetExt(v *ListAiAppWarningByPageResponseBodyExt) *ListAiAppWarningByPageResponseBody
	GetExt() *ListAiAppWarningByPageResponseBodyExt
	SetItems(v []*ListAiAppWarningByPageResponseBodyItems) *ListAiAppWarningByPageResponseBody
	GetItems() []*ListAiAppWarningByPageResponseBodyItems
	SetPageSize(v int32) *ListAiAppWarningByPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAiAppWarningByPageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAiAppWarningByPageResponseBody
	GetTotalCount() *int64
}

type ListAiAppWarningByPageResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The extension field.
	Ext *ListAiAppWarningByPageResponseBodyExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	// The data on the current page.
	Items []*ListAiAppWarningByPageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAiAppWarningByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppWarningByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiAppWarningByPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAiAppWarningByPageResponseBody) GetExt() *ListAiAppWarningByPageResponseBodyExt {
	return s.Ext
}

func (s *ListAiAppWarningByPageResponseBody) GetItems() []*ListAiAppWarningByPageResponseBodyItems {
	return s.Items
}

func (s *ListAiAppWarningByPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiAppWarningByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAiAppWarningByPageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAiAppWarningByPageResponseBody) SetCurrentPage(v int32) *ListAiAppWarningByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBody) SetExt(v *ListAiAppWarningByPageResponseBodyExt) *ListAiAppWarningByPageResponseBody {
	s.Ext = v
	return s
}

func (s *ListAiAppWarningByPageResponseBody) SetItems(v []*ListAiAppWarningByPageResponseBodyItems) *ListAiAppWarningByPageResponseBody {
	s.Items = v
	return s
}

func (s *ListAiAppWarningByPageResponseBody) SetPageSize(v int32) *ListAiAppWarningByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBody) SetRequestId(v string) *ListAiAppWarningByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBody) SetTotalCount(v int64) *ListAiAppWarningByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBody) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
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

type ListAiAppWarningByPageResponseBodyExt struct {
	// The option.
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
}

func (s ListAiAppWarningByPageResponseBodyExt) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppWarningByPageResponseBodyExt) GoString() string {
	return s.String()
}

func (s *ListAiAppWarningByPageResponseBodyExt) GetOption() map[string]interface{} {
	return s.Option
}

func (s *ListAiAppWarningByPageResponseBodyExt) SetOption(v map[string]interface{}) *ListAiAppWarningByPageResponseBodyExt {
	s.Option = v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyExt) Validate() error {
	return dara.Validate(s)
}

type ListAiAppWarningByPageResponseBodyItems struct {
	// appId。
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// name-xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The list of labels.
	Labels []*ListAiAppWarningByPageResponseBodyItemsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The service code.
	//
	// example:
	//
	// baselineCheck_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The trace ID used to correlate and trace alert events.
	//
	// example:
	//
	// 0bc3b4b0********516098843e19bc
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The number of alerts.
	//
	// example:
	//
	// 10
	WarningCount *int64 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
	// The time when the alert was triggered.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	WarningTime *string `json:"WarningTime,omitempty" xml:"WarningTime,omitempty"`
}

func (s ListAiAppWarningByPageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppWarningByPageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAiAppWarningByPageResponseBodyItems) GetAppId() *string {
	return s.AppId
}

func (s *ListAiAppWarningByPageResponseBodyItems) GetAppName() *string {
	return s.AppName
}

func (s *ListAiAppWarningByPageResponseBodyItems) GetLabels() []*ListAiAppWarningByPageResponseBodyItemsLabels {
	return s.Labels
}

func (s *ListAiAppWarningByPageResponseBodyItems) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListAiAppWarningByPageResponseBodyItems) GetTraceId() *string {
	return s.TraceId
}

func (s *ListAiAppWarningByPageResponseBodyItems) GetWarningCount() *int64 {
	return s.WarningCount
}

func (s *ListAiAppWarningByPageResponseBodyItems) GetWarningTime() *string {
	return s.WarningTime
}

func (s *ListAiAppWarningByPageResponseBodyItems) SetAppId(v string) *ListAiAppWarningByPageResponseBodyItems {
	s.AppId = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItems) SetAppName(v string) *ListAiAppWarningByPageResponseBodyItems {
	s.AppName = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItems) SetLabels(v []*ListAiAppWarningByPageResponseBodyItemsLabels) *ListAiAppWarningByPageResponseBodyItems {
	s.Labels = v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItems) SetServiceCode(v string) *ListAiAppWarningByPageResponseBodyItems {
	s.ServiceCode = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItems) SetTraceId(v string) *ListAiAppWarningByPageResponseBodyItems {
	s.TraceId = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItems) SetWarningCount(v int64) *ListAiAppWarningByPageResponseBodyItems {
	s.WarningCount = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItems) SetWarningTime(v string) *ListAiAppWarningByPageResponseBodyItems {
	s.WarningTime = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItems) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAiAppWarningByPageResponseBodyItemsLabels struct {
	// The count.
	//
	// example:
	//
	// 20
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The label name.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The label description.
	//
	// example:
	//
	// desc-xxx
	LabelDesc *string `json:"LabelDesc,omitempty" xml:"LabelDesc,omitempty"`
	// The type.
	//
	// example:
	//
	// promptAttack
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAiAppWarningByPageResponseBodyItemsLabels) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppWarningByPageResponseBodyItemsLabels) GoString() string {
	return s.String()
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) GetCount() *int64 {
	return s.Count
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) GetLabel() *string {
	return s.Label
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) GetLabelDesc() *string {
	return s.LabelDesc
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) GetType() *string {
	return s.Type
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) SetCount(v int64) *ListAiAppWarningByPageResponseBodyItemsLabels {
	s.Count = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) SetLabel(v string) *ListAiAppWarningByPageResponseBodyItemsLabels {
	s.Label = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) SetLabelDesc(v string) *ListAiAppWarningByPageResponseBodyItemsLabels {
	s.LabelDesc = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) SetType(v string) *ListAiAppWarningByPageResponseBodyItemsLabels {
	s.Type = &v
	return s
}

func (s *ListAiAppWarningByPageResponseBodyItemsLabels) Validate() error {
	return dara.Validate(s)
}
