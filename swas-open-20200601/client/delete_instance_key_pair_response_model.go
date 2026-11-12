// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceKeyPairResponseBody) *DeleteInstanceKeyPairResponse
	GetBody() *DeleteInstanceKeyPairResponseBody
}

type DeleteInstanceKeyPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceKeyPairResponse) GetBody() *DeleteInstanceKeyPairResponseBody {
	return s.Body
}

func (s *DeleteInstanceKeyPairResponse) SetHeaders(v map[string]*string) *DeleteInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceKeyPairResponse) SetStatusCode(v int32) *DeleteInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceKeyPairResponse) SetBody(v *DeleteInstanceKeyPairResponseBody) *DeleteInstanceKeyPairResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
