// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiproEmbedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryAiproEmbedConfigRequest
	GetKeyword() *string
}

type QueryAiproEmbedConfigRequest struct {
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s QueryAiproEmbedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiproEmbedConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryAiproEmbedConfigRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryAiproEmbedConfigRequest) SetKeyword(v string) *QueryAiproEmbedConfigRequest {
	s.Keyword = &v
	return s
}

func (s *QueryAiproEmbedConfigRequest) Validate() error {
	return dara.Validate(s)
}
