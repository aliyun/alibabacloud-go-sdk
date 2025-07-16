// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslateReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *GetTranslateReportRequest
	GetApiName() *string
	SetBeginTime(v string) *GetTranslateReportRequest
	GetBeginTime() *string
	SetEndTime(v string) *GetTranslateReportRequest
	GetEndTime() *string
	SetGroup(v string) *GetTranslateReportRequest
	GetGroup() *string
}

type GetTranslateReportRequest struct {
	// This parameter is required.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-03-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-03-01 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s GetTranslateReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateReportRequest) GoString() string {
	return s.String()
}

func (s *GetTranslateReportRequest) GetApiName() *string {
	return s.ApiName
}

func (s *GetTranslateReportRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetTranslateReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTranslateReportRequest) GetGroup() *string {
	return s.Group
}

func (s *GetTranslateReportRequest) SetApiName(v string) *GetTranslateReportRequest {
	s.ApiName = &v
	return s
}

func (s *GetTranslateReportRequest) SetBeginTime(v string) *GetTranslateReportRequest {
	s.BeginTime = &v
	return s
}

func (s *GetTranslateReportRequest) SetEndTime(v string) *GetTranslateReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetTranslateReportRequest) SetGroup(v string) *GetTranslateReportRequest {
	s.Group = &v
	return s
}

func (s *GetTranslateReportRequest) Validate() error {
	return dara.Validate(s)
}
