// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBaselineCheckWhiteRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateBaselineCheckWhiteRecordResponseBodyData) *UpdateBaselineCheckWhiteRecordResponseBody
	GetData() *UpdateBaselineCheckWhiteRecordResponseBodyData
	SetRequestId(v string) *UpdateBaselineCheckWhiteRecordResponseBody
	GetRequestId() *string
}

type UpdateBaselineCheckWhiteRecordResponseBody struct {
	// The data returned.
	Data *UpdateBaselineCheckWhiteRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBaselineCheckWhiteRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineCheckWhiteRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBaselineCheckWhiteRecordResponseBody) GetData() *UpdateBaselineCheckWhiteRecordResponseBodyData {
	return s.Data
}

func (s *UpdateBaselineCheckWhiteRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBaselineCheckWhiteRecordResponseBody) SetData(v *UpdateBaselineCheckWhiteRecordResponseBodyData) *UpdateBaselineCheckWhiteRecordResponseBody {
	s.Data = v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBody) SetRequestId(v string) *UpdateBaselineCheckWhiteRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateBaselineCheckWhiteRecordResponseBodyData struct {
	// The ID of the check item.
	//
	// example:
	//
	// 92
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason why the check item is added to the whitelist.
	//
	// example:
	//
	// Test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the whitelist record.
	//
	// example:
	//
	// 1582
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
	// HOST_BASELINE_WHITE_LIST_23
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

func (s UpdateBaselineCheckWhiteRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineCheckWhiteRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) GetCheckId() *int64 {
	return s.CheckId
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) GetLang() *string {
	return s.Lang
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) SetCheckId(v int64) *UpdateBaselineCheckWhiteRecordResponseBodyData {
	s.CheckId = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) SetLang(v string) *UpdateBaselineCheckWhiteRecordResponseBodyData {
	s.Lang = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) SetReason(v string) *UpdateBaselineCheckWhiteRecordResponseBodyData {
	s.Reason = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) SetRecordId(v int64) *UpdateBaselineCheckWhiteRecordResponseBodyData {
	s.RecordId = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) SetSource(v string) *UpdateBaselineCheckWhiteRecordResponseBodyData {
	s.Source = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) SetTarget(v string) *UpdateBaselineCheckWhiteRecordResponseBodyData {
	s.Target = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) SetTargetType(v string) *UpdateBaselineCheckWhiteRecordResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
