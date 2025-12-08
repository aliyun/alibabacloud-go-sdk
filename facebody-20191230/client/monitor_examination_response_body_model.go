// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorExaminationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MonitorExaminationResponseBodyData) *MonitorExaminationResponseBody
	GetData() *MonitorExaminationResponseBodyData
	SetRequestId(v string) *MonitorExaminationResponseBody
	GetRequestId() *string
}

type MonitorExaminationResponseBody struct {
	Data *MonitorExaminationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D0F6EB94-73B6-4CB8-B266-22D2F221C475
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MonitorExaminationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponseBody) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBody) GetData() *MonitorExaminationResponseBodyData {
	return s.Data
}

func (s *MonitorExaminationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorExaminationResponseBody) SetData(v *MonitorExaminationResponseBodyData) *MonitorExaminationResponseBody {
	s.Data = v
	return s
}

func (s *MonitorExaminationResponseBody) SetRequestId(v string) *MonitorExaminationResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonitorExaminationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MonitorExaminationResponseBodyData struct {
	// example:
	//
	// 0.28805577754974365
	ChatScore  *float32                                      `json:"ChatScore,omitempty" xml:"ChatScore,omitempty"`
	FaceInfo   *MonitorExaminationResponseBodyDataFaceInfo   `json:"FaceInfo,omitempty" xml:"FaceInfo,omitempty" type:"Struct"`
	PersonInfo *MonitorExaminationResponseBodyDataPersonInfo `json:"PersonInfo,omitempty" xml:"PersonInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0.5
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s MonitorExaminationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponseBodyData) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyData) GetChatScore() *float32 {
	return s.ChatScore
}

func (s *MonitorExaminationResponseBodyData) GetFaceInfo() *MonitorExaminationResponseBodyDataFaceInfo {
	return s.FaceInfo
}

func (s *MonitorExaminationResponseBodyData) GetPersonInfo() *MonitorExaminationResponseBodyDataPersonInfo {
	return s.PersonInfo
}

func (s *MonitorExaminationResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *MonitorExaminationResponseBodyData) SetChatScore(v float32) *MonitorExaminationResponseBodyData {
	s.ChatScore = &v
	return s
}

func (s *MonitorExaminationResponseBodyData) SetFaceInfo(v *MonitorExaminationResponseBodyDataFaceInfo) *MonitorExaminationResponseBodyData {
	s.FaceInfo = v
	return s
}

func (s *MonitorExaminationResponseBodyData) SetPersonInfo(v *MonitorExaminationResponseBodyDataPersonInfo) *MonitorExaminationResponseBodyData {
	s.PersonInfo = v
	return s
}

func (s *MonitorExaminationResponseBodyData) SetThreshold(v float32) *MonitorExaminationResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *MonitorExaminationResponseBodyData) Validate() error {
	if s.FaceInfo != nil {
		if err := s.FaceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PersonInfo != nil {
		if err := s.PersonInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MonitorExaminationResponseBodyDataFaceInfo struct {
	// example:
	//
	// 1
	Completeness *float32 `json:"Completeness,omitempty" xml:"Completeness,omitempty"`
	// example:
	//
	// 1
	FaceNumber *int64                                          `json:"FaceNumber,omitempty" xml:"FaceNumber,omitempty"`
	Pose       *MonitorExaminationResponseBodyDataFaceInfoPose `json:"Pose,omitempty" xml:"Pose,omitempty" type:"Struct"`
}

func (s MonitorExaminationResponseBodyDataFaceInfo) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataFaceInfo) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) GetCompleteness() *float32 {
	return s.Completeness
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) GetFaceNumber() *int64 {
	return s.FaceNumber
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) GetPose() *MonitorExaminationResponseBodyDataFaceInfoPose {
	return s.Pose
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) SetCompleteness(v float32) *MonitorExaminationResponseBodyDataFaceInfo {
	s.Completeness = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) SetFaceNumber(v int64) *MonitorExaminationResponseBodyDataFaceInfo {
	s.FaceNumber = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) SetPose(v *MonitorExaminationResponseBodyDataFaceInfoPose) *MonitorExaminationResponseBodyDataFaceInfo {
	s.Pose = v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) Validate() error {
	if s.Pose != nil {
		if err := s.Pose.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MonitorExaminationResponseBodyDataFaceInfoPose struct {
	// example:
	//
	// -0.9185499548912048
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// example:
	//
	// -0.18541647493839264
	Roll *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	// example:
	//
	// 8.095342636108398
	Yaw *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s MonitorExaminationResponseBodyDataFaceInfoPose) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataFaceInfoPose) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) GetPitch() *float32 {
	return s.Pitch
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) GetRoll() *float32 {
	return s.Roll
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) GetYaw() *float32 {
	return s.Yaw
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) SetPitch(v float32) *MonitorExaminationResponseBodyDataFaceInfoPose {
	s.Pitch = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) SetRoll(v float32) *MonitorExaminationResponseBodyDataFaceInfoPose {
	s.Roll = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) SetYaw(v float32) *MonitorExaminationResponseBodyDataFaceInfoPose {
	s.Yaw = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) Validate() error {
	return dara.Validate(s)
}

type MonitorExaminationResponseBodyDataPersonInfo struct {
	CellPhone *MonitorExaminationResponseBodyDataPersonInfoCellPhone `json:"CellPhone,omitempty" xml:"CellPhone,omitempty" type:"Struct"`
	EarPhone  *MonitorExaminationResponseBodyDataPersonInfoEarPhone  `json:"EarPhone,omitempty" xml:"EarPhone,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PersonNumber *int64 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
}

