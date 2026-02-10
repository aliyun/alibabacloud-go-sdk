// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcCloudTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRtcCloudTranscodeResponseBody
	GetRequestId() *string
	SetTaskInfo(v *DescribeRtcCloudTranscodeResponseBodyTaskInfo) *DescribeRtcCloudTranscodeResponseBody
	GetTaskInfo() *DescribeRtcCloudTranscodeResponseBodyTaskInfo
}

type DescribeRtcCloudTranscodeResponseBody struct {
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  *DescribeRtcCloudTranscodeResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeRtcCloudTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcCloudTranscodeResponseBody) GetTaskInfo() *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *DescribeRtcCloudTranscodeResponseBody) SetRequestId(v string) *DescribeRtcCloudTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBody) SetTaskInfo(v *DescribeRtcCloudTranscodeResponseBodyTaskInfo) *DescribeRtcCloudTranscodeResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBody) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcCloudTranscodeResponseBodyTaskInfo struct {
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// myChannel
	ChannelId  *string                                                  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	InputParam *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam `json:"InputParam,omitempty" xml:"InputParam,omitempty" type:"Struct"`
	// example:
	//
	// 600
	MaxIdleTime  *int64                                                       `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	OutputParams []*DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams `json:"OutputParams,omitempty" xml:"OutputParams,omitempty" type:"Repeated"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) GetInputParam() *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam {
	return s.InputParam
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) GetMaxIdleTime() *int64 {
	return s.MaxIdleTime
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) GetOutputParams() []*DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams {
	return s.OutputParams
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) SetAppId(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	s.AppId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) SetChannelId(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) SetInputParam(v *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam) *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	s.InputParam = v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) SetMaxIdleTime(v int64) *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	s.MaxIdleTime = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) SetOutputParams(v []*DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	s.OutputParams = v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) SetStatus(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) SetTaskId(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfo) Validate() error {
	if s.InputParam != nil {
		if err := s.InputParam.Validate(); err != nil {
			return err
		}
	}
	if s.OutputParams != nil {
		for _, item := range s.OutputParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam struct {
	SingleSubUserParam *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam `json:"SingleSubUserParam,omitempty" xml:"SingleSubUserParam,omitempty" type:"Struct"`
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam) GetSingleSubUserParam() *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam {
	return s.SingleSubUserParam
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam) SetSingleSubUserParam(v *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam {
	s.SingleSubUserParam = v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParam) Validate() error {
	if s.SingleSubUserParam != nil {
		if err := s.SingleSubUserParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam struct {
	// example:
	//
	// 0
	SourceType *int64 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType *int64 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) GetSourceType() *int64 {
	return s.SourceType
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) GetStreamType() *int64 {
	return s.StreamType
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) GetUserId() *string {
	return s.UserId
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) SetSourceType(v int64) *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam {
	s.SourceType = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) SetStreamType(v int64) *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam {
	s.StreamType = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) SetUserId(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam {
	s.UserId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoInputParamSingleSubUserParam) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams struct {
	// example:
	//
	// myChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// lhd
	TranscodeTemplate *string `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty"`
	// example:
	//
	// userA_360p
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// eyJhcHBpZCI********
	UserToken *string `json:"UserToken,omitempty" xml:"UserToken,omitempty"`
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) GetTranscodeTemplate() *string {
	return s.TranscodeTemplate
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) GetUserId() *string {
	return s.UserId
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) GetUserToken() *string {
	return s.UserToken
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) SetChannelId(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) SetTranscodeTemplate(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams {
	s.TranscodeTemplate = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) SetUserId(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams {
	s.UserId = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) SetUserToken(v string) *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams {
	s.UserToken = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponseBodyTaskInfoOutputParams) Validate() error {
	return dara.Validate(s)
}
