// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListFaceVerifyDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v *DescribeListFaceVerifyDataResponseBodyMonitorData) *DescribeListFaceVerifyDataResponseBody
	GetMonitorData() *DescribeListFaceVerifyDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeListFaceVerifyDataResponseBody
	GetRequestId() *string
}

type DescribeListFaceVerifyDataResponseBody struct {
	// Returned data.
	MonitorData *DescribeListFaceVerifyDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// ID of this request.
	//
	// example:
	//
	// EBD373EA-07FC-50BC-906F-B8950B6ED462
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeListFaceVerifyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyDataResponseBody) GetMonitorData() *DescribeListFaceVerifyDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeListFaceVerifyDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeListFaceVerifyDataResponseBody) SetMonitorData(v *DescribeListFaceVerifyDataResponseBodyMonitorData) *DescribeListFaceVerifyDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBody) SetRequestId(v string) *DescribeListFaceVerifyDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBody) Validate() error {
	if s.MonitorData != nil {
		if err := s.MonitorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeListFaceVerifyDataResponseBodyMonitorData struct {
	// Face verification data.
	FaceVerifyData []*DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData `json:"FaceVerifyData,omitempty" xml:"FaceVerifyData,omitempty" type:"Repeated"`
}

func (s DescribeListFaceVerifyDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorData) GetFaceVerifyData() []*DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData {
	return s.FaceVerifyData
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorData) SetFaceVerifyData(v []*DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) *DescribeListFaceVerifyDataResponseBodyMonitorData {
	s.FaceVerifyData = v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorData) Validate() error {
	if s.FaceVerifyData != nil {
		for _, item := range s.FaceVerifyData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData struct {
	// Verification statistics time.
	//
	// example:
	//
	// 2025-10-16T00:00:00.000Z
	ConDate *string `json:"ConDate,omitempty" xml:"ConDate,omitempty"`
	// Number of failed verifications.
	//
	// example:
	//
	// 6
	FailCnt *string `json:"FailCnt,omitempty" xml:"FailCnt,omitempty"`
	// Verification scheme.
	//
	// example:
	//
	// Liveness
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000011644
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Number of successful verifications.
	//
	// example:
	//
	// 12
	SuccCnt *string `json:"SuccCnt,omitempty" xml:"SuccCnt,omitempty"`
	// Total number of verifications.
	//
	// example:
	//
	// 18
	TotalCnt *string `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) GetConDate() *string {
	return s.ConDate
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) GetFailCnt() *string {
	return s.FailCnt
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) GetName() *string {
	return s.Name
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) GetSuccCnt() *string {
	return s.SuccCnt
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) GetTotalCnt() *string {
	return s.TotalCnt
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) SetConDate(v string) *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData {
	s.ConDate = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) SetFailCnt(v string) *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData {
	s.FailCnt = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) SetName(v string) *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData {
	s.Name = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) SetSceneId(v string) *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData {
	s.SceneId = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) SetSuccCnt(v string) *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData {
	s.SuccCnt = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) SetTotalCnt(v string) *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData {
	s.TotalCnt = &v
	return s
}

func (s *DescribeListFaceVerifyDataResponseBodyMonitorDataFaceVerifyData) Validate() error {
	return dara.Validate(s)
}
