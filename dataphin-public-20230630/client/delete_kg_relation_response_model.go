// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKgRelationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKgRelationResponseBody) *DeleteKgRelationResponse
	GetBody() *DeleteKgRelationResponseBody
}

type DeleteKgRelationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKgRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteKgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKgRelationResponse) GetBody() *DeleteKgRelationResponseBody {
	return s.Body
}

func (s *DeleteKgRelationResponse) SetHeaders(v map[string]*string) *DeleteKgRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteKgRelationResponse) SetStatusCode(v int32) *DeleteKgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKgRelationResponse) SetBody(v *DeleteKgRelationResponseBody) *DeleteKgRelationResponse {
	s.Body = v
	return s
}

func (s *DeleteKgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
