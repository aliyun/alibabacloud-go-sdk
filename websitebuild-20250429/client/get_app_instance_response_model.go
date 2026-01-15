// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetAppInstanceResponseBody) *GetAppInstanceResponse
	GetBody() *GetAppInstanceResponseBody
}

type GetAppInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppInstanceResponse) GetBody() *GetAppInstanceResponseBody {
	return s.Body
}

func (s *GetAppInstanceResponse) SetHeaders(v map[string]*string) *GetAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceResponse) SetStatusCode(v int32) *GetAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceResponse) SetBody(v *GetAppInstanceResponseBody) *GetAppInstanceResponse {
	s.Body = v
	return s
}

func (s *GetAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
