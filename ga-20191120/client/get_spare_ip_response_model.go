// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpareIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSpareIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSpareIpResponse
	GetStatusCode() *int32
	SetBody(v *GetSpareIpResponseBody) *GetSpareIpResponse
	GetBody() *GetSpareIpResponseBody
}

type GetSpareIpResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpareIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpareIpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSpareIpResponse) GoString() string {
	return s.String()
}

func (s *GetSpareIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSpareIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSpareIpResponse) GetBody() *GetSpareIpResponseBody {
	return s.Body
}

func (s *GetSpareIpResponse) SetHeaders(v map[string]*string) *GetSpareIpResponse {
	s.Headers = v
	return s
}

func (s *GetSpareIpResponse) SetStatusCode(v int32) *GetSpareIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpareIpResponse) SetBody(v *GetSpareIpResponseBody) *GetSpareIpResponse {
	s.Body = v
	return s
}

func (s *GetSpareIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
