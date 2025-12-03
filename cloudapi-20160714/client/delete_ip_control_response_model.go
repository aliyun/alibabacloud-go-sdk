// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpControlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpControlResponseBody) *DeleteIpControlResponse
	GetBody() *DeleteIpControlResponseBody
}

type DeleteIpControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpControlResponse) GetBody() *DeleteIpControlResponseBody {
	return s.Body
}

func (s *DeleteIpControlResponse) SetHeaders(v map[string]*string) *DeleteIpControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpControlResponse) SetStatusCode(v int32) *DeleteIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpControlResponse) SetBody(v *DeleteIpControlResponseBody) *DeleteIpControlResponse {
	s.Body = v
	return s
}

func (s *DeleteIpControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
