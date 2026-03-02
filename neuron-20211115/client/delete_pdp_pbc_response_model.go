// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpPbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePdpPbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePdpPbcResponse
	GetStatusCode() *int32
	SetBody(v *DeletePdpPbcResponseBody) *DeletePdpPbcResponse
	GetBody() *DeletePdpPbcResponseBody
}

type DeletePdpPbcResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePdpPbcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePdpPbcResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpPbcResponse) GoString() string {
	return s.String()
}

func (s *DeletePdpPbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePdpPbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePdpPbcResponse) GetBody() *DeletePdpPbcResponseBody {
	return s.Body
}

func (s *DeletePdpPbcResponse) SetHeaders(v map[string]*string) *DeletePdpPbcResponse {
	s.Headers = v
	return s
}

func (s *DeletePdpPbcResponse) SetStatusCode(v int32) *DeletePdpPbcResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePdpPbcResponse) SetBody(v *DeletePdpPbcResponseBody) *DeletePdpPbcResponse {
	s.Body = v
	return s
}

func (s *DeletePdpPbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
