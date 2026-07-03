// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataBatchIngestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScanNew(v string) *UpdateDataBatchIngestionRequest
	GetAutoScanNew() *string
	SetDataBatchIngestionMode(v string) *UpdateDataBatchIngestionRequest
	GetDataBatchIngestionMode() *string
	SetDataIngestionIds(v []*string) *UpdateDataBatchIngestionRequest
	GetDataIngestionIds() []*string
	SetDataSourceRecognizeEnabled(v bool) *UpdateDataBatchIngestionRequest
	GetDataSourceRecognizeEnabled() *bool
	SetLang(v string) *UpdateDataBatchIngestionRequest
	GetLang() *string
	SetLogUserIds(v []*int64) *UpdateDataBatchIngestionRequest
	GetLogUserIds() []*int64
	SetRegionId(v string) *UpdateDataBatchIngestionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataBatchIngestionRequest
	GetRoleFor() *int64
}

type UpdateDataBatchIngestionRequest struct {
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
	DataIngestionIds []*string `json:"DataIngestionIds,omitempty" xml:"DataIngestionIds,omitempty" type:"Repeated"`
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
	LogUserIds []*int64 `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty" type:"Repeated"`
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

func (s UpdateDataBatchIngestionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataBatchIngestionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataBatchIngestionRequest) GetAutoScanNew() *string {
	return s.AutoScanNew
}

func (s *UpdateDataBatchIngestionRequest) GetDataBatchIngestionMode() *string {
	return s.DataBatchIngestionMode
}

func (s *UpdateDataBatchIngestionRequest) GetDataIngestionIds() []*string {
	return s.DataIngestionIds
}

func (s *UpdateDataBatchIngestionRequest) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *UpdateDataBatchIngestionRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataBatchIngestionRequest) GetLogUserIds() []*int64 {
	return s.LogUserIds
}

func (s *UpdateDataBatchIngestionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataBatchIngestionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataBatchIngestionRequest) SetAutoScanNew(v string) *UpdateDataBatchIngestionRequest {
	s.AutoScanNew = &v
	return s
}

func (s *UpdateDataBatchIngestionRequest) SetDataBatchIngestionMode(v string) *UpdateDataBatchIngestionRequest {
	s.DataBatchIngestionMode = &v
	return s
}

func (s *UpdateDataBatchIngestionRequest) SetDataIngestionIds(v []*string) *UpdateDataBatchIngestionRequest {
	s.DataIngestionIds = v
	return s
}

func (s *UpdateDataBatchIngestionRequest) SetDataSourceRecognizeEnabled(v bool) *UpdateDataBatchIngestionRequest {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *UpdateDataBatchIngestionRequest) SetLang(v string) *UpdateDataBatchIngestionRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataBatchIngestionRequest) SetLogUserIds(v []*int64) *UpdateDataBatchIngestionRequest {
	s.LogUserIds = v
	return s
}

func (s *UpdateDataBatchIngestionRequest) SetRegionId(v string) *UpdateDataBatchIngestionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataBatchIngestionRequest) SetRoleFor(v int64) *UpdateDataBatchIngestionRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataBatchIngestionRequest) Validate() error {
	return dara.Validate(s)
}
