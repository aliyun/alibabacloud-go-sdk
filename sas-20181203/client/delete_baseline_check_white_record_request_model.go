// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBaselineCheckWhiteRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *DeleteBaselineCheckWhiteRecordRequest
	GetCheckIds() []*int64
	SetLang(v string) *DeleteBaselineCheckWhiteRecordRequest
	GetLang() *string
	SetRecordIds(v []*int64) *DeleteBaselineCheckWhiteRecordRequest
	GetRecordIds() []*int64
	SetSource(v string) *DeleteBaselineCheckWhiteRecordRequest
	GetSource() *string
}

type DeleteBaselineCheckWhiteRecordRequest struct {
	// The IDs of check items.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
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
	// The IDs of the whitelist records.
	RecordIds []*int64 `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
	// The data source. Valid values:
	//
	// 	- **default**: host baseline
	//
	// 	- **agentless**: agentless detection
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DeleteBaselineCheckWhiteRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBaselineCheckWhiteRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteBaselineCheckWhiteRecordRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *DeleteBaselineCheckWhiteRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteBaselineCheckWhiteRecordRequest) GetRecordIds() []*int64 {
	return s.RecordIds
}

func (s *DeleteBaselineCheckWhiteRecordRequest) GetSource() *string {
	return s.Source
}

func (s *DeleteBaselineCheckWhiteRecordRequest) SetCheckIds(v []*int64) *DeleteBaselineCheckWhiteRecordRequest {
	s.CheckIds = v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordRequest) SetLang(v string) *DeleteBaselineCheckWhiteRecordRequest {
	s.Lang = &v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordRequest) SetRecordIds(v []*int64) *DeleteBaselineCheckWhiteRecordRequest {
	s.RecordIds = v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordRequest) SetSource(v string) *DeleteBaselineCheckWhiteRecordRequest {
	s.Source = &v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordRequest) Validate() error {
	return dara.Validate(s)
}
