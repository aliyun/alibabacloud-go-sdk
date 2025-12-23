// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSavedQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSavedQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSavedQueryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSavedQueryResponseBody) *UpdateSavedQueryResponse
	GetBody() *UpdateSavedQueryResponseBody
}

type UpdateSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSavedQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSavedQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSavedQueryResponse) GetBody() *UpdateSavedQueryResponseBody {
	return s.Body
}

func (s *UpdateSavedQueryResponse) SetHeaders(v map[string]*string) *UpdateSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *UpdateSavedQueryResponse) SetStatusCode(v int32) *UpdateSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSavedQueryResponse) SetBody(v *UpdateSavedQueryResponseBody) *UpdateSavedQueryResponse {
	s.Body = v
	return s
}

func (s *UpdateSavedQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
