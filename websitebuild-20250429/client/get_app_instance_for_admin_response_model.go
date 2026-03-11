// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppInstanceForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppInstanceForAdminResponse
	GetStatusCode() *int32
	SetBody(v *GetAppInstanceForAdminResponseBody) *GetAppInstanceForAdminResponse
	GetBody() *GetAppInstanceForAdminResponseBody
}

type GetAppInstanceForAdminResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInstanceForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInstanceForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForAdminResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppInstanceForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppInstanceForAdminResponse) GetBody() *GetAppInstanceForAdminResponseBody {
	return s.Body
}

func (s *GetAppInstanceForAdminResponse) SetHeaders(v map[string]*string) *GetAppInstanceForAdminResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceForAdminResponse) SetStatusCode(v int32) *GetAppInstanceForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceForAdminResponse) SetBody(v *GetAppInstanceForAdminResponseBody) *GetAppInstanceForAdminResponse {
	s.Body = v
	return s
}

func (s *GetAppInstanceForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
