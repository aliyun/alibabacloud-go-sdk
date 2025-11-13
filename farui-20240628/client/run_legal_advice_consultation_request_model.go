// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLegalAdviceConsultationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunLegalAdviceConsultationRequest
	GetAppId() *string
	SetAssistant(v *RunLegalAdviceConsultationRequestAssistant) *RunLegalAdviceConsultationRequest
	GetAssistant() *RunLegalAdviceConsultationRequestAssistant
	SetExtra(v *RunLegalAdviceConsultationRequestExtra) *RunLegalAdviceConsultationRequest
	GetExtra() *RunLegalAdviceConsultationRequestExtra
	SetStream(v bool) *RunLegalAdviceConsultationRequest
	GetStream() *bool
	SetThread(v *RunLegalAdviceConsultationRequestThread) *RunLegalAdviceConsultationRequest
	GetThread() *RunLegalAdviceConsultationRequestThread
}

type RunLegalAdviceConsultationRequest struct {
	// example:
	//
	// farui
	AppId     *string                                     `json:"appId,omitempty" xml:"appId,omitempty"`
	Assistant *RunLegalAdviceConsultationRequestAssistant `json:"assistant,omitempty" xml:"assistant,omitempty" type:"Struct"`
	Extra     *RunLegalAdviceConsultationRequestExtra     `json:"extra,omitempty" xml:"extra,omitempty" type:"Struct"`
	// example:
	//
	// true
	Stream *bool                                    `json:"stream,omitempty" xml:"stream,omitempty"`
	Thread *RunLegalAdviceConsultationRequestThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s RunLegalAdviceConsultationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationRequest) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunLegalAdviceConsultationRequest) GetAssistant() *RunLegalAdviceConsultationRequestAssistant {
	return s.Assistant
}

func (s *RunLegalAdviceConsultationRequest) GetExtra() *RunLegalAdviceConsultationRequestExtra {
	return s.Extra
}

func (s *RunLegalAdviceConsultationRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunLegalAdviceConsultationRequest) GetThread() *RunLegalAdviceConsultationRequestThread {
	return s.Thread
}

func (s *RunLegalAdviceConsultationRequest) SetAppId(v string) *RunLegalAdviceConsultationRequest {
	s.AppId = &v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetAssistant(v *RunLegalAdviceConsultationRequestAssistant) *RunLegalAdviceConsultationRequest {
	s.Assistant = v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetExtra(v *RunLegalAdviceConsultationRequestExtra) *RunLegalAdviceConsultationRequest {
	s.Extra = v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetStream(v bool) *RunLegalAdviceConsultationRequest {
	s.Stream = &v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetThread(v *RunLegalAdviceConsultationRequestThread) *RunLegalAdviceConsultationRequest {
	s.Thread = v
	return s
}

func (s *RunLegalAdviceConsultationRequest) Validate() error {
	if s.Assistant != nil {
		if err := s.Assistant.Validate(); err != nil {
			return err
		}
	}
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	if s.Thread != nil {
		if err := s.Thread.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunLegalAdviceConsultationRequestAssistant struct {
	// example:
	//
	// assitant_abc_123
	Id       *string            `json:"id,omitempty" xml:"id,omitempty"`
	MetaData map[string]*string `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// example:
	//
	// legal_advice_consult
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RunLegalAdviceConsultationRequestAssistant) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestAssistant) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestAssistant) GetId() *string {
	return s.Id
}

func (s *RunLegalAdviceConsultationRequestAssistant) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *RunLegalAdviceConsultationRequestAssistant) GetType() *string {
	return s.Type
}

func (s *RunLegalAdviceConsultationRequestAssistant) GetVersion() *string {
	return s.Version
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetId(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Id = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetMetaData(v map[string]*string) *RunLegalAdviceConsultationRequestAssistant {
	s.MetaData = v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetType(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Type = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetVersion(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Version = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) Validate() error {
	return dara.Validate(s)
}

type RunLegalAdviceConsultationRequestExtra struct {
	DeepThink    *bool `json:"deepThink,omitempty" xml:"deepThink,omitempty"`
	OnlineSearch *bool `json:"onlineSearch,omitempty" xml:"onlineSearch,omitempty"`
}

func (s RunLegalAdviceConsultationRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestExtra) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestExtra) GetDeepThink() *bool {
	return s.DeepThink
}

func (s *RunLegalAdviceConsultationRequestExtra) GetOnlineSearch() *bool {
	return s.OnlineSearch
}

func (s *RunLegalAdviceConsultationRequestExtra) SetDeepThink(v bool) *RunLegalAdviceConsultationRequestExtra {
	s.DeepThink = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestExtra) SetOnlineSearch(v bool) *RunLegalAdviceConsultationRequestExtra {
	s.OnlineSearch = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestExtra) Validate() error {
	return dara.Validate(s)
}

type RunLegalAdviceConsultationRequestThread struct {
	Messages []*RunLegalAdviceConsultationRequestThreadMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s RunLegalAdviceConsultationRequestThread) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestThread) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestThread) GetMessages() []*RunLegalAdviceConsultationRequestThreadMessages {
	return s.Messages
}

func (s *RunLegalAdviceConsultationRequestThread) SetMessages(v []*RunLegalAdviceConsultationRequestThreadMessages) *RunLegalAdviceConsultationRequestThread {
	s.Messages = v
	return s
}

func (s *RunLegalAdviceConsultationRequestThread) Validate() error {
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

type RunLegalAdviceConsultationRequestThreadMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunLegalAdviceConsultationRequestThreadMessages) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestThreadMessages) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) GetContent() *string {
	return s.Content
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) GetRole() *string {
	return s.Role
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) SetContent(v string) *RunLegalAdviceConsultationRequestThreadMessages {
	s.Content = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) SetRole(v string) *RunLegalAdviceConsultationRequestThreadMessages {
	s.Role = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) Validate() error {
	return dara.Validate(s)
}
