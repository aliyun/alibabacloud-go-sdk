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
	SetQuery(v string) *GetYaoChiAgentRequest
	GetQuery() *string
	SetSessionId(v string) *GetYaoChiAgentRequest
	GetSessionId() *string
	SetSource(v string) *GetYaoChiAgentRequest
	GetSource() *string
}

type GetYaoChiAgentRequest struct {
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Are there any issues or abnormalities with my instance rm-xxx?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-xxxxxxxxxxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
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
