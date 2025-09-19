// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaselineCheckWhiteRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddBaselineCheckWhiteRecordResponseBodyData) *AddBaselineCheckWhiteRecordResponseBody
	GetData() *AddBaselineCheckWhiteRecordResponseBodyData
	SetRequestId(v string) *AddBaselineCheckWhiteRecordResponseBody
	GetRequestId() *string
}

type AddBaselineCheckWhiteRecordResponseBody struct {
	// The data returned.
	Data *AddBaselineCheckWhiteRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBaselineCheckWhiteRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBaselineCheckWhiteRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddBaselineCheckWhiteRecordResponseBody) GetData() *AddBaselineCheckWhiteRecordResponseBodyData {
	return s.Data
}

func (s *AddBaselineCheckWhiteRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBaselineCheckWhiteRecordResponseBody) SetData(v *AddBaselineCheckWhiteRecordResponseBodyData) *AddBaselineCheckWhiteRecordResponseBody {
	s.Data = v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBody) SetRequestId(v string) *AddBaselineCheckWhiteRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddBaselineCheckWhiteRecordResponseBodyData struct {
	// The ID of the check item.
	//
	// >  You can call the [ListCheckItemWarningSummary](~~ListCheckItemWarningSummary~~) operation to query the IDs of check items.
	//
	// example:
	//
	// 76
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason why the check item is added to the whitelist.
	//
	// example:
	//
	// AutoTest
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 864153
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The data source. Valid values:
	//
	// 	- **default**: server
	//
	// 	- **agentless**: agentless detection
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The object that is added to the whitelist.
	//
	// example:
	//
	// HOST_BASELINE_WHITE_LIST_21
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the assets on which the whitelist rule takes effect. Valid values:
	//
	// 	- **all_instance**: all servers
	//
	// 	- **instance**: specific servers
	//
	// example:
	//
	// instance
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s AddBaselineCheckWhiteRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddBaselineCheckWhiteRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) GetCheckId() *int64 {
	return s.CheckId
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) GetLang() *string {
	return s.Lang
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) GetRecordId() *int64 {
	return s.RecordId
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) SetCheckId(v int64) *AddBaselineCheckWhiteRecordResponseBodyData {
	s.CheckId = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) SetLang(v string) *AddBaselineCheckWhiteRecordResponseBodyData {
	s.Lang = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) SetReason(v string) *AddBaselineCheckWhiteRecordResponseBodyData {
	s.Reason = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) SetRecordId(v int64) *AddBaselineCheckWhiteRecordResponseBodyData {
	s.RecordId = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) SetSource(v string) *AddBaselineCheckWhiteRecordResponseBodyData {
	s.Source = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) SetTarget(v string) *AddBaselineCheckWhiteRecordResponseBodyData {
	s.Target = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) SetTargetType(v string) *AddBaselineCheckWhiteRecordResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
