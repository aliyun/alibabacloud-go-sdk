// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordsToLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddKeywordsToLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddKeywordsToLibResponse
	GetStatusCode() *int32
	SetBody(v *AddKeywordsToLibResponseBody) *AddKeywordsToLibResponse
	GetBody() *AddKeywordsToLibResponseBody
}

type AddKeywordsToLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKeywordsToLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKeywordsToLibResponse) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsToLibResponse) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddKeywordsToLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddKeywordsToLibResponse) GetBody() *AddKeywordsToLibResponseBody {
	return s.Body
}

func (s *AddKeywordsToLibResponse) SetHeaders(v map[string]*string) *AddKeywordsToLibResponse {
	s.Headers = v
	return s
}

func (s *AddKeywordsToLibResponse) SetStatusCode(v int32) *AddKeywordsToLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKeywordsToLibResponse) SetBody(v *AddKeywordsToLibResponseBody) *AddKeywordsToLibResponse {
	s.Body = v
	return s
}

func (s *AddKeywordsToLibResponse) Validate() error {
	return dara.Validate(s)
}
