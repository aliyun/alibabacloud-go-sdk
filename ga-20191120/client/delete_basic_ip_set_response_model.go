// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicIpSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBasicIpSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBasicIpSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBasicIpSetResponseBody) *DeleteBasicIpSetResponse
	GetBody() *DeleteBasicIpSetResponseBody
}

type DeleteBasicIpSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBasicIpSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicIpSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBasicIpSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBasicIpSetResponse) GetBody() *DeleteBasicIpSetResponseBody {
	return s.Body
}

func (s *DeleteBasicIpSetResponse) SetHeaders(v map[string]*string) *DeleteBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicIpSetResponse) SetStatusCode(v int32) *DeleteBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicIpSetResponse) SetBody(v *DeleteBasicIpSetResponseBody) *DeleteBasicIpSetResponse {
	s.Body = v
	return s
}

func (s *DeleteBasicIpSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
