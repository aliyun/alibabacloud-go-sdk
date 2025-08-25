// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v []*string) *ScanImageRequest
	GetScene() []*string
	SetTask(v []*ScanImageRequestTask) *ScanImageRequest
	GetTask() []*ScanImageRequestTask
}

type ScanImageRequest struct {
	// 1
	//
	// This parameter is required.
	//
	// example:
	//
	// porn
	Scene []*string `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
	// 1
	//
	// This parameter is required.
	Task []*ScanImageRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s ScanImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanImageRequest) GoString() string {
	return s.String()
}

func (s *ScanImageRequest) GetScene() []*string {
	return s.Scene
}

func (s *ScanImageRequest) GetTask() []*ScanImageRequestTask {
	return s.Task
}

func (s *ScanImageRequest) SetScene(v []*string) *ScanImageRequest {
	s.Scene = v
	return s
}

func (s *ScanImageRequest) SetTask(v []*ScanImageRequestTask) *ScanImageRequest {
	s.Task = v
	return s
}

func (s *ScanImageRequest) Validate() error {
	return dara.Validate(s)
}

type ScanImageRequestTask struct {
	// example:
	//
	// uuid-xxxx-xxxx-1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// 1
	ImageTimeMillisecond *int64 `json:"ImageTimeMillisecond,omitempty" xml:"ImageTimeMillisecond,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxx.xxx.com/xxx.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 1
	MaxFrames *int32 `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
}

func (s ScanImageRequestTask) String() string {
	return dara.Prettify(s)
}

func (s ScanImageRequestTask) GoString() string {
	return s.String()
}

func (s *ScanImageRequestTask) GetDataId() *string {
	return s.DataId
}

func (s *ScanImageRequestTask) GetImageTimeMillisecond() *int64 {
	return s.ImageTimeMillisecond
}

func (s *ScanImageRequestTask) GetImageURL() *string {
	return s.ImageURL
}

func (s *ScanImageRequestTask) GetInterval() *int32 {
	return s.Interval
}

func (s *ScanImageRequestTask) GetMaxFrames() *int32 {
	return s.MaxFrames
}

func (s *ScanImageRequestTask) SetDataId(v string) *ScanImageRequestTask {
	s.DataId = &v
	return s
}

func (s *ScanImageRequestTask) SetImageTimeMillisecond(v int64) *ScanImageRequestTask {
	s.ImageTimeMillisecond = &v
	return s
}

func (s *ScanImageRequestTask) SetImageURL(v string) *ScanImageRequestTask {
	s.ImageURL = &v
	return s
}

func (s *ScanImageRequestTask) SetInterval(v int32) *ScanImageRequestTask {
	s.Interval = &v
	return s
}

func (s *ScanImageRequestTask) SetMaxFrames(v int32) *ScanImageRequestTask {
	s.MaxFrames = &v
	return s
}

func (s *ScanImageRequestTask) Validate() error {
	return dara.Validate(s)
}
