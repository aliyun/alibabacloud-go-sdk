// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIpProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIpProtectionResponse
	GetStatusCode() *int32
	SetBody(v *GetIpProtectionResponseBody) *GetIpProtectionResponse
	GetBody() *GetIpProtectionResponseBody
}

type GetIpProtectionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIpProtectionResponse) GoString() string {
	return s.String()
}

func (s *GetIpProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIpProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIpProtectionResponse) GetBody() *GetIpProtectionResponseBody {
	return s.Body
}

func (s *GetIpProtectionResponse) SetHeaders(v map[string]*string) *GetIpProtectionResponse {
	s.Headers = v
	return s
}

func (s *GetIpProtectionResponse) SetStatusCode(v int32) *GetIpProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpProtectionResponse) SetBody(v *GetIpProtectionResponseBody) *GetIpProtectionResponse {
	s.Body = v
	return s
}

func (s *GetIpProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
