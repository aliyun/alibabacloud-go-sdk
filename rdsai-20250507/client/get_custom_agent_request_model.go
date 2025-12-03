// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *GetCustomAgentRequest
	GetApiId() *string
	SetCustomAgentId(v string) *GetCustomAgentRequest
	GetCustomAgentId() *string
}

type GetCustomAgentRequest struct {
	// example:
	//
	// app-iBuGU1VxEY42zrQRQfNAn3oj
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// ebe44453-3b41-4c74-94d1-01d088d7xxxx
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
}

func (s GetCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *GetCustomAgentRequest) GetApiId() *string {
	return s.ApiId
}

func (s *GetCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *GetCustomAgentRequest) SetApiId(v string) *GetCustomAgentRequest {
	s.ApiId = &v
	return s
}

func (s *GetCustomAgentRequest) SetCustomAgentId(v string) *GetCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *GetCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
