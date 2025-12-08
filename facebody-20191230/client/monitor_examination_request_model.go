// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorExaminationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *MonitorExaminationRequest
	GetImageURL() *string
	SetType(v int64) *MonitorExaminationRequest
	GetType() *int64
}

type MonitorExaminationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/MonitorExamination/1MonitorExamination1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MonitorExaminationRequest) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationRequest) GoString() string {
	return s.String()
}

func (s *MonitorExaminationRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *MonitorExaminationRequest) GetType() *int64 {
	return s.Type
}

func (s *MonitorExaminationRequest) SetImageURL(v string) *MonitorExaminationRequest {
	s.ImageURL = &v
	return s
}

func (s *MonitorExaminationRequest) SetType(v int64) *MonitorExaminationRequest {
	s.Type = &v
	return s
}

func (s *MonitorExaminationRequest) Validate() error {
	return dara.Validate(s)
}
