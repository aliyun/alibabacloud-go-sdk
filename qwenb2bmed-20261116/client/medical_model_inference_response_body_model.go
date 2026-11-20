// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalModelInferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFrames(v []*MedicalModelInferenceResponseBodyAudioFrames) *MedicalModelInferenceResponseBody
	GetAudioFrames() []*MedicalModelInferenceResponseBodyAudioFrames
	SetCommunication(v *MedicalModelInferenceResponseBodyCommunication) *MedicalModelInferenceResponseBody
	GetCommunication() *MedicalModelInferenceResponseBodyCommunication
	SetData(v *MedicalModelInferenceResponseBodyData) *MedicalModelInferenceResponseBody
	GetData() *MedicalModelInferenceResponseBodyData
	SetError(v *MedicalModelInferenceResponseBodyError) *MedicalModelInferenceResponseBody
	GetError() *MedicalModelInferenceResponseBodyError
	SetErrorCode(v interface{}) *MedicalModelInferenceResponseBody
	GetErrorCode() interface{}
	SetErrorMsg(v string) *MedicalModelInferenceResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *MedicalModelInferenceResponseBody
	GetRequestId() *string
}

type MedicalModelInferenceResponseBody struct {
	// The optional list of audio frames in the response. This can be empty or omitted when no audio is present.
	AudioFrames []*MedicalModelInferenceResponseBodyAudioFrames `json:"audioFrames,omitempty" xml:"audioFrames,omitempty" type:"Repeated"`
	// The information used to associate response frames with the original business request and session.
	Communication *MedicalModelInferenceResponseBodyCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// The business data returned in this response, including messages and optional extended information.
	Data *MedicalModelInferenceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error details returned when a business call fails.
	Error *MedicalModelInferenceResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
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

func (s MedicalModelInferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBody) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBody) GetAudioFrames() []*MedicalModelInferenceResponseBodyAudioFrames {
	return s.AudioFrames
}

func (s *MedicalModelInferenceResponseBody) GetCommunication() *MedicalModelInferenceResponseBodyCommunication {
	return s.Communication
}

func (s *MedicalModelInferenceResponseBody) GetData() *MedicalModelInferenceResponseBodyData {
	return s.Data
}

func (s *MedicalModelInferenceResponseBody) GetError() *MedicalModelInferenceResponseBodyError {
	return s.Error
}

func (s *MedicalModelInferenceResponseBody) GetErrorCode() interface{} {
	return s.ErrorCode
}

