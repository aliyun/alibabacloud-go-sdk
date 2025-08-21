// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuggestShrinkableNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSuggestShrinkableNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSuggestShrinkableNodesResponse
	GetStatusCode() *int32
	SetBody(v *GetSuggestShrinkableNodesResponseBody) *GetSuggestShrinkableNodesResponse
	GetBody() *GetSuggestShrinkableNodesResponseBody
}

type GetSuggestShrinkableNodesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuggestShrinkableNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuggestShrinkableNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSuggestShrinkableNodesResponse) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSuggestShrinkableNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSuggestShrinkableNodesResponse) GetBody() *GetSuggestShrinkableNodesResponseBody {
	return s.Body
}

func (s *GetSuggestShrinkableNodesResponse) SetHeaders(v map[string]*string) *GetSuggestShrinkableNodesResponse {
	s.Headers = v
	return s
}

func (s *GetSuggestShrinkableNodesResponse) SetStatusCode(v int32) *GetSuggestShrinkableNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuggestShrinkableNodesResponse) SetBody(v *GetSuggestShrinkableNodesResponseBody) *GetSuggestShrinkableNodesResponse {
	s.Body = v
	return s
}

func (s *GetSuggestShrinkableNodesResponse) Validate() error {
	return dara.Validate(s)
}
