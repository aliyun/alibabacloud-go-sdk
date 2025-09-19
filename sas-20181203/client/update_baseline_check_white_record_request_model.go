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
	// >  You can call the [ListCheckItemWarningSummary](~~ListCheckItemWarningSummary~~) operation to query the IDs of check items.
	//
	// example:
	//
	// 16
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
	// The ID of the whitelist record.
	//
	// example:
	//
	// 14
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// A list of asset UUIDs from which container names need to be removed from the whitelist.
	RemoveContainerUuids []*string `json:"RemoveContainerUuids,omitempty" xml:"RemoveContainerUuids,omitempty" type:"Repeated"`
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
	// The type of the assets on which the whitelist rule takes effect. Valid values:
	//
	// 	- **all_instance**: all servers
	//
	// 	- **instance**: specific servers
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
