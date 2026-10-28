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
	// Optional. The additional information, which is a JSON string.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	ImageKeys *string `json:"ImageKeys,omitempty" xml:"ImageKeys,omitempty"`
	// The natural language description of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// Are there any issues or abnormalities with my instance rm-xxx?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Optional. The session ID, which is a UUID string. If this parameter is not specified, a new session is created. To maintain context in a conversation, use the same session ID.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-xxxxxxxxxxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Optional. The source cloud service. Specifies the cloud service to which the Q&A belongs. Default value: yaochi.
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
