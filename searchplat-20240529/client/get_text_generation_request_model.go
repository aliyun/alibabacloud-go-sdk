// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsiLevel(v string) *GetTextGenerationRequest
	GetCsiLevel() *string
	SetEnableSearch(v bool) *GetTextGenerationRequest
	GetEnableSearch() *bool
	SetMessages(v []*GetTextGenerationRequestMessages) *GetTextGenerationRequest
	GetMessages() []*GetTextGenerationRequestMessages
	SetParameters(v map[string]interface{}) *GetTextGenerationRequest
	GetParameters() map[string]interface{}
	SetStream(v bool) *GetTextGenerationRequest
	GetStream() *bool
}

type GetTextGenerationRequest struct {
	CsiLevel     *string `json:"csi_level,omitempty" xml:"csi_level,omitempty"`
	EnableSearch *bool   `json:"enable_search,omitempty" xml:"enable_search,omitempty"`
	// This parameter is required.
	Messages   []*GetTextGenerationRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	Parameters map[string]interface{}              `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Stream     *bool                               `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s GetTextGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTextGenerationRequest) GoString() string {
	return s.String()
}

func (s *GetTextGenerationRequest) GetCsiLevel() *string {
	return s.CsiLevel
}

func (s *GetTextGenerationRequest) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *GetTextGenerationRequest) GetMessages() []*GetTextGenerationRequestMessages {
	return s.Messages
}

func (s *GetTextGenerationRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetTextGenerationRequest) GetStream() *bool {
	return s.Stream
}

func (s *GetTextGenerationRequest) SetCsiLevel(v string) *GetTextGenerationRequest {
	s.CsiLevel = &v
	return s
}

func (s *GetTextGenerationRequest) SetEnableSearch(v bool) *GetTextGenerationRequest {
	s.EnableSearch = &v
	return s
}

func (s *GetTextGenerationRequest) SetMessages(v []*GetTextGenerationRequestMessages) *GetTextGenerationRequest {
	s.Messages = v
	return s
}

func (s *GetTextGenerationRequest) SetParameters(v map[string]interface{}) *GetTextGenerationRequest {
	s.Parameters = v
	return s
}

func (s *GetTextGenerationRequest) SetStream(v bool) *GetTextGenerationRequest {
	s.Stream = &v
	return s
}

func (s *GetTextGenerationRequest) Validate() error {
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

type GetTextGenerationRequestMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetTextGenerationRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s GetTextGenerationRequestMessages) GoString() string {
	return s.String()
}

func (s *GetTextGenerationRequestMessages) GetContent() *string {
	return s.Content
}

func (s *GetTextGenerationRequestMessages) GetRole() *string {
	return s.Role
}

func (s *GetTextGenerationRequestMessages) SetContent(v string) *GetTextGenerationRequestMessages {
	s.Content = &v
	return s
}

func (s *GetTextGenerationRequestMessages) SetRole(v string) *GetTextGenerationRequestMessages {
	s.Role = &v
	return s
}

func (s *GetTextGenerationRequestMessages) Validate() error {
	return dara.Validate(s)
}