func (s MonitorExaminationResponseBodyDataPersonInfo) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataPersonInfo) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) GetCellPhone() *MonitorExaminationResponseBodyDataPersonInfoCellPhone {
	return s.CellPhone
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) GetEarPhone() *MonitorExaminationResponseBodyDataPersonInfoEarPhone {
	return s.EarPhone
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) GetPersonNumber() *int64 {
	return s.PersonNumber
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) SetCellPhone(v *MonitorExaminationResponseBodyDataPersonInfoCellPhone) *MonitorExaminationResponseBodyDataPersonInfo {
	s.CellPhone = v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) SetEarPhone(v *MonitorExaminationResponseBodyDataPersonInfoEarPhone) *MonitorExaminationResponseBodyDataPersonInfo {
	s.EarPhone = v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) SetPersonNumber(v int64) *MonitorExaminationResponseBodyDataPersonInfo {
	s.PersonNumber = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) Validate() error {
	if s.CellPhone != nil {
		if err := s.CellPhone.Validate(); err != nil {
			return err
		}
	}
	if s.EarPhone != nil {
		if err := s.EarPhone.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MonitorExaminationResponseBodyDataPersonInfoCellPhone struct {
	// example:
	//
	// 0.39076218008995056
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 0.6
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s MonitorExaminationResponseBodyDataPersonInfoCellPhone) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataPersonInfoCellPhone) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataPersonInfoCellPhone) GetScore() *float32 {
	return s.Score
}

func (s *MonitorExaminationResponseBodyDataPersonInfoCellPhone) GetThreshold() *float32 {
	return s.Threshold
}

func (s *MonitorExaminationResponseBodyDataPersonInfoCellPhone) SetScore(v float32) *MonitorExaminationResponseBodyDataPersonInfoCellPhone {
	s.Score = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfoCellPhone) SetThreshold(v float32) *MonitorExaminationResponseBodyDataPersonInfoCellPhone {
	s.Threshold = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfoCellPhone) Validate() error {
	return dara.Validate(s)
}

type MonitorExaminationResponseBodyDataPersonInfoEarPhone struct {
	// example:
	//
	// 0.7980290651321411
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 0.6
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s MonitorExaminationResponseBodyDataPersonInfoEarPhone) String() string {
	return dara.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataPersonInfoEarPhone) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataPersonInfoEarPhone) GetScore() *float32 {
	return s.Score
}

func (s *MonitorExaminationResponseBodyDataPersonInfoEarPhone) GetThreshold() *float32 {
	return s.Threshold
}

func (s *MonitorExaminationResponseBodyDataPersonInfoEarPhone) SetScore(v float32) *MonitorExaminationResponseBodyDataPersonInfoEarPhone {
	s.Score = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfoEarPhone) SetThreshold(v float32) *MonitorExaminationResponseBodyDataPersonInfoEarPhone {
	s.Threshold = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfoEarPhone) Validate() error {
	return dara.Validate(s)
}
