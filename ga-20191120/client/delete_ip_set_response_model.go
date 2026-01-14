// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpSetResponseBody) *DeleteIpSetResponse
	GetBody() *DeleteIpSetResponseBody
}

type DeleteIpSetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpSetResponse) GetBody() *DeleteIpSetResponseBody {
	return s.Body
}

func (s *DeleteIpSetResponse) SetHeaders(v map[string]*string) *DeleteIpSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpSetResponse) SetStatusCode(v int32) *DeleteIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpSetResponse) SetBody(v *DeleteIpSetResponseBody) *DeleteIpSetResponse {
	s.Body = v
	return s
}

func (s *DeleteIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
