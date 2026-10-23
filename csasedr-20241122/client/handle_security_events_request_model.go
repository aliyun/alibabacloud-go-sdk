// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSecurityEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarkBatch(v bool) *HandleSecurityEventsRequest
	GetMarkBatch() *bool
	SetMarkMissParam(v string) *HandleSecurityEventsRequest
	GetMarkMissParam() *string
	SetOperationCode(v string) *HandleSecurityEventsRequest
	GetOperationCode() *string
	SetOperationParams(v string) *HandleSecurityEventsRequest
	GetOperationParams() *string
	SetRemark(v string) *HandleSecurityEventsRequest
	GetRemark() *string
	SetSecurityEventIds(v []*int64) *HandleSecurityEventsRequest
	GetSecurityEventIds() []*int64
}

type HandleSecurityEventsRequest struct {
	// Specifies whether to add to the whitelist in batches. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	MarkBatch *bool `json:"MarkBatch,omitempty" xml:"MarkBatch,omitempty"`
	// The JSON array string used when adding an advanced whitelist or precise defense whitelist rule. This parameter applies only when OperationCode is set to `advance_mark_mis_info` or `defense_mark_mis_info`. Do not pass this parameter for other operations.
	//
	// Each rule contains the following fields:
	//
	// - `field`: The matching field, which is obtained from the FiledName in MarkFieldsSource returned by DescribeSecurityEventOperations.
	//
	// - `fieldValue`: The user-defined matching value, which does not need to be equal to the original field value of the current alert.
	//
	// - `operate`: The matching operator, which must belong to the SupportedMisType list returned for the field. Do not use a general fixed enumeration instead of the supported range for the current field.
	//
	// - `assetScope`: The effective scope object of the whitelist rule, which includes matchMode (scope mode) and the corresponding matchTargetIds (scope target list).
	//
	// The assetScope parameter specifies the effective scope of the rule. If you only handle the alert and do not add a whitelist rule, you can omit this parameter or pass an empty array string `[]`.
	//
	// example:
	//
	// [{"field":"processPath","fieldValue":"/tmp/evil","operate":"strEqual","assetScope":{"matchMode":"DeviceGroupNormal","matchTargetIds":["device-group-1"]}}]
	MarkMissParam *string `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
	// The operation code for handling the endpoint alert. Valid values:
	//
	// - advance_mark_mis_info: Adds to the advanced whitelist.
	//
	// - defense_mark_mis_info: Adds to the precise defense whitelist.
	//
	// - mark_mis_info: Marks as a false positive (standard whitelisting).
	//
	// - rm_mark_mis_info: Removes from the standard or advanced whitelist.
	//
	// - rm_defense_mark_mis_info: Removes from the precise defense whitelist.
	//
	// - ignore: Ignores the alert.
	//
	// - manual_handled: Marks as manually handled.
	//
	// - cancle_manual: Cancels the manually handled status.
	//
	// The preceding list is not a fixed set of operations that can be performed on every alert. The actual available operations are also restricted by feature switches, dynamic whitelists, alert operation configurations, and the current status. Before you call this operation, call DescribeSecurityEventOperations to query the currently available operations for the alert and check UserCanOperate. Use only the OperationCode values that are allowed in the query result.
	//
	// This parameter is required.
	//
	// example:
	//
	// ignore
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The operation parameters, represented as a JSON object string. If this parameter is not passed or an empty string is passed, the backend processes it as `{}`. The specific parameters must correspond to the OperationCode and are subject to the OperationParams returned by DescribeSecurityEventOperations.
	//
	// example:
	//
	// {}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// The remark for the handling operation.
	//
	// example:
	//
	// remark test.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IDs of the security alert events to handle.
	//
	// This parameter is required.
	SecurityEventIds []*int64 `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
}

func (s HandleSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequest) GetMarkBatch() *bool {
	return s.MarkBatch
}

func (s *HandleSecurityEventsRequest) GetMarkMissParam() *string {
	return s.MarkMissParam
}

func (s *HandleSecurityEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *HandleSecurityEventsRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *HandleSecurityEventsRequest) GetRemark() *string {
	return s.Remark
}

func (s *HandleSecurityEventsRequest) GetSecurityEventIds() []*int64 {
	return s.SecurityEventIds
}

func (s *HandleSecurityEventsRequest) SetMarkBatch(v bool) *HandleSecurityEventsRequest {
	s.MarkBatch = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetMarkMissParam(v string) *HandleSecurityEventsRequest {
	s.MarkMissParam = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationCode(v string) *HandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationParams(v string) *HandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetRemark(v string) *HandleSecurityEventsRequest {
	s.Remark = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetSecurityEventIds(v []*int64) *HandleSecurityEventsRequest {
	s.SecurityEventIds = v
	return s
}

func (s *HandleSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}
