// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeLicensePlateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeLicensePlateResponseBodyData) *RecognizeLicensePlateResponseBody
	GetData() *RecognizeLicensePlateResponseBodyData
	SetRequestId(v string) *RecognizeLicensePlateResponseBody
	GetRequestId() *string
}

type RecognizeLicensePlateResponseBody struct {
	Data *RecognizeLicensePlateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 3F10DAC3-CF4A-487C-BF33-3B8EB9AA12F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeLicensePlateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBody) GetData() *RecognizeLicensePlateResponseBodyData {
	return s.Data
}

func (s *RecognizeLicensePlateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeLicensePlateResponseBody) SetData(v *RecognizeLicensePlateResponseBodyData) *RecognizeLicensePlateResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeLicensePlateResponseBody) SetRequestId(v string) *RecognizeLicensePlateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeLicensePlateResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeLicensePlateResponseBodyData struct {
	Plates []*RecognizeLicensePlateResponseBodyDataPlates `json:"Plates,omitempty" xml:"Plates,omitempty" type:"Repeated"`
}

func (s RecognizeLicensePlateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyData) GetPlates() []*RecognizeLicensePlateResponseBodyDataPlates {
	return s.Plates
}

func (s *RecognizeLicensePlateResponseBodyData) SetPlates(v []*RecognizeLicensePlateResponseBodyDataPlates) *RecognizeLicensePlateResponseBodyData {
	s.Plates = v
	return s
}

func (s *RecognizeLicensePlateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeLicensePlateResponseBodyDataPlates struct {
	// example:
	//
	// 0.99745339155197144
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	PlateNumber *string  `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	PlateType   *string  `json:"PlateType,omitempty" xml:"PlateType,omitempty"`
	// example:
	//
	// 1
	PlateTypeConfidence *float32                                                `json:"PlateTypeConfidence,omitempty" xml:"PlateTypeConfidence,omitempty"`
	Positions           []*RecognizeLicensePlateResponseBodyDataPlatesPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
	Roi                 *RecognizeLicensePlateResponseBodyDataPlatesRoi         `json:"Roi,omitempty" xml:"Roi,omitempty" type:"Struct"`
}

func (s RecognizeLicensePlateResponseBodyDataPlates) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlates) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) GetConfidence() *float32 {
	return s.Confidence
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) GetPlateNumber() *string {
	return s.PlateNumber
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) GetPlateType() *string {
	return s.PlateType
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) GetPlateTypeConfidence() *float32 {
	return s.PlateTypeConfidence
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) GetPositions() []*RecognizeLicensePlateResponseBodyDataPlatesPositions {
	return s.Positions
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) GetRoi() *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	return s.Roi
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetConfidence(v float32) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Confidence = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateNumber(v string) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateNumber = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateType(v string) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateType = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateTypeConfidence(v float32) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateTypeConfidence = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPositions(v []*RecognizeLicensePlateResponseBodyDataPlatesPositions) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Positions = v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetRoi(v *RecognizeLicensePlateResponseBodyDataPlatesRoi) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Roi = v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) Validate() error {
	return dara.Validate(s)
}

type RecognizeLicensePlateResponseBodyDataPlatesPositions struct {
	// example:
	//
	// 466
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 293
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeLicensePlateResponseBodyDataPlatesPositions) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlatesPositions) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesPositions) GetX() *int64 {
	return s.X
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesPositions) GetY() *int64 {
	return s.Y
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesPositions) SetX(v int64) *RecognizeLicensePlateResponseBodyDataPlatesPositions {
	s.X = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesPositions) SetY(v int64) *RecognizeLicensePlateResponseBodyDataPlatesPositions {
	s.Y = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesPositions) Validate() error {
	return dara.Validate(s)
}

type RecognizeLicensePlateResponseBodyDataPlatesRoi struct {
	// example:
	//
	// 53
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 141
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 294
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 256
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeLicensePlateResponseBodyDataPlatesRoi) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlatesRoi) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) GetH() *int32 {
	return s.H
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) GetW() *int32 {
	return s.W
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) GetX() *int32 {
	return s.X
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) GetY() *int32 {
	return s.Y
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetH(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.H = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetW(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.W = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetX(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.X = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetY(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.Y = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) Validate() error {
	return dara.Validate(s)
}
