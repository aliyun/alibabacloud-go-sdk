// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOmniSearchQuery interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *OmniSearchQuery
	GetQuery() *string
	SetSessionId(v string) *OmniSearchQuery
	GetSessionId() *string
	SetUseCot(v bool) *OmniSearchQuery
	GetUseCot() *bool
	SetUseModel(v string) *OmniSearchQuery
	GetUseModel() *string
}

type OmniSearchQuery struct {
	// The request parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// What is the weather in Hangzhou today?
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The session ID. This is a reserved field and is not currently supported.
	//
	// example:
	//
	// session_id_123
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to use CoT. This is a reserved field and is not currently supported.
	//
	// example:
	//
	// false
	UseCot *bool `json:"useCot,omitempty" xml:"useCot,omitempty"`
	// The model version to use. Valid values: pro for the deep version, and lite for the fast version.
	//
	// example:
	//
	// pro
	UseModel *string `json:"useModel,omitempty" xml:"useModel,omitempty"`
}

func (s OmniSearchQuery) String() string {
	return dara.Prettify(s)
}

func (s OmniSearchQuery) GoString() string {
	return s.String()
}

func (s *OmniSearchQuery) GetQuery() *string {
	return s.Query
}

func (s *OmniSearchQuery) GetSessionId() *string {
	return s.SessionId
}

func (s *OmniSearchQuery) GetUseCot() *bool {
	return s.UseCot
}

func (s *OmniSearchQuery) GetUseModel() *string {
	return s.UseModel
}

func (s *OmniSearchQuery) SetQuery(v string) *OmniSearchQuery {
	s.Query = &v
	return s
}

func (s *OmniSearchQuery) SetSessionId(v string) *OmniSearchQuery {
	s.SessionId = &v
	return s
}

func (s *OmniSearchQuery) SetUseCot(v bool) *OmniSearchQuery {
	s.UseCot = &v
	return s
}

func (s *OmniSearchQuery) SetUseModel(v string) *OmniSearchQuery {
	s.UseModel = &v
	return s
}

func (s *OmniSearchQuery) Validate() error {
	return dara.Validate(s)
}
