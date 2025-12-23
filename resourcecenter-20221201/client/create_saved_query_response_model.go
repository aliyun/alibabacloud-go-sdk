// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavedQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSavedQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSavedQueryResponse
	GetStatusCode() *int32
	SetBody(v *CreateSavedQueryResponseBody) *CreateSavedQueryResponse
	GetBody() *CreateSavedQueryResponseBody
}

type CreateSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSavedQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSavedQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSavedQueryResponse) GetBody() *CreateSavedQueryResponseBody {
	return s.Body
}

func (s *CreateSavedQueryResponse) SetHeaders(v map[string]*string) *CreateSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *CreateSavedQueryResponse) SetStatusCode(v int32) *CreateSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSavedQueryResponse) SetBody(v *CreateSavedQueryResponseBody) *CreateSavedQueryResponse {
	s.Body = v
	return s
}

func (s *CreateSavedQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
