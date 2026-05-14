// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkBindClientTelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkBindClientTelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkBindClientTelResponse
	GetStatusCode() *int32
	SetBody(v *ClinkBindClientTelResponseBody) *ClinkBindClientTelResponse
	GetBody() *ClinkBindClientTelResponseBody
}

type ClinkBindClientTelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkBindClientTelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkBindClientTelResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkBindClientTelResponse) GoString() string {
	return s.String()
}

func (s *ClinkBindClientTelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkBindClientTelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkBindClientTelResponse) GetBody() *ClinkBindClientTelResponseBody {
	return s.Body
}

func (s *ClinkBindClientTelResponse) SetHeaders(v map[string]*string) *ClinkBindClientTelResponse {
	s.Headers = v
	return s
}

func (s *ClinkBindClientTelResponse) SetStatusCode(v int32) *ClinkBindClientTelResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkBindClientTelResponse) SetBody(v *ClinkBindClientTelResponseBody) *ClinkBindClientTelResponse {
	s.Body = v
	return s
}

func (s *ClinkBindClientTelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
