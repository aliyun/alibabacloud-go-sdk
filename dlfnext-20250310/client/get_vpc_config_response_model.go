// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcConfigResponseBody) *GetVpcConfigResponse
	GetBody() *GetVpcConfigResponseBody
}

type GetVpcConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVpcConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcConfigResponse) GetBody() *GetVpcConfigResponseBody {
	return s.Body
}

func (s *GetVpcConfigResponse) SetHeaders(v map[string]*string) *GetVpcConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVpcConfigResponse) SetStatusCode(v int32) *GetVpcConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcConfigResponse) SetBody(v *GetVpcConfigResponseBody) *GetVpcConfigResponse {
	s.Body = v
	return s
}

func (s *GetVpcConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
