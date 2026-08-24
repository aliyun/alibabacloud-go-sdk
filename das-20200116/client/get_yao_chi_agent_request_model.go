// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYaoChiAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtraInfo(v string) *GetYaoChiAgentRequest
	GetExtraInfo() *string
	SetImageKeys(v string) *GetYaoChiAgentRequest
	GetImageKeys() *string
	SetQuery(v string) *GetYaoChiAgentRequest
	GetQuery() *string
	SetSessionId(v string) *GetYaoChiAgentRequest
	GetSessionId() *string
	SetSource(v string) *GetYaoChiAgentRequest
	GetSource() *string
}

type GetYaoChiAgentRequest struct {
	// The additional information in JSON string format. This parameter is optional.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	ImageKeys *string `json:"ImageKeys,omitempty" xml:"ImageKeys,omitempty"`
	// The natural language description of the question.
	//
	// This parameter is required.
	//
	// example:
	//
	// Are there any issues or abnormalities with my instance rm-xxx?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The session ID in UUID string format. This parameter is optional. If you do not specify this parameter, a new session is created. To maintain context across a conversation, use the same session ID.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-xxxxxxxxxxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The cloud service source that specifies which cloud service the question belongs to. This parameter is optional. Default value: yaochi.
	//
	// example:
	//
	// yaochi
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetYaoChiAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYaoChiAgentRequest) GoString() string {
	return s.String()
}

func (s *GetYaoChiAgentRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetYaoChiAgentRequest) GetImageKeys() *string {
	return s.ImageKeys
}

func (s *GetYaoChiAgentRequest) GetQuery() *string {
	return s.Query
}

func (s *GetYaoChiAgentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetYaoChiAgentRequest) GetSource() *string {
	return s.Source
}

func (s *GetYaoChiAgentRequest) SetExtraInfo(v string) *GetYaoChiAgentRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GetYaoChiAgentRequest) SetImageKeys(v string) *GetYaoChiAgentRequest {
	s.ImageKeys = &v
	return s
}

func (s *GetYaoChiAgentRequest) SetQuery(v string) *GetYaoChiAgentRequest {
	s.Query = &v
	return s
}

func (s *GetYaoChiAgentRequest) SetSessionId(v string) *GetYaoChiAgentRequest {
	s.SessionId = &v
	return s
}

func (s *GetYaoChiAgentRequest) SetSource(v string) *GetYaoChiAgentRequest {
	s.Source = &v
	return s
}

func (s *GetYaoChiAgentRequest) Validate() error {
	return dara.Validate(s)
}
