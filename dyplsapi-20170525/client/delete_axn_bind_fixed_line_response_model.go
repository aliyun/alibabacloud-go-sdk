// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxnBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAxnBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAxnBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAxnBindFixedLineResponseBody) *DeleteAxnBindFixedLineResponse
	GetBody() *DeleteAxnBindFixedLineResponseBody
}

type DeleteAxnBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxnBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxnBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAxnBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAxnBindFixedLineResponse) GetBody() *DeleteAxnBindFixedLineResponseBody {
	return s.Body
}

func (s *DeleteAxnBindFixedLineResponse) SetHeaders(v map[string]*string) *DeleteAxnBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxnBindFixedLineResponse) SetStatusCode(v int32) *DeleteAxnBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxnBindFixedLineResponse) SetBody(v *DeleteAxnBindFixedLineResponseBody) *DeleteAxnBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *DeleteAxnBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
