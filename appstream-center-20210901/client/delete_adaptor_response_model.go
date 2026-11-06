// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdaptorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAdaptorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAdaptorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAdaptorResponseBody) *DeleteAdaptorResponse
	GetBody() *DeleteAdaptorResponseBody
}

type DeleteAdaptorResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdaptorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdaptorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdaptorResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdaptorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAdaptorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAdaptorResponse) GetBody() *DeleteAdaptorResponseBody {
	return s.Body
}

func (s *DeleteAdaptorResponse) SetHeaders(v map[string]*string) *DeleteAdaptorResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdaptorResponse) SetStatusCode(v int32) *DeleteAdaptorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdaptorResponse) SetBody(v *DeleteAdaptorResponseBody) *DeleteAdaptorResponse {
	s.Body = v
	return s
}

func (s *DeleteAdaptorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
