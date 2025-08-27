// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCorpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubCorpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubCorpResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubCorpResponseBody) *CreateSubCorpResponse
	GetBody() *CreateSubCorpResponseBody
}

type CreateSubCorpResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubCorpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubCorpResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCorpResponse) GoString() string {
	return s.String()
}

func (s *CreateSubCorpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubCorpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubCorpResponse) GetBody() *CreateSubCorpResponseBody {
	return s.Body
}

func (s *CreateSubCorpResponse) SetHeaders(v map[string]*string) *CreateSubCorpResponse {
	s.Headers = v
	return s
}

func (s *CreateSubCorpResponse) SetStatusCode(v int32) *CreateSubCorpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubCorpResponse) SetBody(v *CreateSubCorpResponseBody) *CreateSubCorpResponse {
	s.Body = v
	return s
}

func (s *CreateSubCorpResponse) Validate() error {
	return dara.Validate(s)
}
