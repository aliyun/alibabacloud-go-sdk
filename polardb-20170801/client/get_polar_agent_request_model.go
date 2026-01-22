// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtraInfo(v string) *GetPolarAgentRequest
	GetExtraInfo() *string
	SetQuery(v string) *GetPolarAgentRequest
	GetQuery() *string
	SetSessionId(v string) *GetPolarAgentRequest
	GetSessionId() *string
	SetSource(v string) *GetPolarAgentRequest
	GetSource() *string
}

type GetPolarAgentRequest struct {
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c4d242f3-c909-4846-91d9-f84c238a9820
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// c4d242f3-c909-4846-91d9-f84c238a9820
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// polardb-console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetPolarAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolarAgentRequest) GoString() string {
	return s.String()
}

func (s *GetPolarAgentRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetPolarAgentRequest) GetQuery() *string {
	return s.Query
}

func (s *GetPolarAgentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetPolarAgentRequest) GetSource() *string {
	return s.Source
}

func (s *GetPolarAgentRequest) SetExtraInfo(v string) *GetPolarAgentRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GetPolarAgentRequest) SetQuery(v string) *GetPolarAgentRequest {
	s.Query = &v
	return s
}

func (s *GetPolarAgentRequest) SetSessionId(v string) *GetPolarAgentRequest {
	s.SessionId = &v
	return s
}

func (s *GetPolarAgentRequest) SetSource(v string) *GetPolarAgentRequest {
	s.Source = &v
	return s
}

func (s *GetPolarAgentRequest) Validate() error {
	return dara.Validate(s)
}
