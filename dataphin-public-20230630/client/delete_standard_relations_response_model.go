// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardRelationsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardRelationsResponseBody) *DeleteStandardRelationsResponse
	GetBody() *DeleteStandardRelationsResponseBody
}

type DeleteStandardRelationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRelationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardRelationsResponse) GetBody() *DeleteStandardRelationsResponseBody {
	return s.Body
}

func (s *DeleteStandardRelationsResponse) SetHeaders(v map[string]*string) *DeleteStandardRelationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardRelationsResponse) SetStatusCode(v int32) *DeleteStandardRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardRelationsResponse) SetBody(v *DeleteStandardRelationsResponseBody) *DeleteStandardRelationsResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
