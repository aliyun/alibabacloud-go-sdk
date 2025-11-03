// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBgpPeerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBgpPeerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBgpPeerResponse
	GetStatusCode() *int32
	SetBody(v *CreateBgpPeerResponseBody) *CreateBgpPeerResponse
	GetBody() *CreateBgpPeerResponseBody
}

type CreateBgpPeerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBgpPeerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBgpPeerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBgpPeerResponse) GoString() string {
	return s.String()
}

func (s *CreateBgpPeerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBgpPeerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBgpPeerResponse) GetBody() *CreateBgpPeerResponseBody {
	return s.Body
}

func (s *CreateBgpPeerResponse) SetHeaders(v map[string]*string) *CreateBgpPeerResponse {
	s.Headers = v
	return s
}

func (s *CreateBgpPeerResponse) SetStatusCode(v int32) *CreateBgpPeerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBgpPeerResponse) SetBody(v *CreateBgpPeerResponseBody) *CreateBgpPeerResponse {
	s.Body = v
	return s
}

func (s *CreateBgpPeerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
