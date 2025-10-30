// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxbBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAxbBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAxbBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAxbBindFixedLineResponseBody) *DeleteAxbBindFixedLineResponse
	GetBody() *DeleteAxbBindFixedLineResponseBody
}

type DeleteAxbBindFixedLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxbBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxbBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxbBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAxbBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAxbBindFixedLineResponse) GetBody() *DeleteAxbBindFixedLineResponseBody {
	return s.Body
}

func (s *DeleteAxbBindFixedLineResponse) SetHeaders(v map[string]*string) *DeleteAxbBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxbBindFixedLineResponse) SetStatusCode(v int32) *DeleteAxbBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxbBindFixedLineResponse) SetBody(v *DeleteAxbBindFixedLineResponseBody) *DeleteAxbBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *DeleteAxbBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
