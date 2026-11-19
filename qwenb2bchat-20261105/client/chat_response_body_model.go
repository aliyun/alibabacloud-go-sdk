// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioFrames(v []*ChatResponseBodyAudioFrames) *ChatResponseBody
	GetAudioFrames() []*ChatResponseBodyAudioFrames
	SetCommunication(v *ChatResponseBodyCommunication) *ChatResponseBody
	GetCommunication() *ChatResponseBodyCommunication
	SetData(v *ChatResponseBodyData) *ChatResponseBody
	GetData() *ChatResponseBodyData
	SetError(v *ChatResponseBodyError) *ChatResponseBody
	GetError() *ChatResponseBodyError
	SetErrorCode(v interface{}) *ChatResponseBody
	GetErrorCode() interface{}
	SetErrorMsg(v string) *ChatResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ChatResponseBody
	GetRequestId() *string
}

type ChatResponseBody struct {
	// The optional list of audio frames in the response. This list can be empty or omitted when no audio is present.
	AudioFrames []*ChatResponseBodyAudioFrames `json:"audioFrames,omitempty" xml:"audioFrames,omitempty" type:"Repeated"`
	// The information used to associate the response frame with the original business request and session.
	Communication *ChatResponseBodyCommunication `json:"communication,omitempty" xml:"communication,omitempty" type:"Struct"`
	// The business data returned in the response, including messages and optional extended information.
	Data *ChatResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error details returned when a business call fails.
	Error *ChatResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
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
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) GetAudioFrames() []*ChatResponseBodyAudioFrames {
	return s.AudioFrames
}

func (s *ChatResponseBody) GetCommunication() *ChatResponseBodyCommunication {
	return s.Communication
}

func (s *ChatResponseBody) GetData() *ChatResponseBodyData {
	return s.Data
}

func (s *ChatResponseBody) GetError() *ChatResponseBodyError {
	return s.Error
}

func (s *ChatResponseBody) GetErrorCode() interface{} {
	return s.ErrorCode
}

func (s *ChatResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBody) SetAudioFrames(v []*ChatResponseBodyAudioFrames) *ChatResponseBody {
	s.AudioFrames = v
	return s
}

func (s *ChatResponseBody) SetCommunication(v *ChatResponseBodyCommunication) *ChatResponseBody {
	s.Communication = v
	return s
}

func (s *ChatResponseBody) SetData(v *ChatResponseBodyData) *ChatResponseBody {
	s.Data = v
	return s
}

func (s *ChatResponseBody) SetError(v *ChatResponseBodyError) *ChatResponseBody {
	s.Error = v
	return s
}

func (s *ChatResponseBody) SetErrorCode(v interface{}) *ChatResponseBody {
	s.ErrorCode = v
	return s
}

func (s *ChatResponseBody) SetErrorMsg(v string) *ChatResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) Validate() error {
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

type ChatResponseBodyAudioFrames struct {
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
	// SUQz
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Indicates whether the frame is the last audio frame.
	Last *bool `json:"last,omitempty" xml:"last,omitempty"`
	// The error description.
	//
	// example:
	//
	// Speech synthesis is currently unavailable
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The business request correlation ID.
	//
	// example:
	//
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The type of the audio or error frame.
	//
	// example:
	//
	// audio
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChatResponseBodyAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyAudioFrames) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyAudioFrames) GetCode() *string {
	return s.Code
}

func (s *ChatResponseBodyAudioFrames) GetData() *string {
	return s.Data
}

func (s *ChatResponseBodyAudioFrames) GetLast() *bool {
	return s.Last
}

func (s *ChatResponseBodyAudioFrames) GetMessage() *string {
	return s.Message
}

