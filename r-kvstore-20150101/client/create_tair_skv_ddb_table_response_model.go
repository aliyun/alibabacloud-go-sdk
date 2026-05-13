// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairSkvDdbTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTairSkvDdbTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTairSkvDdbTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateTairSkvDdbTableResponseBody) *CreateTairSkvDdbTableResponse
	GetBody() *CreateTairSkvDdbTableResponseBody
}

type CreateTairSkvDdbTableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTairSkvDdbTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTairSkvDdbTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTairSkvDdbTableResponse) GoString() string {
	return s.String()
}

func (s *CreateTairSkvDdbTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTairSkvDdbTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTairSkvDdbTableResponse) GetBody() *CreateTairSkvDdbTableResponseBody {
	return s.Body
}

func (s *CreateTairSkvDdbTableResponse) SetHeaders(v map[string]*string) *CreateTairSkvDdbTableResponse {
	s.Headers = v
	return s
}

func (s *CreateTairSkvDdbTableResponse) SetStatusCode(v int32) *CreateTairSkvDdbTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTairSkvDdbTableResponse) SetBody(v *CreateTairSkvDdbTableResponseBody) *CreateTairSkvDdbTableResponse {
	s.Body = v
	return s
}

func (s *CreateTairSkvDdbTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
