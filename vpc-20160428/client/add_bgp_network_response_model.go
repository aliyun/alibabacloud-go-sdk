// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBgpNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBgpNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBgpNetworkResponse
	GetStatusCode() *int32
	SetBody(v *AddBgpNetworkResponseBody) *AddBgpNetworkResponse
	GetBody() *AddBgpNetworkResponseBody
}

type AddBgpNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBgpNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBgpNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBgpNetworkResponse) GoString() string {
	return s.String()
}

func (s *AddBgpNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBgpNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBgpNetworkResponse) GetBody() *AddBgpNetworkResponseBody {
	return s.Body
}

func (s *AddBgpNetworkResponse) SetHeaders(v map[string]*string) *AddBgpNetworkResponse {
	s.Headers = v
	return s
}

func (s *AddBgpNetworkResponse) SetStatusCode(v int32) *AddBgpNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBgpNetworkResponse) SetBody(v *AddBgpNetworkResponseBody) *AddBgpNetworkResponse {
	s.Body = v
	return s
}

func (s *AddBgpNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