func (s *ChatResponseBodyAudioFrames) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBodyAudioFrames) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyAudioFrames) SetCode(v string) *ChatResponseBodyAudioFrames {
	s.Code = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetData(v string) *ChatResponseBodyAudioFrames {
	s.Data = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetLast(v bool) *ChatResponseBodyAudioFrames {
	s.Last = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetMessage(v string) *ChatResponseBodyAudioFrames {
	s.Message = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetRequestId(v string) *ChatResponseBodyAudioFrames {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) SetType(v string) *ChatResponseBodyAudioFrames {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyAudioFrames) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyCommunication struct {
	// The business request ID, if present.
	//
	// example:
	//
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The identifier of the business response frame.
	//
	// example:
	//
	// 1
	ResponseId *int64 `json:"responseId,omitempty" xml:"responseId,omitempty"`
	// The customer session ID.
	//
	// example:
	//
	// session-example-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ChatResponseBodyCommunication) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyCommunication) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyCommunication) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBodyCommunication) GetResponseId() *int64 {
	return s.ResponseId
}

func (s *ChatResponseBodyCommunication) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatResponseBodyCommunication) SetRequestId(v string) *ChatResponseBodyCommunication {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBodyCommunication) SetResponseId(v int64) *ChatResponseBodyCommunication {
	s.ResponseId = &v
	return s
}

func (s *ChatResponseBodyCommunication) SetSessionId(v string) *ChatResponseBodyCommunication {
	s.SessionId = &v
	return s
}

func (s *ChatResponseBodyCommunication) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyData struct {
	// The optional debug information included in the business response.
	//
	// example:
	//
	// {"intentLabel":"image_search"}
	Debug interface{} `json:"debug,omitempty" xml:"debug,omitempty"`
	// The business extended information, which may include inputResources.
	ExtraInfo *ChatResponseBodyDataExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	// The list of messages returned in the business response or the current streaming response frame.
	Messages []*ChatResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s ChatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyData) GetDebug() interface{} {
	return s.Debug
}

func (s *ChatResponseBodyData) GetExtraInfo() *ChatResponseBodyDataExtraInfo {
	return s.ExtraInfo
}

func (s *ChatResponseBodyData) GetMessages() []*ChatResponseBodyDataMessages {
	return s.Messages
}

func (s *ChatResponseBodyData) SetDebug(v interface{}) *ChatResponseBodyData {
	s.Debug = v
	return s
}

func (s *ChatResponseBodyData) SetExtraInfo(v *ChatResponseBodyDataExtraInfo) *ChatResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *ChatResponseBodyData) SetMessages(v []*ChatResponseBodyDataMessages) *ChatResponseBodyData {
	s.Messages = v
	return s
}

func (s *ChatResponseBodyData) Validate() error {
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

type ChatResponseBodyDataExtraInfo struct {
	// The transfer and storage results of the input resources for the current round.
	InputResources []*ChatResponseBodyDataExtraInfoInputResources `json:"inputResources,omitempty" xml:"inputResources,omitempty" type:"Repeated"`
}

func (s ChatResponseBodyDataExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataExtraInfo) GetInputResources() []*ChatResponseBodyDataExtraInfoInputResources {
	return s.InputResources
}

func (s *ChatResponseBodyDataExtraInfo) SetInputResources(v []*ChatResponseBodyDataExtraInfoInputResources) *ChatResponseBodyDataExtraInfo {
	s.InputResources = v
	return s
}

func (s *ChatResponseBodyDataExtraInfo) Validate() error {
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

type ChatResponseBodyDataExtraInfoInputResources struct {
	// The position of the original content block.
	//
	// example:
	//
	// 1
	ContentIndex *int32 `json:"contentIndex,omitempty" xml:"contentIndex,omitempty"`
	// The file name of the transferred and stored resource.
	//
	// example:
	//
	// cat.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The ID of the transferred and stored resource.
	//
	// example:
	//
	// ws_example_001
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the transferred and stored resource.
	//
	// example:
	//
	// image
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The URL of the transferred and stored resource.
	//
	// example:
	//
	// https://cdn.example.com/uploads/cat.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ChatResponseBodyDataExtraInfoInputResources) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataExtraInfoInputResources) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetContentIndex() *int32 {
	return s.ContentIndex
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetFileName() *string {
	return s.FileName
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyDataExtraInfoInputResources) GetUrl() *string {
	return s.Url
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetContentIndex(v int32) *ChatResponseBodyDataExtraInfoInputResources {
	s.ContentIndex = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetFileName(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.FileName = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetResourceId(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.ResourceId = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetType(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) SetUrl(v string) *ChatResponseBodyDataExtraInfoInputResources {
	s.Url = &v
	return s
}

func (s *ChatResponseBodyDataExtraInfoInputResources) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyDataMessages struct {
	// The text or content of the corresponding message type.
	//
	// example:
	//
	// Image as follows [(image_waterfall_1)]
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The extended metadata associated with the message.
	MetaData *ChatResponseBodyDataMessagesMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
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

func (s ChatResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataMessages) GetContent() interface{} {
	return s.Content
}

func (s *ChatResponseBodyDataMessages) GetMetaData() *ChatResponseBodyDataMessagesMetaData {
	return s.MetaData
}

func (s *ChatResponseBodyDataMessages) GetMimeType() *string {
	return s.MimeType
}

func (s *ChatResponseBodyDataMessages) GetStatus() *string {
	return s.Status
}

func (s *ChatResponseBodyDataMessages) SetContent(v interface{}) *ChatResponseBodyDataMessages {
	s.Content = v
	return s
}

func (s *ChatResponseBodyDataMessages) SetMetaData(v *ChatResponseBodyDataMessagesMetaData) *ChatResponseBodyDataMessages {
	s.MetaData = v
	return s
}

func (s *ChatResponseBodyDataMessages) SetMimeType(v string) *ChatResponseBodyDataMessages {
	s.MimeType = &v
	return s
}

func (s *ChatResponseBodyDataMessages) SetStatus(v string) *ChatResponseBodyDataMessages {
	s.Status = &v
	return s
}

func (s *ChatResponseBodyDataMessages) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatResponseBodyDataMessagesMetaData struct {
	// The display title of the process module.
	//
	// example:
	//
	// Issue analysis
	BarTitle *string `json:"barTitle,omitempty" xml:"barTitle,omitempty"`
	// The complete content snapshot of the process module.
	//
	// example:
	//
	// {"body":"Analyzing the issue","status":"processing"}
	Content interface{} `json:"content,omitempty" xml:"content,omitempty"`
	// The list of structured content associated with the body, such as image-text cards, returned by message type.
	MultiLoad []*ChatResponseBodyDataMessagesMetaDataMultiLoad `json:"multiLoad,omitempty" xml:"multiLoad,omitempty" type:"Repeated"`
	// The process module identifier within the current request.
	//
	// example:
	//
	// thinking-1
	SourceSeq *string `json:"sourceSeq,omitempty" xml:"sourceSeq,omitempty"`
	// The process module type.
	//
	// example:
	//
	// bar_thinking
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChatResponseBodyDataMessagesMetaData) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataMessagesMetaData) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataMessagesMetaData) GetBarTitle() *string {
	return s.BarTitle
}

func (s *ChatResponseBodyDataMessagesMetaData) GetContent() interface{} {
	return s.Content
}

func (s *ChatResponseBodyDataMessagesMetaData) GetMultiLoad() []*ChatResponseBodyDataMessagesMetaDataMultiLoad {
	return s.MultiLoad
}

func (s *ChatResponseBodyDataMessagesMetaData) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *ChatResponseBodyDataMessagesMetaData) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyDataMessagesMetaData) SetBarTitle(v string) *ChatResponseBodyDataMessagesMetaData {
	s.BarTitle = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetContent(v interface{}) *ChatResponseBodyDataMessagesMetaData {
	s.Content = v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetMultiLoad(v []*ChatResponseBodyDataMessagesMetaDataMultiLoad) *ChatResponseBodyDataMessagesMetaData {
	s.MultiLoad = v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetSourceSeq(v string) *ChatResponseBodyDataMessagesMetaData {
	s.SourceSeq = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) SetType(v string) *ChatResponseBodyDataMessagesMetaData {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaData) Validate() error {
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

type ChatResponseBodyDataMessagesMetaDataMultiLoad struct {
	// The card content.
	//
	// example:
	//
	// {"list":[{"imageUrl":"https://images.example.com/cat.png"}]}
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

func (s ChatResponseBodyDataMessagesMetaDataMultiLoad) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyDataMessagesMetaDataMultiLoad) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) GetContent() interface{} {
	return s.Content
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) GetSourceSeq() *string {
	return s.SourceSeq
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) SetContent(v interface{}) *ChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Content = v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) SetSourceSeq(v string) *ChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.SourceSeq = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) SetType(v string) *ChatResponseBodyDataMessagesMetaDataMultiLoad {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyDataMessagesMetaDataMultiLoad) Validate() error {
	return dara.Validate(s)
}

type ChatResponseBodyError struct {
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
	// req-example-001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ChatResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyError) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyError) GetCode() interface{} {
	return s.Code
}

func (s *ChatResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *ChatResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBodyError) SetCode(v interface{}) *ChatResponseBodyError {
	s.Code = v
	return s
}

func (s *ChatResponseBodyError) SetMessage(v string) *ChatResponseBodyError {
	s.Message = &v
	return s
}

func (s *ChatResponseBodyError) SetRequestId(v string) *ChatResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBodyError) Validate() error {
	return dara.Validate(s)
}
