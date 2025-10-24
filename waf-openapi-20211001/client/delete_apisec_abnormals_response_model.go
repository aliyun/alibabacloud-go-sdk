// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApisecAbnormalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApisecAbnormalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApisecAbnormalsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApisecAbnormalsResponseBody) *DeleteApisecAbnormalsResponse
	GetBody() *DeleteApisecAbnormalsResponseBody
}

type DeleteApisecAbnormalsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApisecAbnormalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApisecAbnormalsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApisecAbnormalsResponse) GoString() string {
	return s.String()
}

func (s *DeleteApisecAbnormalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApisecAbnormalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApisecAbnormalsResponse) GetBody() *DeleteApisecAbnormalsResponseBody {
	return s.Body
}

func (s *DeleteApisecAbnormalsResponse) SetHeaders(v map[string]*string) *DeleteApisecAbnormalsResponse {
	s.Headers = v
	return s
}

func (s *DeleteApisecAbnormalsResponse) SetStatusCode(v int32) *DeleteApisecAbnormalsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApisecAbnormalsResponse) SetBody(v *DeleteApisecAbnormalsResponseBody) *DeleteApisecAbnormalsResponse {
	s.Body = v
	return s
}

func (s *DeleteApisecAbnormalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
