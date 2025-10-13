// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigMapResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigMapResponseBody) *DeleteConfigMapResponse
	GetBody() *DeleteConfigMapResponseBody
}

type DeleteConfigMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigMapResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigMapResponse) GetBody() *DeleteConfigMapResponseBody {
	return s.Body
}

func (s *DeleteConfigMapResponse) SetHeaders(v map[string]*string) *DeleteConfigMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigMapResponse) SetStatusCode(v int32) *DeleteConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigMapResponse) SetBody(v *DeleteConfigMapResponseBody) *DeleteConfigMapResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