func (s *MedicalModelInferenceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MedicalModelInferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBody) SetAudioFrames(v []*MedicalModelInferenceResponseBodyAudioFrames) *MedicalModelInferenceResponseBody {
	s.AudioFrames = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetCommunication(v *MedicalModelInferenceResponseBodyCommunication) *MedicalModelInferenceResponseBody {
	s.Communication = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetData(v *MedicalModelInferenceResponseBodyData) *MedicalModelInferenceResponseBody {
	s.Data = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetError(v *MedicalModelInferenceResponseBodyError) *MedicalModelInferenceResponseBody {
	s.Error = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetErrorCode(v interface{}) *MedicalModelInferenceResponseBody {
	s.ErrorCode = v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetErrorMsg(v string) *MedicalModelInferenceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *MedicalModelInferenceResponseBody) SetRequestId(v string) *MedicalModelInferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBody) Validate() error {
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

type MedicalModelInferenceResponseBodyAudioFrames struct {
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
	// The request correlation ID.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The type of the audio or error frame.
	//
	// example:
	//
	// audio
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalModelInferenceResponseBodyAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyAudioFrames) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetCode() *string {
	return s.Code
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetData() *string {
	return s.Data
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetLast() *bool {
	return s.Last
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetMessage() *string {
	return s.Message
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetCode(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Code = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetData(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Data = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetLast(v bool) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Last = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetMessage(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Message = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetRequestId(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) SetType(v string) *MedicalModelInferenceResponseBodyAudioFrames {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyAudioFrames) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyCommunication struct {
	// The business request ID. This is the same as the business request ID when present.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the business response frame.
	//
	// example:
	//
	// 1
	ResponseId *int64 `json:"responseId,omitempty" xml:"responseId,omitempty"`
	// The ID of the customer session.
	//
	// example:
	//
	// med-session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s MedicalModelInferenceResponseBodyCommunication) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyCommunication) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBodyCommunication) GetResponseId() *int64 {
	return s.ResponseId
}

func (s *MedicalModelInferenceResponseBodyCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *MedicalModelInferenceResponseBodyCommunication) SetRequestId(v string) *MedicalModelInferenceResponseBodyCommunication {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyCommunication) SetResponseId(v int64) *MedicalModelInferenceResponseBodyCommunication {
	s.ResponseId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyCommunication) SetSessionId(v string) *MedicalModelInferenceResponseBodyCommunication {
	s.SessionId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyCommunication) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyData struct {
	// The optional debug information included in the business response.
	//
	// example:
	//
	// {}
	Debug interface{} `json:"debug,omitempty" xml:"debug,omitempty"`
	// The extended business information, which may include inputResources.
	ExtraInfo *MedicalModelInferenceResponseBodyDataExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	// The list of messages returned in this business response or the current streaming response frame.
	Messages []*MedicalModelInferenceResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s MedicalModelInferenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyData) GetDebug() interface{} {
	return s.Debug
}

func (s *MedicalModelInferenceResponseBodyData) GetExtraInfo() *MedicalModelInferenceResponseBodyDataExtraInfo {
	return s.ExtraInfo
}

func (s *MedicalModelInferenceResponseBodyData) GetMessages() []*MedicalModelInferenceResponseBodyDataMessages {
	return s.Messages
}

func (s *MedicalModelInferenceResponseBodyData) SetDebug(v interface{}) *MedicalModelInferenceResponseBodyData {
	s.Debug = v
	return s
}

func (s *MedicalModelInferenceResponseBodyData) SetExtraInfo(v *MedicalModelInferenceResponseBodyDataExtraInfo) *MedicalModelInferenceResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *MedicalModelInferenceResponseBodyData) SetMessages(v []*MedicalModelInferenceResponseBodyDataMessages) *MedicalModelInferenceResponseBodyData {
	s.Messages = v
	return s
}

func (s *MedicalModelInferenceResponseBodyData) Validate() error {
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

type MedicalModelInferenceResponseBodyDataExtraInfo struct {
	// The transfer results of the input resources for this round.
	InputResources []*MedicalModelInferenceResponseBodyDataExtraInfoInputResources `json:"inputResources,omitempty" xml:"inputResources,omitempty" type:"Repeated"`
}

func (s MedicalModelInferenceResponseBodyDataExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfo) GetInputResources() []*MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	return s.InputResources
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfo) SetInputResources(v []*MedicalModelInferenceResponseBodyDataExtraInfoInputResources) *MedicalModelInferenceResponseBodyDataExtraInfo {
	s.InputResources = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfo) Validate() error {
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

type MedicalModelInferenceResponseBodyDataExtraInfoInputResources struct {
	// The index of the original content block.
	//
	// example:
	//
	// 1
	ContentIndex *int32 `json:"contentIndex,omitempty" xml:"contentIndex,omitempty"`
	// The name of the resource file.
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
	// The type of the input resource.
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

func (s MedicalModelInferenceResponseBodyDataExtraInfoInputResources) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetContentIndex() *int32 {
	return s.ContentIndex
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetFileName() *string {
	return s.FileName
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) GetUrl() *string {
	return s.Url
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetContentIndex(v int32) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.ContentIndex = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetFileName(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.FileName = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetResourceId(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.ResourceId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetType(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) SetUrl(v string) *MedicalModelInferenceResponseBodyDataExtraInfoInputResources {
	s.Url = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataExtraInfoInputResources) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyDataMessages struct {
	// The text or content corresponding to the message type.
	//
	// example:
	//
	// Hemoglobin is used to carry oxygen. When interpreting the report, you can refer to the reference ranges provided on the report
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The extended metadata associated with this message.
	MetaData *MedicalModelInferenceResponseBodyDataMessagesMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// The type of the message.
	//
	// example:
	//
	// multi_load/iframe
	MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	// The status of the business message.
	//
	// example:
	//
	// complete
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s MedicalModelInferenceResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetMetaData() *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	return s.MetaData
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetMimeType() *string {
	return s.MimeType
}

func (s *MedicalModelInferenceResponseBodyDataMessages) GetStatus() *string {
	return s.Status
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetContent(v interface{}) *MedicalModelInferenceResponseBodyDataMessages {
	s.Content = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetMetaData(v *MedicalModelInferenceResponseBodyDataMessagesMetaData) *MedicalModelInferenceResponseBodyDataMessages {
	s.MetaData = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetMimeType(v string) *MedicalModelInferenceResponseBodyDataMessages {
	s.MimeType = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) SetStatus(v string) *MedicalModelInferenceResponseBodyDataMessages {
	s.Status = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessages) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MedicalModelInferenceResponseBodyDataMessagesMetaData struct {
	// The title of the status bar.
	//
	// example:
	//
	// Analyzing the issue
	BarTitle *string `json:"barTitle,omitempty" xml:"barTitle,omitempty"`
	// The content of the metadata.
	//
	// example:
	//
	// {"body":"Analyzing the user\\"s issue.","status":"processing"}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The list of structured content associated with the body, such as image-text cards, returned by message type.
	MultiLoad []*MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad `json:"multiLoad,omitempty" xml:"multiLoad,omitempty" type:"Repeated"`
	// The source identifier.
	//
	// example:
	//
	// bar_thinking_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// The type of the metadata.
	//
	// example:
	//
	// bar_thinking
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaData) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaData) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetBarTitle() *string {
	return s.BarTitle
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetContent() interface{} {
	return s.Content
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetMultiLoad() []*MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	return s.MultiLoad
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetBarTitle(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.BarTitle = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetContent(v interface{}) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.Content = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetMultiLoad(v []*MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.MultiLoad = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetSourceSeq(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.SourceSeq = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) SetType(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaData {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaData) Validate() error {
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

type MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad struct {
	// The content of the card.
	//
	// example:
	//
	// {"list":[{"image_url":"https://example.com/image.jpg","image_width":"800","image_height":"600"}]}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The reference identifier of the card.
	//
	// example:
	//
	// image_waterfall_1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// The type of the card.
	//
	// example:
	//
	// image_waterfall
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GetContent() interface{} {
	return s.Content
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) GetType() *string {
	return s.Type
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) SetContent(v interface{}) *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	s.Content = v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) SetSourceSeq(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	s.SourceSeq = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) SetType(v string) *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad {
	s.Type = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyDataMessagesMetaDataMultiLoad) Validate() error {
	return dara.Validate(s)
}

type MedicalModelInferenceResponseBodyError struct {
	// The business error code used by the client to identify the failure type.
	//
	// example:
	//
	// INVALID_REQUEST
	Code interface{} `json:"code,omitempty" xml:"code,omitempty"`
	// The descriptive text of the business error.
	//
	// example:
	//
	// Invalid request parameters
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The business request correlation ID of the failed call, used for troubleshooting.
	//
	// example:
	//
	// med-req-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MedicalModelInferenceResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponseBodyError) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponseBodyError) GetCode() interface{} {
	return s.Code
}

func (s *MedicalModelInferenceResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *MedicalModelInferenceResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *MedicalModelInferenceResponseBodyError) SetCode(v interface{}) *MedicalModelInferenceResponseBodyError {
	s.Code = v
	return s
}

func (s *MedicalModelInferenceResponseBodyError) SetMessage(v string) *MedicalModelInferenceResponseBodyError {
	s.Message = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyError) SetRequestId(v string) *MedicalModelInferenceResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *MedicalModelInferenceResponseBodyError) Validate() error {
	return dara.Validate(s)
}
