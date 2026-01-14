// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpSetsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpSetsResponseBody) *DeleteIpSetsResponse
	GetBody() *DeleteIpSetsResponseBody
}

type DeleteIpSetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpSetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpSetsResponse) GetBody() *DeleteIpSetsResponseBody {
	return s.Body
}

func (s *DeleteIpSetsResponse) SetHeaders(v map[string]*string) *DeleteIpSetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpSetsResponse) SetStatusCode(v int32) *DeleteIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpSetsResponse) SetBody(v *DeleteIpSetsResponseBody) *DeleteIpSetsResponse {
	s.Body = v
	return s
}

func (s *DeleteIpSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
