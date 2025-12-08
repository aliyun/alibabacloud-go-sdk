// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iMonitorExaminationAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *MonitorExaminationAdvanceRequest
	GetImageURLObject() io.Reader
	SetType(v int64) *MonitorExaminationAdvanceRequest
	GetType() *int64
}

type MonitorExaminationAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/MonitorExamination/1MonitorExamination1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MonitorExaminationAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MonitorExaminationAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *MonitorExaminationAdvanceRequest) GetType() *int64 {
	return s.Type
}

func (s *MonitorExaminationAdvanceRequest) SetImageURLObject(v io.Reader) *MonitorExaminationAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *MonitorExaminationAdvanceRequest) SetType(v int64) *MonitorExaminationAdvanceRequest {
	s.Type = &v
	return s
}

func (s *MonitorExaminationAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
