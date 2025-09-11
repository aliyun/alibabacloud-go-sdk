// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageTranslateTaskResponseBody
	GetCode() *string
	SetData(v *GetImageTranslateTaskResponseBodyData) *GetImageTranslateTaskResponseBody
	GetData() *GetImageTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *GetImageTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetImageTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageTranslateTaskResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *GetImageTranslateTaskResponseBody
	GetSynchro() *bool
}

type GetImageTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetImageTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 377A48D7-7CFA-53F9-8CA2-14FE3F2774B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// true
	Synchro *bool `json:"synchro,omitempty" xml:"synchro,omitempty"`
}

func (s GetImageTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageTranslateTaskResponseBody) GetData() *GetImageTranslateTaskResponseBodyData {
	return s.Data
}

func (s *GetImageTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetImageTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageTranslateTaskResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetImageTranslateTaskResponseBody) SetCode(v string) *GetImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetData(v *GetImageTranslateTaskResponseBodyData) *GetImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetHttpStatusCode(v string) *GetImageTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetMessage(v string) *GetImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetRequestId(v string) *GetImageTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetSuccess(v bool) *GetImageTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetSynchro(v bool) *GetImageTranslateTaskResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyData struct {
	// example:
	//
	// 213e391517328463424251152ec9fb
	TraceId     *string                                           `json:"traceId,omitempty" xml:"traceId,omitempty"`
	Translation *GetImageTranslateTaskResponseBodyDataTranslation `json:"translation,omitempty" xml:"translation,omitempty" type:"Struct"`
}

func (s GetImageTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyData) GetTraceId() *string {
	return s.TraceId
}

func (s *GetImageTranslateTaskResponseBodyData) GetTranslation() *GetImageTranslateTaskResponseBodyDataTranslation {
	return s.Translation
}

