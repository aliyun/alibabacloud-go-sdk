// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelProviderResponseBody) *DeleteModelProviderResponse
	GetBody() *DeleteModelProviderResponseBody
}

type DeleteModelProviderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelProviderResponse) GetBody() *DeleteModelProviderResponseBody {
	return s.Body
}

func (s *DeleteModelProviderResponse) SetHeaders(v map[string]*string) *DeleteModelProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelProviderResponse) SetStatusCode(v int32) *DeleteModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelProviderResponse) SetBody(v *DeleteModelProviderResponseBody) *DeleteModelProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
