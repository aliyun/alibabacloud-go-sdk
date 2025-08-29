// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTargetItemReportDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *QueryTrafficControlTargetItemReportDetailRequest
	GetDate() *string
	SetEnvironment(v string) *QueryTrafficControlTargetItemReportDetailRequest
	GetEnvironment() *string
	SetInstanceId(v string) *QueryTrafficControlTargetItemReportDetailRequest
	GetInstanceId() *string
}

type QueryTrafficControlTargetItemReportDetailRequest struct {
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// This parameter is required.
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) GetDate() *string {
	return s.Date
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) SetDate(v string) *QueryTrafficControlTargetItemReportDetailRequest {
	s.Date = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) SetEnvironment(v string) *QueryTrafficControlTargetItemReportDetailRequest {
	s.Environment = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) SetInstanceId(v string) *QueryTrafficControlTargetItemReportDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailRequest) Validate() error {
	return dara.Validate(s)
}
