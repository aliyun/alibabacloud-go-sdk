// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMessageResponseBodyData) *GetMessageResponseBody
	GetData() *GetMessageResponseBodyData
	SetRequestId(v string) *GetMessageResponseBody
	GetRequestId() *string
}

type GetMessageResponseBody struct {
	// The data returned.
	Data *GetMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// 8724BC18-904D-5A0D-BFF4-F0554F0037E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageResponseBody) GetData() *GetMessageResponseBodyData {
	return s.Data
}

func (s *GetMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageResponseBody) SetData(v *GetMessageResponseBodyData) *GetMessageResponseBody {
	s.Data = v
	return s
}

func (s *GetMessageResponseBody) SetRequestId(v string) *GetMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMessageResponseBodyData struct {
	// Message content.
	//
	// example:
	//
	// {********}
	MessageContent *string `json:"MessageContent,omitempty" xml:"MessageContent,omitempty"`
	// Message ID.
	//
	// example:
	//
	// m-123***7890
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Message level.
	//
	// example:
	//
	// WARNING
	MessageLevel *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
	// Message name.
	//
	// example:
	//
	// MyMessage
	MessageName *string `json:"MessageName,omitempty" xml:"MessageName,omitempty"`
	// Message source ID.
	//
	// example:
	//
	// p-123***7890
	MessageSourceId *string `json:"MessageSourceId,omitempty" xml:"MessageSourceId,omitempty"`
	// Message source region ID.
	//
	// example:
	//
	// cn-hangzhou
	MessageSourceRegionId *string `json:"MessageSourceRegionId,omitempty" xml:"MessageSourceRegionId,omitempty"`
	// Message source type.
	//
	// example:
	//
	// PROTECTION_POLICY
	MessageSourceType *string `json:"MessageSourceType,omitempty" xml:"MessageSourceType,omitempty"`
	// Message time.
	//
	// example:
	//
	// 1740019609
	MessageTime *int64 `json:"MessageTime,omitempty" xml:"MessageTime,omitempty"`
	// Message type.
	//
	// example:
	//
	// SUB_PROTECTION_POLICY_MODIFIED
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMessageResponseBodyData) GetMessageContent() *string {
	return s.MessageContent
}

func (s *GetMessageResponseBodyData) GetMessageId() *string {
	return s.MessageId
}

func (s *GetMessageResponseBodyData) GetMessageLevel() *string {
	return s.MessageLevel
}

func (s *GetMessageResponseBodyData) GetMessageName() *string {
	return s.MessageName
}

func (s *GetMessageResponseBodyData) GetMessageSourceId() *string {
	return s.MessageSourceId
}

func (s *GetMessageResponseBodyData) GetMessageSourceRegionId() *string {
	return s.MessageSourceRegionId
}

func (s *GetMessageResponseBodyData) GetMessageSourceType() *string {
	return s.MessageSourceType
}

func (s *GetMessageResponseBodyData) GetMessageTime() *int64 {
	return s.MessageTime
}

func (s *GetMessageResponseBodyData) GetMessageType() *string {
	return s.MessageType
}

func (s *GetMessageResponseBodyData) SetMessageContent(v string) *GetMessageResponseBodyData {
	s.MessageContent = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageId(v string) *GetMessageResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageLevel(v string) *GetMessageResponseBodyData {
	s.MessageLevel = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageName(v string) *GetMessageResponseBodyData {
	s.MessageName = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageSourceId(v string) *GetMessageResponseBodyData {
	s.MessageSourceId = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageSourceRegionId(v string) *GetMessageResponseBodyData {
	s.MessageSourceRegionId = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageSourceType(v string) *GetMessageResponseBodyData {
	s.MessageSourceType = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageTime(v int64) *GetMessageResponseBodyData {
	s.MessageTime = &v
	return s
}

func (s *GetMessageResponseBodyData) SetMessageType(v string) *GetMessageResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *GetMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
