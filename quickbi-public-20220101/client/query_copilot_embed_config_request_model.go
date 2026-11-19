// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopilotEmbedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryCopilotEmbedConfigRequest
	GetKeyword() *string
}

type QueryCopilotEmbedConfigRequest struct {
	// The name of the embedding configuration module. Fuzzy search is supported.
	//
	// example:
	//
	// 06-ELive
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s QueryCopilotEmbedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCopilotEmbedConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryCopilotEmbedConfigRequest) SetKeyword(v string) *QueryCopilotEmbedConfigRequest {
	s.Keyword = &v
	return s
}

func (s *QueryCopilotEmbedConfigRequest) Validate() error {
	return dara.Validate(s)
}
