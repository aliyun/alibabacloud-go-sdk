// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcCloudRecordingFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeRtcCloudRecordingFilesRequest
	GetTaskId() *string
}

type DescribeRtcCloudRecordingFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRtcCloudRecordingFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudRecordingFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudRecordingFilesRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRtcCloudRecordingFilesRequest) SetTaskId(v string) *DescribeRtcCloudRecordingFilesRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRtcCloudRecordingFilesRequest) Validate() error {
	return dara.Validate(s)
}
