// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSavedQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSavedQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSavedQueryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSavedQueryResponseBody) *DeleteSavedQueryResponse
	GetBody() *DeleteSavedQueryResponseBody
}

type DeleteSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSavedQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSavedQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSavedQueryResponse) GetBody() *DeleteSavedQueryResponseBody {
	return s.Body
}

func (s *DeleteSavedQueryResponse) SetHeaders(v map[string]*string) *DeleteSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *DeleteSavedQueryResponse) SetStatusCode(v int32) *DeleteSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSavedQueryResponse) SetBody(v *DeleteSavedQueryResponseBody) *DeleteSavedQueryResponse {
	s.Body = v
	return s
}

func (s *DeleteSavedQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
