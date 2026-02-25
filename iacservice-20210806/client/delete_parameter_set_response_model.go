// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteParameterSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteParameterSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteParameterSetResponseBody) *DeleteParameterSetResponse
	GetBody() *DeleteParameterSetResponseBody
}

type DeleteParameterSetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParameterSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteParameterSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteParameterSetResponse) GetBody() *DeleteParameterSetResponseBody {
	return s.Body
}

func (s *DeleteParameterSetResponse) SetHeaders(v map[string]*string) *DeleteParameterSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterSetResponse) SetStatusCode(v int32) *DeleteParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParameterSetResponse) SetBody(v *DeleteParameterSetResponseBody) *DeleteParameterSetResponse {
	s.Body = v
	return s
}

func (s *DeleteParameterSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
