// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginProtectionResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginProtectionResponseBody) *GetOriginProtectionResponse
	GetBody() *GetOriginProtectionResponseBody
}

type GetOriginProtectionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionResponse) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginProtectionResponse) GetBody() *GetOriginProtectionResponseBody {
	return s.Body
}

func (s *GetOriginProtectionResponse) SetHeaders(v map[string]*string) *GetOriginProtectionResponse {
	s.Headers = v
	return s
}

func (s *GetOriginProtectionResponse) SetStatusCode(v int32) *GetOriginProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginProtectionResponse) SetBody(v *GetOriginProtectionResponseBody) *GetOriginProtectionResponse {
	s.Body = v
	return s
}

func (s *GetOriginProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