func (s *GetImageTranslateTaskResponseBodyData) SetTraceId(v string) *GetImageTranslateTaskResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyData) SetTranslation(v *GetImageTranslateTaskResponseBodyDataTranslation) *GetImageTranslateTaskResponseBodyData {
	s.Translation = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslation struct {
	// example:
	//
	// 0
	Angle         *int64                                                           `json:"angle,omitempty" xml:"angle,omitempty"`
	BoundingBoxes []*GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes `json:"boundingBoxes,omitempty" xml:"boundingBoxes,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	BoxesCount *int64 `json:"boxesCount,omitempty" xml:"boxesCount,omitempty"`
	// example:
	//
	// 800
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// 800
	OrgHeight *int64 `json:"orgHeight,omitempty" xml:"orgHeight,omitempty"`
	// example:
	//
	// 800
	OrgWidth   *int64                                                        `json:"orgWidth,omitempty" xml:"orgWidth,omitempty"`
	TableInfos []*GetImageTranslateTaskResponseBodyDataTranslationTableInfos `json:"tableInfos,omitempty" xml:"tableInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 800
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslation) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslation) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetAngle() *int64 {
	return s.Angle
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetBoundingBoxes() []*GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	return s.BoundingBoxes
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetBoxesCount() *int64 {
	return s.BoxesCount
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetHeight() *int64 {
	return s.Height
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetOrgHeight() *int64 {
	return s.OrgHeight
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetOrgWidth() *int64 {
	return s.OrgWidth
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetTableInfos() []*GetImageTranslateTaskResponseBodyDataTranslationTableInfos {
	return s.TableInfos
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) GetWidth() *int64 {
	return s.Width
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetAngle(v int64) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.Angle = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetBoundingBoxes(v []*GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.BoundingBoxes = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetBoxesCount(v int64) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.BoxesCount = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetHeight(v int64) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.Height = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetOrgHeight(v int64) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.OrgHeight = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetOrgWidth(v int64) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.OrgWidth = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetTableInfos(v []*GetImageTranslateTaskResponseBodyDataTranslationTableInfos) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.TableInfos = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) SetWidth(v int64) *GetImageTranslateTaskResponseBodyDataTranslation {
	s.Width = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslation) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes struct {
	// example:
	//
	// 1
	Confidence *float32 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// example:
	//
	// 0
	Direction *int64                                                                  `json:"direction,omitempty" xml:"direction,omitempty"`
	DownLeft  *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft  `json:"downLeft,omitempty" xml:"downLeft,omitempty" type:"Struct"`
	DownRight *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight `json:"downRight,omitempty" xml:"downRight,omitempty" type:"Struct"`
	// example:
	//
	// 1
	TableCellId *int64 `json:"tableCellId,omitempty" xml:"tableCellId,omitempty"`
	// example:
	//
	// tbl-1dd15f7e-e373-4da9-858e-8785db1a2954
	TableId *int64  `json:"tableId,omitempty" xml:"tableId,omitempty"`
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// {
	//
	//           "en": "Promotional price:"
	//
	//         }
	Translation map[string]interface{}                                                `json:"translation,omitempty" xml:"translation,omitempty"`
	UpLeft      *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft  `json:"upLeft,omitempty" xml:"upLeft,omitempty" type:"Struct"`
	UpRight     *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight `json:"upRight,omitempty" xml:"upRight,omitempty" type:"Struct"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetConfidence() *float32 {
	return s.Confidence
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetDirection() *int64 {
	return s.Direction
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetDownLeft() *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft {
	return s.DownLeft
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetDownRight() *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight {
	return s.DownRight
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetTableCellId() *int64 {
	return s.TableCellId
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetTableId() *int64 {
	return s.TableId
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetText() *string {
	return s.Text
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetTranslation() map[string]interface{} {
	return s.Translation
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetUpLeft() *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft {
	return s.UpLeft
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) GetUpRight() *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight {
	return s.UpRight
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetConfidence(v float32) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.Confidence = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetDirection(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.Direction = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetDownLeft(v *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.DownLeft = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetDownRight(v *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.DownRight = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetTableCellId(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.TableCellId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetTableId(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.TableId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetText(v string) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.Text = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetTranslation(v map[string]interface{}) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.Translation = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetUpLeft(v *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.UpLeft = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) SetUpRight(v *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes {
	s.UpRight = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxes) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft struct {
	// example:
	//
	// 10
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 694
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) GetX() *int64 {
	return s.X
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) GetY() *int64 {
	return s.Y
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) SetX(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft {
	s.X = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) SetY(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft {
	s.Y = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownLeft) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight struct {
	// example:
	//
	// 97
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 694
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) GetX() *int64 {
	return s.X
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) GetY() *int64 {
	return s.Y
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) SetX(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight {
	s.X = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) SetY(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight {
	s.Y = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesDownRight) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft struct {
	// example:
	//
	// 10
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 669
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) GetX() *int64 {
	return s.X
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) GetY() *int64 {
	return s.Y
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) SetX(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft {
	s.X = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) SetY(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft {
	s.Y = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpLeft) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight struct {
	// example:
	//
	// 11
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 22
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) GetX() *int64 {
	return s.X
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) GetY() *int64 {
	return s.Y
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) SetX(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight {
	s.X = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) SetY(v int64) *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight {
	s.Y = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationBoundingBoxesUpRight) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationTableInfos struct {
	CellInfos []*GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos `json:"cellInfos,omitempty" xml:"cellInfos,omitempty" type:"Repeated"`
	// example:
	//
	// tbl-f16944be-5955-466c-aa6c-940e4ed99a09
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// example:
	//
	// 50
	XCellSize *int64 `json:"xCellSize,omitempty" xml:"xCellSize,omitempty"`
	// example:
	//
	// 50
	YCellSize *int64 `json:"yCellSize,omitempty" xml:"yCellSize,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationTableInfos) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationTableInfos) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) GetCellInfos() []*GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	return s.CellInfos
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) GetTableId() *int64 {
	return s.TableId
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) GetXCellSize() *int64 {
	return s.XCellSize
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) GetYCellSize() *int64 {
	return s.YCellSize
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) SetCellInfos(v []*GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) *GetImageTranslateTaskResponseBodyDataTranslationTableInfos {
	s.CellInfos = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) SetTableId(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfos {
	s.TableId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) SetXCellSize(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfos {
	s.XCellSize = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) SetYCellSize(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfos {
	s.YCellSize = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfos) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos struct {
	Pos []*GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TableCellId *int64  `json:"tableCellId,omitempty" xml:"tableCellId,omitempty"`
	Text        *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// 2
	Xec *int64 `json:"xec,omitempty" xml:"xec,omitempty"`
	// example:
	//
	// 1
	Xsc *int64 `json:"xsc,omitempty" xml:"xsc,omitempty"`
	// example:
	//
	// 1
	Yec *int64 `json:"yec,omitempty" xml:"yec,omitempty"`
	// example:
	//
	// 3
	Ysc *int64 `json:"ysc,omitempty" xml:"ysc,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GetPos() []*GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos {
	return s.Pos
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GetTableCellId() *int64 {
	return s.TableCellId
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GetText() *string {
	return s.Text
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GetXec() *int64 {
	return s.Xec
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GetXsc() *int64 {
	return s.Xsc
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GetYec() *int64 {
	return s.Yec
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) GetYsc() *int64 {
	return s.Ysc
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) SetPos(v []*GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	s.Pos = v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) SetTableCellId(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	s.TableCellId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) SetText(v string) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	s.Text = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) SetXec(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	s.Xec = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) SetXsc(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	s.Xsc = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) SetYec(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	s.Yec = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) SetYsc(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos {
	s.Ysc = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfos) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos struct {
	// example:
	//
	// 33
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 11
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) GetX() *int64 {
	return s.X
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) GetY() *int64 {
	return s.Y
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) SetX(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos {
	s.X = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) SetY(v int64) *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos {
	s.Y = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyDataTranslationTableInfosCellInfosPos) Validate() error {
	return dara.Validate(s)
}
