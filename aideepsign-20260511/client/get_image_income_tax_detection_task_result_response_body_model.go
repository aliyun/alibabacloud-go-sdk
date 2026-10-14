// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageIncomeTaxDetectionTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*GetImageIncomeTaxDetectionTaskResultResponseBodyItems) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetItems() []*GetImageIncomeTaxDetectionTaskResultResponseBodyItems
	SetMessage(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetRequestId() *string
	SetResult(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetResult() *string
	SetStatus(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody
	GetTaskId() *string
}

type GetImageIncomeTaxDetectionTaskResultResponseBody struct {
	// **The business error code.*	- The value "OK" is returned if the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// **The HTTP status code.*	- The value 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// **The list of detected abnormal items.*	- Returned only when `Status` is `succeeded`. An empty array indicates that no abnormal items are detected.
	Items []*GetImageIncomeTaxDetectionTaskResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The failure reason when the task fails.
	//
	// example:
	//
	// poll timeout: no result within 30m0s
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The raw downstream detection result (extend.result string) for troubleshooting and passing through new fields.
	//
	// example:
	//
	// {"items": []}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// **The task status.*	- Valid values: `running` (in progress), `succeeded` (completed), `failed` (failed).
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// **The task ID.**
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageIncomeTaxDetectionTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageIncomeTaxDetectionTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetItems() []*GetImageIncomeTaxDetectionTaskResultResponseBodyItems {
	return s.Items
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetResult() *string {
	return s.Result
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetCode(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetHttpStatusCode(v int32) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetItems(v []*GetImageIncomeTaxDetectionTaskResultResponseBodyItems) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.Items = v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetMessage(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetRequestId(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetResult(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.Result = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetStatus(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetSuccess(v bool) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) SetTaskId(v string) *GetImageIncomeTaxDetectionTaskResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetImageIncomeTaxDetectionTaskResultResponseBodyItems struct {
	// The text description of the abnormal item.
	//
	// example:
	//
	// Month sequence anomaly: 2022-05 -> 2022-11 (should be in descending or equal order)
	AbnormalItem *string `json:"AbnormalItem,omitempty" xml:"AbnormalItem,omitempty"`
	// The bounding box position of the abnormal item in the image. May be an empty array for image-wide abnormalities.
	Coordinates []*GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	// The detection stage identifier for the abnormality.
	//
	// example:
	//
	// stage2
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
}

func (s GetImageIncomeTaxDetectionTaskResultResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetImageIncomeTaxDetectionTaskResultResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItems) GetAbnormalItem() *string {
	return s.AbnormalItem
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItems) GetCoordinates() []*GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates {
	return s.Coordinates
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItems) GetStage() *string {
	return s.Stage
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItems) SetAbnormalItem(v string) *GetImageIncomeTaxDetectionTaskResultResponseBodyItems {
	s.AbnormalItem = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItems) SetCoordinates(v []*GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) *GetImageIncomeTaxDetectionTaskResultResponseBodyItems {
	s.Coordinates = v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItems) SetStage(v string) *GetImageIncomeTaxDetectionTaskResultResponseBodyItems {
	s.Stage = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItems) Validate() error {
	if s.Coordinates != nil {
		for _, item := range s.Coordinates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates struct {
	// The X-coordinate of the upper-left corner (pixels).
	//
	// example:
	//
	// 928
	X1 *int32 `json:"X1,omitempty" xml:"X1,omitempty"`
	// The X-coordinate of the lower-right corner of the bounding box (pixels).
	//
	// example:
	//
	// 1147
	X2 *int32 `json:"X2,omitempty" xml:"X2,omitempty"`
	// The Y-coordinate of the upper-left corner (pixels).
	//
	// example:
	//
	// 792
	Y1 *int32 `json:"Y1,omitempty" xml:"Y1,omitempty"`
	// The Y-coordinate of the lower-right corner of the bounding box (pixels).
	//
	// example:
	//
	// 847
	Y2 *int32 `json:"Y2,omitempty" xml:"Y2,omitempty"`
}

func (s GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) String() string {
	return dara.Prettify(s)
}

func (s GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) GoString() string {
	return s.String()
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) GetX1() *int32 {
	return s.X1
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) GetX2() *int32 {
	return s.X2
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) GetY1() *int32 {
	return s.Y1
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) GetY2() *int32 {
	return s.Y2
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) SetX1(v int32) *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates {
	s.X1 = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) SetX2(v int32) *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates {
	s.X2 = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) SetY1(v int32) *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates {
	s.Y1 = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) SetY2(v int32) *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates {
	s.Y2 = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponseBodyItemsCoordinates) Validate() error {
	return dara.Validate(s)
}
