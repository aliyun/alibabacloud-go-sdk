// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQaChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QaChatRequest
	GetAppId() *string
	SetMessage(v *QaChatRequestMessage) *QaChatRequest
	GetMessage() *QaChatRequestMessage
	SetOptions(v map[string]interface{}) *QaChatRequest
	GetOptions() map[string]interface{}
	SetSessionId(v string) *QaChatRequest
	GetSessionId() *string
}

type QaChatRequest struct {
	// example:
	//
	// 2047140750220754946
	AppId   *string               `json:"appId,omitempty" xml:"appId,omitempty"`
	Message *QaChatRequestMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
	// example:
	//
	// {
	//
	//   "debug": true
	//
	// }
	Options map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
	// example:
	//
	// b2a979e79799489fbde56119bf8c4dc7
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s QaChatRequest) String() string {
	return dara.Prettify(s)
}

func (s QaChatRequest) GoString() string {
	return s.String()
}

func (s *QaChatRequest) GetAppId() *string {
	return s.AppId
}

func (s *QaChatRequest) GetMessage() *QaChatRequestMessage {
	return s.Message
}

func (s *QaChatRequest) GetOptions() map[string]interface{} {
	return s.Options
}

func (s *QaChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *QaChatRequest) SetAppId(v string) *QaChatRequest {
	s.AppId = &v
	return s
}

func (s *QaChatRequest) SetMessage(v *QaChatRequestMessage) *QaChatRequest {
	s.Message = v
	return s
}

func (s *QaChatRequest) SetOptions(v map[string]interface{}) *QaChatRequest {
	s.Options = v
	return s
}

func (s *QaChatRequest) SetSessionId(v string) *QaChatRequest {
	s.SessionId = &v
	return s
}

func (s *QaChatRequest) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QaChatRequestMessage struct {
	Parts []*QaChatRequestMessageParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s QaChatRequestMessage) String() string {
	return dara.Prettify(s)
}

func (s QaChatRequestMessage) GoString() string {
	return s.String()
}

func (s *QaChatRequestMessage) GetParts() []*QaChatRequestMessageParts {
	return s.Parts
}

func (s *QaChatRequestMessage) GetRole() *string {
	return s.Role
}

func (s *QaChatRequestMessage) SetParts(v []*QaChatRequestMessageParts) *QaChatRequestMessage {
	s.Parts = v
	return s
}

func (s *QaChatRequestMessage) SetRole(v string) *QaChatRequestMessage {
	s.Role = &v
	return s
}

func (s *QaChatRequestMessage) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QaChatRequestMessageParts struct {
	// example:
	//
	// {
	//
	//   "templateId": "456789"
	//
	// }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// image/png
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	// example:
	//
	// 帮我搜索下今天的天气
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// https://meeting.dingtalk.com/j/4sSPAxWaPbM
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QaChatRequestMessageParts) String() string {
	return dara.Prettify(s)
}

func (s QaChatRequestMessageParts) GoString() string {
	return s.String()
}

func (s *QaChatRequestMessageParts) GetData() interface{} {
	return s.Data
}

func (s *QaChatRequestMessageParts) GetMediaType() *string {
	return s.MediaType
}

func (s *QaChatRequestMessageParts) GetText() *string {
	return s.Text
}

func (s *QaChatRequestMessageParts) GetType() *string {
	return s.Type
}

func (s *QaChatRequestMessageParts) GetUrl() *string {
	return s.Url
}

func (s *QaChatRequestMessageParts) SetData(v interface{}) *QaChatRequestMessageParts {
	s.Data = v
	return s
}

func (s *QaChatRequestMessageParts) SetMediaType(v string) *QaChatRequestMessageParts {
	s.MediaType = &v
	return s
}

func (s *QaChatRequestMessageParts) SetText(v string) *QaChatRequestMessageParts {
	s.Text = &v
	return s
}

func (s *QaChatRequestMessageParts) SetType(v string) *QaChatRequestMessageParts {
	s.Type = &v
	return s
}

func (s *QaChatRequestMessageParts) SetUrl(v string) *QaChatRequestMessageParts {
	s.Url = &v
	return s
}

func (s *QaChatRequestMessageParts) Validate() error {
	return dara.Validate(s)
}
