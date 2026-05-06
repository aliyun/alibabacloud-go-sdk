// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKeylessServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetKeylessServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetKeylessServerResponse
	GetStatusCode() *int32
	SetBody(v *SetKeylessServerResponseBody) *SetKeylessServerResponse
	GetBody() *SetKeylessServerResponseBody
}

type SetKeylessServerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetKeylessServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetKeylessServerResponse) String() string {
	return dara.Prettify(s)
}

func (s SetKeylessServerResponse) GoString() string {
	return s.String()
}

func (s *SetKeylessServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetKeylessServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetKeylessServerResponse) GetBody() *SetKeylessServerResponseBody {
	return s.Body
}

func (s *SetKeylessServerResponse) SetHeaders(v map[string]*string) *SetKeylessServerResponse {
	s.Headers = v
	return s
}

func (s *SetKeylessServerResponse) SetStatusCode(v int32) *SetKeylessServerResponse {
	s.StatusCode = &v
	return s
}

func (s *SetKeylessServerResponse) SetBody(v *SetKeylessServerResponseBody) *SetKeylessServerResponse {
	s.Body = v
	return s
}

func (s *SetKeylessServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
