// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataBatchIngestionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScanNew(v string) *UpdateDataBatchIngestionShrinkRequest
	GetAutoScanNew() *string
	SetDataBatchIngestionMode(v string) *UpdateDataBatchIngestionShrinkRequest
	GetDataBatchIngestionMode() *string
	SetDataIngestionIdsShrink(v string) *UpdateDataBatchIngestionShrinkRequest
	GetDataIngestionIdsShrink() *string
	SetDataSourceRecognizeEnabled(v bool) *UpdateDataBatchIngestionShrinkRequest
	GetDataSourceRecognizeEnabled() *bool
	SetLang(v string) *UpdateDataBatchIngestionShrinkRequest
	GetLang() *string
	SetLogUserIdsShrink(v string) *UpdateDataBatchIngestionShrinkRequest
	GetLogUserIdsShrink() *string
	SetRegionId(v string) *UpdateDataBatchIngestionShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataBatchIngestionShrinkRequest
	GetRoleFor() *int64
}

type UpdateDataBatchIngestionShrinkRequest struct {
	// Specifies whether to automatically discover new users.
	//
	// - enabled: Enables the feature.
	//
	// - disabled: Disables the feature.
	//
	// example:
	//
	// enabled
	AutoScanNew *string `json:"AutoScanNew,omitempty" xml:"AutoScanNew,omitempty"`
	// The mode for batch data ingestion. Valid values:
	//
	// - full
	//
	// - increment
	//
	// example:
	//
	// full
	DataBatchIngestionMode *string `json:"DataBatchIngestionMode,omitempty" xml:"DataBatchIngestionMode,omitempty"`
	// The list of ingestion policy IDs.
	DataIngestionIdsShrink *string `json:"DataIngestionIds,omitempty" xml:"DataIngestionIds,omitempty"`
	// Specifies whether to automatically discover new Logstores.
	//
	// example:
	//
	// true
	DataSourceRecognizeEnabled *bool `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of user IDs for batch data ingestion.
	LogUserIdsShrink *string `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty"`
	// The region of the Data Management hub for threat analysis. Select a region for the management hub based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose perspective the administrator wants to switch to.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataBatchIngestionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataBatchIngestionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetAutoScanNew() *string {
	return s.AutoScanNew
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetDataBatchIngestionMode() *string {
	return s.DataBatchIngestionMode
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetDataIngestionIdsShrink() *string {
	return s.DataIngestionIdsShrink
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetLogUserIdsShrink() *string {
	return s.LogUserIdsShrink
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataBatchIngestionShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetAutoScanNew(v string) *UpdateDataBatchIngestionShrinkRequest {
	s.AutoScanNew = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetDataBatchIngestionMode(v string) *UpdateDataBatchIngestionShrinkRequest {
	s.DataBatchIngestionMode = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetDataIngestionIdsShrink(v string) *UpdateDataBatchIngestionShrinkRequest {
	s.DataIngestionIdsShrink = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetDataSourceRecognizeEnabled(v bool) *UpdateDataBatchIngestionShrinkRequest {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetLang(v string) *UpdateDataBatchIngestionShrinkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetLogUserIdsShrink(v string) *UpdateDataBatchIngestionShrinkRequest {
	s.LogUserIdsShrink = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetRegionId(v string) *UpdateDataBatchIngestionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) SetRoleFor(v int64) *UpdateDataBatchIngestionShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataBatchIngestionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
