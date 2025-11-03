// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChainResponseBody) *DeleteChainResponse
	GetBody() *DeleteChainResponseBody
}

type DeleteChainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChainResponse) GoString() string {
	return s.String()
}

func (s *DeleteChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChainResponse) GetBody() *DeleteChainResponseBody {
	return s.Body
}

func (s *DeleteChainResponse) SetHeaders(v map[string]*string) *DeleteChainResponse {
	s.Headers = v
	return s
}

func (s *DeleteChainResponse) SetStatusCode(v int32) *DeleteChainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChainResponse) SetBody(v *DeleteChainResponseBody) *DeleteChainResponse {
	s.Body = v
	return s
}

func (s *DeleteChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
