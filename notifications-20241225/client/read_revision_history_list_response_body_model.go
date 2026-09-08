// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadRevisionHistoryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadRevisionHistoryListResponseBody
	GetCode() *string
	SetData(v *ReadRevisionHistoryListResponseBodyData) *ReadRevisionHistoryListResponseBody
	GetData() *ReadRevisionHistoryListResponseBodyData
	SetHttpCode(v int32) *ReadRevisionHistoryListResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ReadRevisionHistoryListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadRevisionHistoryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadRevisionHistoryListResponseBody
	GetSuccess() *bool
}

type ReadRevisionHistoryListResponseBody struct {
	// The error code returned by the system. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	Data *ReadRevisionHistoryListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// 2xx
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned when the call failed.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadRevisionHistoryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadRevisionHistoryListResponseBody) GetData() *ReadRevisionHistoryListResponseBodyData {
	return s.Data
}

func (s *ReadRevisionHistoryListResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ReadRevisionHistoryListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadRevisionHistoryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadRevisionHistoryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadRevisionHistoryListResponseBody) SetCode(v string) *ReadRevisionHistoryListResponseBody {
	s.Code = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBody) SetData(v *ReadRevisionHistoryListResponseBodyData) *ReadRevisionHistoryListResponseBody {
	s.Data = v
	return s
}

