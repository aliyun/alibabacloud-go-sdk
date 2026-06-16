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
	// Application ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2052929167853146113
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// User message object containing role and multimodal content.
	//
	// This parameter is required.
	Message *QaChatRequestMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
	// No input required
	//
	// example:
	//
	// {
	//
	//   "debug": true
	//
	// }
	Options map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
	// Q&A session ID, used to track multiple Q&A interactions from the same user.
	//
	// example:
	//
	// req_123456789
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
	// Individual content block, differentiated by `type`
	Parts []*QaChatRequestMessageParts `json:"parts,omitempty" xml:"parts,omitempty" type:"Repeated"`
	// Message role, currently only supports the `"user"` role
	//
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
	// Required when type = "data". The data object structure is as follows:
	//
	// - type: String type, required, indicates the data subtype. Currently supported value is "template", indicating a video template.
	//
	// - videoId: String type, conditionally required. Only required when type = "template", indicating the video template ID; can be ignored or set to null for other types.
	//
	// example:
	//
	// {
	//
	//   "type": "template",
	//
	//   "videoId": "xxxx"
	//
	// }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Required when `type="file"`.
	//
	// 	- Media type, currently only supports image formats JPG/PNG/WEBP/JPEG, maximum 5
	//
	// example:
	//
	// image/png
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	// Required when `type="text"`.
	//
	// 	- Text content, maximum 1024 characters
	//
	// example:
	//
	// 请问这个视频讲了什么？
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// Fixed content block type, only supports `"text"` / `"file"` / `"data"`
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Required when `type="file"`. Supports the following two types, with format support for JPG/PNG/WEBP/JPEG:
	//
	// • Media resource CDN URL, currently supports images, maximum 5;
	//
	// • Image encoding, upload image files using base64 encoded strings (supports bitmap formats), maximum 5
	//
	// example:
	//
	// https://example.com/img.jpg
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
