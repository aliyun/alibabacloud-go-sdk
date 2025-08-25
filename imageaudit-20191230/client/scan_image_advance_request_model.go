// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iScanImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v []*string) *ScanImageAdvanceRequest
	GetScene() []*string
	SetTask(v []*ScanImageAdvanceRequestTask) *ScanImageAdvanceRequest
	GetTask() []*ScanImageAdvanceRequestTask
}

type ScanImageAdvanceRequest struct {
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
	Task []*ScanImageAdvanceRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s ScanImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ScanImageAdvanceRequest) GetScene() []*string {
	return s.Scene
}

func (s *ScanImageAdvanceRequest) GetTask() []*ScanImageAdvanceRequestTask {
	return s.Task
}

func (s *ScanImageAdvanceRequest) SetScene(v []*string) *ScanImageAdvanceRequest {
	s.Scene = v
	return s
}

func (s *ScanImageAdvanceRequest) SetTask(v []*ScanImageAdvanceRequestTask) *ScanImageAdvanceRequest {
	s.Task = v
	return s
}

func (s *ScanImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type ScanImageAdvanceRequestTask struct {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 1
	MaxFrames *int32 `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
}

func (s ScanImageAdvanceRequestTask) String() string {
	return dara.Prettify(s)
}

func (s ScanImageAdvanceRequestTask) GoString() string {
	return s.String()
}

func (s *ScanImageAdvanceRequestTask) GetDataId() *string {
	return s.DataId
}

func (s *ScanImageAdvanceRequestTask) GetImageTimeMillisecond() *int64 {
	return s.ImageTimeMillisecond
}

func (s *ScanImageAdvanceRequestTask) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *ScanImageAdvanceRequestTask) GetInterval() *int32 {
	return s.Interval
}

func (s *ScanImageAdvanceRequestTask) GetMaxFrames() *int32 {
	return s.MaxFrames
}

func (s *ScanImageAdvanceRequestTask) SetDataId(v string) *ScanImageAdvanceRequestTask {
	s.DataId = &v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetImageTimeMillisecond(v int64) *ScanImageAdvanceRequestTask {
	s.ImageTimeMillisecond = &v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetImageURLObject(v io.Reader) *ScanImageAdvanceRequestTask {
	s.ImageURLObject = v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetInterval(v int32) *ScanImageAdvanceRequestTask {
	s.Interval = &v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetMaxFrames(v int32) *ScanImageAdvanceRequestTask {
	s.MaxFrames = &v
	return s
}

func (s *ScanImageAdvanceRequestTask) Validate() error {
	return dara.Validate(s)
}
