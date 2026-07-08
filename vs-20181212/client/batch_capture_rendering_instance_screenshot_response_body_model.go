// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCaptureRenderingInstanceScreenshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatatest(v *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest) *BatchCaptureRenderingInstanceScreenshotResponseBody
	GetDatatest() *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest
	SetFailedCount(v int64) *BatchCaptureRenderingInstanceScreenshotResponseBody
	GetFailedCount() *int64
	SetFailedItems(v []*BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) *BatchCaptureRenderingInstanceScreenshotResponseBody
	GetFailedItems() []*BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems
	SetRequestId(v string) *BatchCaptureRenderingInstanceScreenshotResponseBody
	GetRequestId() *string
	SetSuccessCount(v int64) *BatchCaptureRenderingInstanceScreenshotResponseBody
	GetSuccessCount() *int64
	SetSuccessItems(v []*BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) *BatchCaptureRenderingInstanceScreenshotResponseBody
	GetSuccessItems() []*BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems
}

type BatchCaptureRenderingInstanceScreenshotResponseBody struct {
	// The dry run result.
	Datatest *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest `json:"Datatest,omitempty" xml:"Datatest,omitempty" type:"Struct"`
	// The number of failed instances.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The list of instances for which screenshots failed.
	FailedItems []*BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of successful instances.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The list of successful instances.
	SuccessItems []*BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems `json:"SuccessItems,omitempty" xml:"SuccessItems,omitempty" type:"Repeated"`
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) GetDatatest() *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest {
	return s.Datatest
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) GetFailedItems() []*BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems {
	return s.FailedItems
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) GetSuccessItems() []*BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems {
	return s.SuccessItems
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) SetDatatest(v *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest) *BatchCaptureRenderingInstanceScreenshotResponseBody {
	s.Datatest = v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) SetFailedCount(v int64) *BatchCaptureRenderingInstanceScreenshotResponseBody {
	s.FailedCount = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) SetFailedItems(v []*BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) *BatchCaptureRenderingInstanceScreenshotResponseBody {
	s.FailedItems = v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) SetRequestId(v string) *BatchCaptureRenderingInstanceScreenshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) SetSuccessCount(v int64) *BatchCaptureRenderingInstanceScreenshotResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) SetSuccessItems(v []*BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) *BatchCaptureRenderingInstanceScreenshotResponseBody {
	s.SuccessItems = v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBody) Validate() error {
	if s.Datatest != nil {
		if err := s.Datatest.Validate(); err != nil {
			return err
		}
	}
	if s.FailedItems != nil {
		for _, item := range s.FailedItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessItems != nil {
		for _, item := range s.SuccessItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest struct {
	// The dry run result.
	Result *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest) GetResult() *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult {
	return s.Result
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest) SetResult(v *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult) *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest {
	s.Result = v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatest) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult struct {
	// The number of successful instances.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult) SetSuccessCount(v int64) *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult {
	s.SuccessCount = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyDatatestResult) Validate() error {
	return dara.Validate(s)
}

type BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems struct {
	// The error code of the failure.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message of the failure.
	//
	// example:
	//
	// Not Applied
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The instance ID of the cloud application service instance.
	//
	// example:
	//
	// render-072da95539d3402da90353b244191722
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) GetErrCode() *string {
	return s.ErrCode
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) SetErrCode(v string) *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems {
	s.ErrCode = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) SetErrMessage(v string) *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) SetRenderingInstanceId(v string) *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems {
	s.RenderingInstanceId = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodyFailedItems) Validate() error {
	return dara.Validate(s)
}

type BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems struct {
	// The time when the screenshot was created.
	//
	// example:
	//
	// 2026-05-19T14:46:37+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The download URL of the screenshot.
	//
	// example:
	//
	// https://testts-1.oss-cn-beijing.aliyuncs.com/app/test-zip-file.zip
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The instance ID of the cloud application service instance.
	//
	// example:
	//
	// render-1ada8cd82783407b99fa202826fc6447
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) GetCreationTime() *string {
	return s.CreationTime
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) SetCreationTime(v string) *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems {
	s.CreationTime = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) SetDownloadUrl(v string) *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems {
	s.DownloadUrl = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) SetRenderingInstanceId(v string) *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems {
	s.RenderingInstanceId = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotResponseBodySuccessItems) Validate() error {
	return dara.Validate(s)
}
