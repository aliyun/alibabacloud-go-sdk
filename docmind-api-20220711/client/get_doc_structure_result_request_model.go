// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocStructureResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDocStructureResultRequest
	GetId() *string
	SetImageStrategy(v string) *GetDocStructureResultRequest
	GetImageStrategy() *string
	SetRevealMarkdown(v bool) *GetDocStructureResultRequest
	GetRevealMarkdown() *bool
	SetUseUrlResponseBody(v bool) *GetDocStructureResultRequest
	GetUseUrlResponseBody() *bool
}

type GetDocStructureResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageStrategy      *string `json:"ImageStrategy,omitempty" xml:"ImageStrategy,omitempty"`
	RevealMarkdown     *bool   `json:"RevealMarkdown,omitempty" xml:"RevealMarkdown,omitempty"`
	UseUrlResponseBody *bool   `json:"UseUrlResponseBody,omitempty" xml:"UseUrlResponseBody,omitempty"`
}

func (s GetDocStructureResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocStructureResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocStructureResultRequest) GetId() *string {
	return s.Id
}

func (s *GetDocStructureResultRequest) GetImageStrategy() *string {
	return s.ImageStrategy
}

func (s *GetDocStructureResultRequest) GetRevealMarkdown() *bool {
	return s.RevealMarkdown
}

func (s *GetDocStructureResultRequest) GetUseUrlResponseBody() *bool {
	return s.UseUrlResponseBody
}

func (s *GetDocStructureResultRequest) SetId(v string) *GetDocStructureResultRequest {
	s.Id = &v
	return s
}

func (s *GetDocStructureResultRequest) SetImageStrategy(v string) *GetDocStructureResultRequest {
	s.ImageStrategy = &v
	return s
}

func (s *GetDocStructureResultRequest) SetRevealMarkdown(v bool) *GetDocStructureResultRequest {
	s.RevealMarkdown = &v
	return s
}

func (s *GetDocStructureResultRequest) SetUseUrlResponseBody(v bool) *GetDocStructureResultRequest {
	s.UseUrlResponseBody = &v
	return s
}

func (s *GetDocStructureResultRequest) Validate() error {
	return dara.Validate(s)
}
