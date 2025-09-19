// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaselineCheckWhiteRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *AddBaselineCheckWhiteRecordRequest
	GetCheckId() *int64
	SetLang(v string) *AddBaselineCheckWhiteRecordRequest
	GetLang() *string
	SetReason(v string) *AddBaselineCheckWhiteRecordRequest
	GetReason() *string
	SetSource(v string) *AddBaselineCheckWhiteRecordRequest
	GetSource() *string
	SetTargetType(v string) *AddBaselineCheckWhiteRecordRequest
	GetTargetType() *string
}

type AddBaselineCheckWhiteRecordRequest struct {
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
	// AutoRun
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The data source. If you leave this parameter empty, the default value is used. Valid values:
	//
	// 	- **default**: server
	//
	// 	- **agentless**: agentless detection
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
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

func (s AddBaselineCheckWhiteRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBaselineCheckWhiteRecordRequest) GoString() string {
	return s.String()
}

func (s *AddBaselineCheckWhiteRecordRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *AddBaselineCheckWhiteRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *AddBaselineCheckWhiteRecordRequest) GetReason() *string {
	return s.Reason
}

func (s *AddBaselineCheckWhiteRecordRequest) GetSource() *string {
	return s.Source
}

func (s *AddBaselineCheckWhiteRecordRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *AddBaselineCheckWhiteRecordRequest) SetCheckId(v int64) *AddBaselineCheckWhiteRecordRequest {
	s.CheckId = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordRequest) SetLang(v string) *AddBaselineCheckWhiteRecordRequest {
	s.Lang = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordRequest) SetReason(v string) *AddBaselineCheckWhiteRecordRequest {
	s.Reason = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordRequest) SetSource(v string) *AddBaselineCheckWhiteRecordRequest {
	s.Source = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordRequest) SetTargetType(v string) *AddBaselineCheckWhiteRecordRequest {
	s.TargetType = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordRequest) Validate() error {
	return dara.Validate(s)
}
