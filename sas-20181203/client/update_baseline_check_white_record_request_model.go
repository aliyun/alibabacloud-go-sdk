// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBaselineCheckWhiteRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *UpdateBaselineCheckWhiteRecordRequest
	GetCheckId() *int64
	SetLang(v string) *UpdateBaselineCheckWhiteRecordRequest
	GetLang() *string
	SetReason(v string) *UpdateBaselineCheckWhiteRecordRequest
	GetReason() *string
	SetRecordId(v int64) *UpdateBaselineCheckWhiteRecordRequest
	GetRecordId() *int64
	SetRemoveContainerUuids(v []*string) *UpdateBaselineCheckWhiteRecordRequest
	GetRemoveContainerUuids() []*string
	SetSource(v string) *UpdateBaselineCheckWhiteRecordRequest
	GetSource() *string
	SetTargetType(v string) *UpdateBaselineCheckWhiteRecordRequest
	GetTargetType() *string
}

type UpdateBaselineCheckWhiteRecordRequest struct {
	// The ID of the check item.
	//
	// > Call the [ListCheckItemWarningSummary](~~ListCheckItemWarningSummary~~) operation to obtain the check item ID.
	//
	// example:
	//
	// 16
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason for adding the whitelist entry.
	//
	// example:
	//
	// Manually processed.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the whitelist record.
	//
	// example:
	//
	// 14
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The list of asset UUIDs for which container names are to be removed from the whitelist.
	RemoveContainerUuids []*string `json:"RemoveContainerUuids,omitempty" xml:"RemoveContainerUuids,omitempty" type:"Repeated"`
	// The data source. Valid values:
	//
	// - **default**: host
	//
	// - **agentless**: agentless.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the target on which the whitelist takes effect. Valid values:
	//
	// - **all_instance**: all servers
	//
	// - **instance**: specific servers.
	//
	// example:
	//
	// all_instance
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateBaselineCheckWhiteRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineCheckWhiteRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateBaselineCheckWhiteRecordRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *UpdateBaselineCheckWhiteRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateBaselineCheckWhiteRecordRequest) GetReason() *string {
	return s.Reason
}

func (s *UpdateBaselineCheckWhiteRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateBaselineCheckWhiteRecordRequest) GetRemoveContainerUuids() []*string {
	return s.RemoveContainerUuids
}

func (s *UpdateBaselineCheckWhiteRecordRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateBaselineCheckWhiteRecordRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateBaselineCheckWhiteRecordRequest) SetCheckId(v int64) *UpdateBaselineCheckWhiteRecordRequest {
	s.CheckId = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordRequest) SetLang(v string) *UpdateBaselineCheckWhiteRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordRequest) SetReason(v string) *UpdateBaselineCheckWhiteRecordRequest {
	s.Reason = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordRequest) SetRecordId(v int64) *UpdateBaselineCheckWhiteRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordRequest) SetRemoveContainerUuids(v []*string) *UpdateBaselineCheckWhiteRecordRequest {
	s.RemoveContainerUuids = v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordRequest) SetSource(v string) *UpdateBaselineCheckWhiteRecordRequest {
	s.Source = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordRequest) SetTargetType(v string) *UpdateBaselineCheckWhiteRecordRequest {
	s.TargetType = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordRequest) Validate() error {
	return dara.Validate(s)
}
