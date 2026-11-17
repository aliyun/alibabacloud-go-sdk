// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertDirectConversationToWorkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConvertDirectConversationToWorkResponseBody
	GetCode() *string
	SetConversionId(v string) *ConvertDirectConversationToWorkResponseBody
	GetConversionId() *string
	SetConvertedRounds(v int32) *ConvertDirectConversationToWorkResponseBody
	GetConvertedRounds() *int32
	SetHistoryScope(v string) *ConvertDirectConversationToWorkResponseBody
	GetHistoryScope() *string
	SetMessage(v string) *ConvertDirectConversationToWorkResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConvertDirectConversationToWorkResponseBody
	GetRequestId() *string
	SetRetryable(v bool) *ConvertDirectConversationToWorkResponseBody
	GetRetryable() *bool
	SetSourceEpoch(v int64) *ConvertDirectConversationToWorkResponseBody
	GetSourceEpoch() *int64
	SetSourceSessionId(v string) *ConvertDirectConversationToWorkResponseBody
	GetSourceSessionId() *string
	SetStatus(v string) *ConvertDirectConversationToWorkResponseBody
	GetStatus() *string
	SetTargetSessionId(v string) *ConvertDirectConversationToWorkResponseBody
	GetTargetSessionId() *string
	SetThroughSequence(v int64) *ConvertDirectConversationToWorkResponseBody
	GetThroughSequence() *int64
	SetWorkMode(v string) *ConvertDirectConversationToWorkResponseBody
	GetWorkMode() *string
}

type ConvertDirectConversationToWorkResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The session conversion record ID.
	//
	// example:
	//
	// exampleConversionId
	ConversionId *string `json:"conversionId,omitempty" xml:"conversionId,omitempty"`
	// The number of converted conversation turns.
	//
	// example:
	//
	// 2
	ConvertedRounds *int32 `json:"convertedRounds,omitempty" xml:"convertedRounds,omitempty"`
	// The scope of the converted historical messages.
	//
	// example:
	//
	// direct
	HistoryScope *string `json:"historyScope,omitempty" xml:"historyScope,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Specifies whether retries are supported.
	//
	// example:
	//
	// false
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
	// The source session version identifier.
	//
	// example:
	//
	// 0
	SourceEpoch *int64 `json:"sourceEpoch,omitempty" xml:"sourceEpoch,omitempty"`
	// The source session ID.
	//
	// example:
	//
	// exampleSourceSessionId
	SourceSessionId *string `json:"sourceSessionId,omitempty" xml:"sourceSessionId,omitempty"`
	// The conversion status.
	//
	// example:
	//
	// succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The target work session ID.
	//
	// example:
	//
	// exampleTargetSessionId
	TargetSessionId *string `json:"targetSessionId,omitempty" xml:"targetSessionId,omitempty"`
	// The cutoff message sequence number for the conversion.
	//
	// example:
	//
	// 2
	ThroughSequence *int64 `json:"throughSequence,omitempty" xml:"throughSequence,omitempty"`
	// The work session mode.
	//
	// example:
	//
	// work
	WorkMode *string `json:"workMode,omitempty" xml:"workMode,omitempty"`
}

func (s ConvertDirectConversationToWorkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertDirectConversationToWorkResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertDirectConversationToWorkResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConvertDirectConversationToWorkResponseBody) GetConversionId() *string {
	return s.ConversionId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetConvertedRounds() *int32 {
	return s.ConvertedRounds
}

func (s *ConvertDirectConversationToWorkResponseBody) GetHistoryScope() *string {
	return s.HistoryScope
}

func (s *ConvertDirectConversationToWorkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConvertDirectConversationToWorkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetRetryable() *bool {
	return s.Retryable
}

func (s *ConvertDirectConversationToWorkResponseBody) GetSourceEpoch() *int64 {
	return s.SourceEpoch
}

func (s *ConvertDirectConversationToWorkResponseBody) GetSourceSessionId() *string {
	return s.SourceSessionId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ConvertDirectConversationToWorkResponseBody) GetTargetSessionId() *string {
	return s.TargetSessionId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetThroughSequence() *int64 {
	return s.ThroughSequence
}

func (s *ConvertDirectConversationToWorkResponseBody) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ConvertDirectConversationToWorkResponseBody) SetCode(v string) *ConvertDirectConversationToWorkResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetConversionId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.ConversionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetConvertedRounds(v int32) *ConvertDirectConversationToWorkResponseBody {
	s.ConvertedRounds = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetHistoryScope(v string) *ConvertDirectConversationToWorkResponseBody {
	s.HistoryScope = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetMessage(v string) *ConvertDirectConversationToWorkResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetRequestId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetRetryable(v bool) *ConvertDirectConversationToWorkResponseBody {
	s.Retryable = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetSourceEpoch(v int64) *ConvertDirectConversationToWorkResponseBody {
	s.SourceEpoch = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetSourceSessionId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.SourceSessionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetStatus(v string) *ConvertDirectConversationToWorkResponseBody {
	s.Status = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetTargetSessionId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.TargetSessionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetThroughSequence(v int64) *ConvertDirectConversationToWorkResponseBody {
	s.ThroughSequence = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetWorkMode(v string) *ConvertDirectConversationToWorkResponseBody {
	s.WorkMode = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) Validate() error {
	return dara.Validate(s)
}
