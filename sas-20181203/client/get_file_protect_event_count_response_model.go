// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectEventCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectEventCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectEventCountResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectEventCountResponseBody) *GetFileProtectEventCountResponse
	GetBody() *GetFileProtectEventCountResponseBody
}

type GetFileProtectEventCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectEventCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectEventCountResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectEventCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectEventCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectEventCountResponse) GetBody() *GetFileProtectEventCountResponseBody {
	return s.Body
}

func (s *GetFileProtectEventCountResponse) SetHeaders(v map[string]*string) *GetFileProtectEventCountResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectEventCountResponse) SetStatusCode(v int32) *GetFileProtectEventCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectEventCountResponse) SetBody(v *GetFileProtectEventCountResponseBody) *GetFileProtectEventCountResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectEventCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
