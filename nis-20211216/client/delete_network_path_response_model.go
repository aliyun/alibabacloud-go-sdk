// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkPathResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkPathResponseBody) *DeleteNetworkPathResponse
	GetBody() *DeleteNetworkPathResponseBody
}

type DeleteNetworkPathResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkPathResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPathResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkPathResponse) GetBody() *DeleteNetworkPathResponseBody {
	return s.Body
}

func (s *DeleteNetworkPathResponse) SetHeaders(v map[string]*string) *DeleteNetworkPathResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkPathResponse) SetStatusCode(v int32) *DeleteNetworkPathResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkPathResponse) SetBody(v *DeleteNetworkPathResponseBody) *DeleteNetworkPathResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
