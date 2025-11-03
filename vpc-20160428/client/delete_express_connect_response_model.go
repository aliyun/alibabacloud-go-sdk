// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExpressConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExpressConnectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExpressConnectResponseBody) *DeleteExpressConnectResponse
	GetBody() *DeleteExpressConnectResponseBody
}

type DeleteExpressConnectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExpressConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExpressConnectResponse) GetBody() *DeleteExpressConnectResponseBody {
	return s.Body
}

func (s *DeleteExpressConnectResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectResponse) SetStatusCode(v int32) *DeleteExpressConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectResponse) SetBody(v *DeleteExpressConnectResponseBody) *DeleteExpressConnectResponse {
	s.Body = v
	return s
}

func (s *DeleteExpressConnectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
