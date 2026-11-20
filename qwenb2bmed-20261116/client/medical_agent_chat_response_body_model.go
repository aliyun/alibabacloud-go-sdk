// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAgentChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFrames(v []*MedicalAgentChatResponseBodyAudioFrames) *MedicalAgentChatResponseBody
	GetAudioFrames() []*MedicalAgentChatResponseBodyAudioFrames
	SetCommunication(v *MedicalAgentChatResponseBodyCommunication) *MedicalAgentChatResponseBody
	GetCommunication() *MedicalAgentChatResponseBodyCommunication
	SetData(v *MedicalAgentChatResponseBodyData) *MedicalAgentChatResponseBody
	GetData() *MedicalAgentChatResponseBodyData
	SetError(v *MedicalAgentChatResponseBodyError) *MedicalAgentChatResponseBody
	GetError() *MedicalAgentChatResponseBodyError
	SetErrorCode(v interface{}) *MedicalAgentChatResponseBody
	GetErrorCode() interface{}
	SetErrorMsg(v string) *MedicalAgentChatResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *MedicalAgentChatResponseBody
	GetRequestId() *string
}

type MedicalAgentChatResponseBody struct {
	// The optional list of audio frames in the response. This list can be empty or omitted when no audio is present.
	AudioFrames []*MedicalAgentChatResponseBodyAudioFrames `json:"audioFrames,omitempty" xml:"audioFrames,omitempty" type:"Repeated"`
	// The information used to correlate response frames with the original business request and session.
	Communication *MedicalAgentChatResponseBodyCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// The business data returned in this response, including messages and optional extended information.
	Data *MedicalAgentChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error details returned when a business call fails.
	Error *MedicalAgentChatResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The business status code.
	//
	// example:
	//
	// 0
	ErrorCode interface{} `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business status description.
	//
	// example:
	//
	// success
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The optional business request correlation ID. For the POP request ID, see x-acs-request-id.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MedicalAgentChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBody) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBody) GetAudioFrames() []*MedicalAgentChatResponseBodyAudioFrames {
	return s.AudioFrames
}

func (s *MedicalAgentChatResponseBody) GetCommunication() *MedicalAgentChatResponseBodyCommunication {
	return s.Communication
}

func (s *MedicalAgentChatResponseBody) GetData() *MedicalAgentChatResponseBodyData {
	return s.Data
}

func (s *MedicalAgentChatResponseBody) GetError() *MedicalAgentChatResponseBodyError {
	return s.Error
}

func (s *MedicalAgentChatResponseBody) GetErrorCode() interface{} {
	return s.ErrorCode
}

func (s *MedicalAgentChatResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MedicalAgentChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBody) SetAudioFrames(v []*MedicalAgentChatResponseBodyAudioFrames) *MedicalAgentChatResponseBody {
	s.AudioFrames = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetCommunication(v *MedicalAgentChatResponseBodyCommunication) *MedicalAgentChatResponseBody {
	s.Communication = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetData(v *MedicalAgentChatResponseBodyData) *MedicalAgentChatResponseBody {
	s.Data = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetError(v *MedicalAgentChatResponseBodyError) *MedicalAgentChatResponseBody {
	s.Error = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetErrorCode(v interface{}) *MedicalAgentChatResponseBody {
	s.ErrorCode = v
	return s
}

func (s *MedicalAgentChatResponseBody) SetErrorMsg(v string) *MedicalAgentChatResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *MedicalAgentChatResponseBody) SetRequestId(v string) *MedicalAgentChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBody) Validate() error {
	if s.AudioFrames != nil {
		for _, item := range s.AudioFrames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Communication != nil {
		if err := s.Communication.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MedicalAgentChatResponseBodyAudioFrames struct {
	// The error code.
	//
	// example:
	//
	// DOWNSTREAM_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The encoded audio data.
	//
	// example:
	//
	// AAEC
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Indicates whether this is the last audio frame.
	Last *bool `json:"last,omitempty" xml:"last,omitempty"`
	// The error description.
	//
	// example:
	//
	// Audio generation failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The business request correlation ID.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The audio or error frame type.
	//
	// example:
	//
	// audio
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalAgentChatResponseBodyAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyAudioFrames) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetCode() *string {
	return s.Code
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetData() *string {
	return s.Data
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetLast() *bool {
	return s.Last
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetMessage() *string {
	return s.Message
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBodyAudioFrames) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetCode(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Code = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetData(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Data = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetLast(v bool) *MedicalAgentChatResponseBodyAudioFrames {
	s.Last = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetMessage(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Message = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetRequestId(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) SetType(v string) *MedicalAgentChatResponseBodyAudioFrames {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyAudioFrames) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyCommunication struct {
	// The business request ID, which is identical to the original request ID when present.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The business response frame ID.
	//
	// example:
	//
	// 1
	ResponseId *int64 `json:"responseId,omitempty" xml:"responseId,omitempty"`
	// The customer session ID.
	//
	// example:
	//
	// med-session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s MedicalAgentChatResponseBodyCommunication) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyCommunication) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBodyCommunication) GetResponseId() *int64 {
	return s.ResponseId
}

func (s *MedicalAgentChatResponseBodyCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *MedicalAgentChatResponseBodyCommunication) SetRequestId(v string) *MedicalAgentChatResponseBodyCommunication {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyCommunication) SetResponseId(v int64) *MedicalAgentChatResponseBodyCommunication {
	s.ResponseId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyCommunication) SetSessionId(v string) *MedicalAgentChatResponseBodyCommunication {
	s.SessionId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyCommunication) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyData struct {
	// The optional debug information included in the business response.
	//
	// example:
	//
	// {}
	Debug interface{} `json:"debug,omitempty" xml:"debug,omitempty"`
	// The business extended information, which may include inputResources.
	ExtraInfo *MedicalAgentChatResponseBodyDataExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	// The list of messages returned in this business response or the current streaming response frame.
	Messages []*MedicalAgentChatResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s MedicalAgentChatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyData) GetDebug() interface{} {
	return s.Debug
}

func (s *MedicalAgentChatResponseBodyData) GetExtraInfo() *MedicalAgentChatResponseBodyDataExtraInfo {
	return s.ExtraInfo
}

func (s *MedicalAgentChatResponseBodyData) GetMessages() []*MedicalAgentChatResponseBodyDataMessages {
	return s.Messages
}

func (s *MedicalAgentChatResponseBodyData) SetDebug(v interface{}) *MedicalAgentChatResponseBodyData {
	s.Debug = v
	return s
}

func (s *MedicalAgentChatResponseBodyData) SetExtraInfo(v *MedicalAgentChatResponseBodyDataExtraInfo) *MedicalAgentChatResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *MedicalAgentChatResponseBodyData) SetMessages(v []*MedicalAgentChatResponseBodyDataMessages) *MedicalAgentChatResponseBodyData {
	s.Messages = v
	return s
}

func (s *MedicalAgentChatResponseBodyData) Validate() error {
	if s.ExtraInfo != nil {
		if err := s.ExtraInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MedicalAgentChatResponseBodyDataExtraInfo struct {
	// The transfer results of the input resources for this round.
	InputResources []*MedicalAgentChatResponseBodyDataExtraInfoInputResources `json:"inputResources,omitempty" xml:"inputResources,omitempty" type:"Repeated"`
}

func (s MedicalAgentChatResponseBodyDataExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataExtraInfo) GetInputResources() []*MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	return s.InputResources
}

func (s *MedicalAgentChatResponseBodyDataExtraInfo) SetInputResources(v []*MedicalAgentChatResponseBodyDataExtraInfoInputResources) *MedicalAgentChatResponseBodyDataExtraInfo {
	s.InputResources = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfo) Validate() error {
	if s.InputResources != nil {
		for _, item := range s.InputResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MedicalAgentChatResponseBodyDataExtraInfoInputResources struct {
	// The index of the original content block.
	//
	// example:
	//
	// 1
	ContentIndex *int32 `json:"contentIndex,omitempty" xml:"contentIndex,omitempty"`
	// The resource file name.
	//
	// example:
	//
	// report.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The ID of the transferred resource.
	//
	// example:
	//
	// ws_123
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The input resource type.
	//
	// example:
	//
	// image
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The URL of the transferred resource.
	//
	// example:
	//
	// https://example.com/report.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MedicalAgentChatResponseBodyDataExtraInfoInputResources) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataExtraInfoInputResources) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetContentIndex() *int32 {
	return s.ContentIndex
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetFileName() *string {
	return s.FileName
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) GetUrl() *string {
	return s.Url
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetContentIndex(v int32) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.ContentIndex = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetFileName(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.FileName = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetResourceId(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.ResourceId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetType(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) SetUrl(v string) *MedicalAgentChatResponseBodyDataExtraInfoInputResources {
	s.Url = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataExtraInfoInputResources) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyDataMessages struct {
	// The text or content corresponding to the message type.
	//
	// example:
	//
	// Hemoglobin carries oxygen. Refer to the reference ranges on the report for interpretation
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The extended metadata associated with this message.
	MetaData *MedicalAgentChatResponseBodyDataMessagesMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// The message type.
	//
	// example:
	//
	// multi_load/iframe
	MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	// The business message status.
	//
	// example:
	//
	// complete
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s MedicalAgentChatResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetMetaData() *MedicalAgentChatResponseBodyDataMessagesMetaData {
	return s.MetaData
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetMimeType() *string {
	return s.MimeType
}

func (s *MedicalAgentChatResponseBodyDataMessages) GetStatus() *string {
	return s.Status
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetContent(v interface{}) *MedicalAgentChatResponseBodyDataMessages {
	s.Content = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetMetaData(v *MedicalAgentChatResponseBodyDataMessagesMetaData) *MedicalAgentChatResponseBodyDataMessages {
	s.MetaData = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetMimeType(v string) *MedicalAgentChatResponseBodyDataMessages {
	s.MimeType = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) SetStatus(v string) *MedicalAgentChatResponseBodyDataMessages {
	s.Status = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessages) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MedicalAgentChatResponseBodyDataMessagesMetaData struct {
	// The status bar title.
	//
	// example:
	//
	// Analyzing the issue
	BarTitle *string `json:"barTitle,omitempty" xml:"barTitle,omitempty"`
	// The metadata content.
	//
	// example:
	//
	// {"body":"Analyzing the user\\"s issue.","status":"processing"}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The list of structured content associated with the body, such as image-text cards, returned by message type.
	MultiLoad []*MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad `json:"multiLoad,omitempty" xml:"multiLoad,omitempty" type:"Repeated"`
	// The source identifier.
	//
	// example:
	//
	// bar_thinking_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// The metadata type.
	//
	// example:
	//
	// bar_thinking
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalAgentChatResponseBodyDataMessagesMetaData) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataMessagesMetaData) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetBarTitle() *string {
	return s.BarTitle
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetContent() interface{} {
	return s.Content
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetMultiLoad() []*MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	return s.MultiLoad
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetBarTitle(v string) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.BarTitle = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetContent(v interface{}) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.Content = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetMultiLoad(v []*MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.MultiLoad = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetSourceSeq(v string) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.SourceSeq = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) SetType(v string) *MedicalAgentChatResponseBodyDataMessagesMetaData {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaData) Validate() error {
	if s.MultiLoad != nil {
		for _, item := range s.MultiLoad {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad struct {
	// The card content.
	//
	// example:
	//
	// {"list":[{"image_url":"https://example.com/image.jpg","image_width":"800","image_height":"600"}]}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The card reference identifier.
	//
	// example:
	//
	// image_waterfall_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// The card type.
	//
	// example:
	//
	// image_waterfall
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GetContent() interface{} {
	return s.Content
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) GetType() *string {
	return s.Type
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) SetContent(v interface{}) *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Content = v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) SetSourceSeq(v string) *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.SourceSeq = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) SetType(v string) *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Type = &v
	return s
}

func (s *MedicalAgentChatResponseBodyDataMessagesMetaDataMultiLoad) Validate() error {
	return dara.Validate(s)
}

type MedicalAgentChatResponseBodyError struct {
	// The business error code used by the client to identify the failure type.
	//
	// example:
	//
	// INVALID_REQUEST
	Code interface{} `json:"code,omitempty" xml:"code,omitempty"`
	// The descriptive text for the business error.
	//
	// example:
	//
	// Invalid request parameters
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The business request correlation ID for the failed call, used to troubleshoot the issue.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MedicalAgentChatResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponseBodyError) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponseBodyError) GetCode() interface{} {
	return s.Code
}

func (s *MedicalAgentChatResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *MedicalAgentChatResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalAgentChatResponseBodyError) SetCode(v interface{}) *MedicalAgentChatResponseBodyError {
	s.Code = v
	return s
}

func (s *MedicalAgentChatResponseBodyError) SetMessage(v string) *MedicalAgentChatResponseBodyError {
	s.Message = &v
	return s
}

func (s *MedicalAgentChatResponseBodyError) SetRequestId(v string) *MedicalAgentChatResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *MedicalAgentChatResponseBodyError) Validate() error {
	return dara.Validate(s)
}