func (s *ReadRevisionHistoryListResponseBody) SetHttpCode(v int32) *ReadRevisionHistoryListResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBody) SetMessage(v string) *ReadRevisionHistoryListResponseBody {
	s.Message = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBody) SetRequestId(v string) *ReadRevisionHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBody) SetSuccess(v bool) *ReadRevisionHistoryListResponseBody {
	s.Success = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadRevisionHistoryListResponseBodyData struct {
	// The maximum number of entries.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of data.
	//
	// example:
	//
	// AAAAAT0x7j2M1Og+SpZ8n4WEjfo=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// A single row of returned data.
	Rows []*ReadRevisionHistoryListResponseBodyDataRows `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// The total number of messages in the category.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ReadRevisionHistoryListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadRevisionHistoryListResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadRevisionHistoryListResponseBodyData) GetRows() []*ReadRevisionHistoryListResponseBodyDataRows {
	return s.Rows
}

func (s *ReadRevisionHistoryListResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ReadRevisionHistoryListResponseBodyData) SetMaxResults(v int32) *ReadRevisionHistoryListResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyData) SetNextToken(v string) *ReadRevisionHistoryListResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyData) SetRows(v []*ReadRevisionHistoryListResponseBodyDataRows) *ReadRevisionHistoryListResponseBodyData {
	s.Rows = v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyData) SetTotalCount(v int32) *ReadRevisionHistoryListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyData) Validate() error {
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadRevisionHistoryListResponseBodyDataRows struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1305851476425884
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The event type code.
	//
	// example:
	//
	// prod_edu_content
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The channel group.
	//
	// example:
	//
	// base
	ChannelGroupCode *string `json:"ChannelGroupCode,omitempty" xml:"ChannelGroupCode,omitempty"`
	// The modified value.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "channelType": "sms",
	//
	//         "checkedState": "YES"
	//
	//     },
	//
	//     {
	//
	//         "channelType": "pmsg",
	//
	//         "checkedState": "NO"
	//
	//     },
	//
	//     {
	//
	//         "channelType": "email",
	//
	//         "checkedState": "NO"
	//
	//     }
	//
	// ]
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// The revision item code. Valid values:
	//
	// example:
	//
	// receive_channel
	OperationItemCode *string `json:"OperationItemCode,omitempty" xml:"OperationItemCode,omitempty"`
	// The revision item name.
	//
	// example:
	//
	// Reception channel
	OperationItemName *string `json:"OperationItemName,omitempty" xml:"OperationItemName,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1743583672981
	OperationTimestamp *int64 `json:"OperationTimestamp,omitempty" xml:"OperationTimestamp,omitempty"`
	// The IP address of the operator.
	//
	// example:
	//
	// /
	OperatorIp *string `json:"OperatorIp,omitempty" xml:"OperatorIp,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// 1662077279821892
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The UID of the operator.
	//
	// example:
	//
	// 1062132414049864
	OperatorUid *int64 `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The original value.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "channelType": "sms",
	//
	//         "checkedState": "NO"
	//
	//     },
	//
	//     {
	//
	//         "channelType": "pmsg",
	//
	//         "checkedState": "NO"
	//
	//     },
	//
	//     {
	//
	//         "channelType": "email",
	//
	//         "checkedState": "NO"
	//
	//     }
	//
	// ]
	OriginalValue *string `json:"OriginalValue,omitempty" xml:"OriginalValue,omitempty"`
	// The pagination information.
	PageSpec *ReadRevisionHistoryListResponseBodyDataRowsPageSpec `json:"PageSpec,omitempty" xml:"PageSpec,omitempty" type:"Struct"`
	// The remarks.
	//
	// example:
	//
	// /
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
}

func (s ReadRevisionHistoryListResponseBodyDataRows) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListResponseBodyDataRows) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetChannelGroupCode() *string {
	return s.ChannelGroupCode
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetNewValue() *string {
	return s.NewValue
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetOperationItemCode() *string {
	return s.OperationItemCode
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetOperationItemName() *string {
	return s.OperationItemName
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetOperationTimestamp() *int64 {
	return s.OperationTimestamp
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetOperatorIp() *string {
	return s.OperatorIp
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetOperatorUid() *int64 {
	return s.OperatorUid
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetOriginalValue() *string {
	return s.OriginalValue
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetPageSpec() *ReadRevisionHistoryListResponseBodyDataRowsPageSpec {
	return s.PageSpec
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) GetRemarks() *string {
	return s.Remarks
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetAliUid(v int64) *ReadRevisionHistoryListResponseBodyDataRows {
	s.AliUid = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetCategoryCode(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.CategoryCode = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetChannelGroupCode(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.ChannelGroupCode = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetNewValue(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.NewValue = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetOperationItemCode(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.OperationItemCode = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetOperationItemName(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.OperationItemName = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetOperationTimestamp(v int64) *ReadRevisionHistoryListResponseBodyDataRows {
	s.OperationTimestamp = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetOperatorIp(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.OperatorIp = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetOperatorName(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.OperatorName = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetOperatorUid(v int64) *ReadRevisionHistoryListResponseBodyDataRows {
	s.OperatorUid = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetOriginalValue(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.OriginalValue = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetPageSpec(v *ReadRevisionHistoryListResponseBodyDataRowsPageSpec) *ReadRevisionHistoryListResponseBodyDataRows {
	s.PageSpec = v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) SetRemarks(v string) *ReadRevisionHistoryListResponseBodyDataRows {
	s.Remarks = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRows) Validate() error {
	if s.PageSpec != nil {
		if err := s.PageSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadRevisionHistoryListResponseBodyDataRowsPageSpec struct {
	// The maximum number of entries.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of data.
	//
	// example:
	//
	// AAAAAT0x7j2M1Og+SpZ8n4WEjfo=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ReadRevisionHistoryListResponseBodyDataRowsPageSpec) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListResponseBodyDataRowsPageSpec) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListResponseBodyDataRowsPageSpec) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadRevisionHistoryListResponseBodyDataRowsPageSpec) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadRevisionHistoryListResponseBodyDataRowsPageSpec) SetMaxResults(v int32) *ReadRevisionHistoryListResponseBodyDataRowsPageSpec {
	s.MaxResults = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRowsPageSpec) SetNextToken(v string) *ReadRevisionHistoryListResponseBodyDataRowsPageSpec {
	s.NextToken = &v
	return s
}

func (s *ReadRevisionHistoryListResponseBodyDataRowsPageSpec) Validate() error {
	return dara.Validate(s)
}
