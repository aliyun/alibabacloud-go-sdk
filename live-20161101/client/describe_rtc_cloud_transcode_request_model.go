// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcCloudTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRtcCloudTranscodeRequest
	GetAppId() *string
	SetTaskId(v string) *DescribeRtcCloudTranscodeRequest
	GetTaskId() *string
}

type DescribeRtcCloudTranscodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRtcCloudTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudTranscodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudTranscodeRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcCloudTranscodeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRtcCloudTranscodeRequest) SetAppId(v string) *DescribeRtcCloudTranscodeRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeRequest) SetTaskId(v string) *DescribeRtcCloudTranscodeRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
