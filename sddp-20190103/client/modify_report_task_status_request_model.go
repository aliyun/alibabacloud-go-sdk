// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReportTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *ModifyReportTaskStatusRequest
	GetFeatureType() *int32
	SetLang(v string) *ModifyReportTaskStatusRequest
	GetLang() *string
	SetReportTaskStatus(v int32) *ModifyReportTaskStatusRequest
	GetReportTaskStatus() *int32
}

type ModifyReportTaskStatusRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese
	//
	// - **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The status to set for the report task. Valid values:
	//
	// - **0**: Disable
	//
	// - **1**: Enable
	//
	// 	Notice:
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ReportTaskStatus *int32 `json:"ReportTaskStatus,omitempty" xml:"ReportTaskStatus,omitempty"`
}

func (s ModifyReportTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReportTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyReportTaskStatusRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *ModifyReportTaskStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyReportTaskStatusRequest) GetReportTaskStatus() *int32 {
	return s.ReportTaskStatus
}

func (s *ModifyReportTaskStatusRequest) SetFeatureType(v int32) *ModifyReportTaskStatusRequest {
	s.FeatureType = &v
	return s
}

func (s *ModifyReportTaskStatusRequest) SetLang(v string) *ModifyReportTaskStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyReportTaskStatusRequest) SetReportTaskStatus(v int32) *ModifyReportTaskStatusRequest {
	s.ReportTaskStatus = &v
	return s
}

func (s *ModifyReportTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
