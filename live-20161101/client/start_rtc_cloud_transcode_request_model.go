// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartRtcCloudTranscodeRequest
	GetAppId() *string
	SetChannelId(v string) *StartRtcCloudTranscodeRequest
	GetChannelId() *string
	SetInputParam(v *StartRtcCloudTranscodeRequestInputParam) *StartRtcCloudTranscodeRequest
	GetInputParam() *StartRtcCloudTranscodeRequestInputParam
	SetMaxIdleTime(v int64) *StartRtcCloudTranscodeRequest
	GetMaxIdleTime() *int64
	SetOutputParams(v []*StartRtcCloudTranscodeRequestOutputParams) *StartRtcCloudTranscodeRequest
	GetOutputParams() []*StartRtcCloudTranscodeRequestOutputParams
}

type StartRtcCloudTranscodeRequest struct {
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
	// myChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	InputParam *StartRtcCloudTranscodeRequestInputParam `json:"InputParam,omitempty" xml:"InputParam,omitempty" type:"Struct"`
	// example:
	//
	// 600
	MaxIdleTime *int64 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// This parameter is required.
	OutputParams []*StartRtcCloudTranscodeRequestOutputParams `json:"OutputParams,omitempty" xml:"OutputParams,omitempty" type:"Repeated"`
}

func (s StartRtcCloudTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudTranscodeRequest) GoString() string {
	return s.String()
}

func (s *StartRtcCloudTranscodeRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartRtcCloudTranscodeRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRtcCloudTranscodeRequest) GetInputParam() *StartRtcCloudTranscodeRequestInputParam {
	return s.InputParam
}

func (s *StartRtcCloudTranscodeRequest) GetMaxIdleTime() *int64 {
	return s.MaxIdleTime
}

func (s *StartRtcCloudTranscodeRequest) GetOutputParams() []*StartRtcCloudTranscodeRequestOutputParams {
	return s.OutputParams
}

func (s *StartRtcCloudTranscodeRequest) SetAppId(v string) *StartRtcCloudTranscodeRequest {
	s.AppId = &v
	return s
}

func (s *StartRtcCloudTranscodeRequest) SetChannelId(v string) *StartRtcCloudTranscodeRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRtcCloudTranscodeRequest) SetInputParam(v *StartRtcCloudTranscodeRequestInputParam) *StartRtcCloudTranscodeRequest {
	s.InputParam = v
	return s
}

func (s *StartRtcCloudTranscodeRequest) SetMaxIdleTime(v int64) *StartRtcCloudTranscodeRequest {
	s.MaxIdleTime = &v
	return s
}

func (s *StartRtcCloudTranscodeRequest) SetOutputParams(v []*StartRtcCloudTranscodeRequestOutputParams) *StartRtcCloudTranscodeRequest {
	s.OutputParams = v
	return s
}

func (s *StartRtcCloudTranscodeRequest) Validate() error {
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

type StartRtcCloudTranscodeRequestInputParam struct {
	// This parameter is required.
	SingleSubUserParam *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam `json:"SingleSubUserParam,omitempty" xml:"SingleSubUserParam,omitempty" type:"Struct"`
}

func (s StartRtcCloudTranscodeRequestInputParam) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudTranscodeRequestInputParam) GoString() string {
	return s.String()
}

func (s *StartRtcCloudTranscodeRequestInputParam) GetSingleSubUserParam() *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam {
	return s.SingleSubUserParam
}

func (s *StartRtcCloudTranscodeRequestInputParam) SetSingleSubUserParam(v *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) *StartRtcCloudTranscodeRequestInputParam {
	s.SingleSubUserParam = v
	return s
}

func (s *StartRtcCloudTranscodeRequestInputParam) Validate() error {
	if s.SingleSubUserParam != nil {
		if err := s.SingleSubUserParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartRtcCloudTranscodeRequestInputParamSingleSubUserParam struct {
	// example:
	//
	// 0
	SourceType *int64 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType *int64 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) GoString() string {
	return s.String()
}

func (s *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) GetSourceType() *int64 {
	return s.SourceType
}

func (s *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) GetStreamType() *int64 {
	return s.StreamType
}

func (s *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) GetUserId() *string {
	return s.UserId
}

func (s *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) SetSourceType(v int64) *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam {
	s.SourceType = &v
	return s
}

func (s *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) SetStreamType(v int64) *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam {
	s.StreamType = &v
	return s
}

func (s *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) SetUserId(v string) *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam {
	s.UserId = &v
	return s
}

func (s *StartRtcCloudTranscodeRequestInputParamSingleSubUserParam) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudTranscodeRequestOutputParams struct {
	// This parameter is required.
	//
	// example:
	//
	// myChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lhd
	TranscodeTemplate *string `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// userA_360p
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eyJhcHBpZCI********
	UserToken *string `json:"UserToken,omitempty" xml:"UserToken,omitempty"`
}

func (s StartRtcCloudTranscodeRequestOutputParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudTranscodeRequestOutputParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudTranscodeRequestOutputParams) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRtcCloudTranscodeRequestOutputParams) GetTranscodeTemplate() *string {
	return s.TranscodeTemplate
}

func (s *StartRtcCloudTranscodeRequestOutputParams) GetUserId() *string {
	return s.UserId
}

func (s *StartRtcCloudTranscodeRequestOutputParams) GetUserToken() *string {
	return s.UserToken
}

func (s *StartRtcCloudTranscodeRequestOutputParams) SetChannelId(v string) *StartRtcCloudTranscodeRequestOutputParams {
	s.ChannelId = &v
	return s
}

func (s *StartRtcCloudTranscodeRequestOutputParams) SetTranscodeTemplate(v string) *StartRtcCloudTranscodeRequestOutputParams {
	s.TranscodeTemplate = &v
	return s
}

func (s *StartRtcCloudTranscodeRequestOutputParams) SetUserId(v string) *StartRtcCloudTranscodeRequestOutputParams {
	s.UserId = &v
	return s
}

func (s *StartRtcCloudTranscodeRequestOutputParams) SetUserToken(v string) *StartRtcCloudTranscodeRequestOutputParams {
	s.UserToken = &v
	return s
}

func (s *StartRtcCloudTranscodeRequestOutputParams) Validate() error {
	return dara.Validate(s)
}
