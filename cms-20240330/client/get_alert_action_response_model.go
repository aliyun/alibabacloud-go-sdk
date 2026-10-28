// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertActionResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertActionResponseBody) *GetAlertActionResponse
	GetBody() *GetAlertActionResponseBody
}

type GetAlertActionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertActionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionResponse) GoString() string {
	return s.String()
}

func (s *GetAlertActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertActionResponse) GetBody() *GetAlertActionResponseBody {
	return s.Body
}

func (s *GetAlertActionResponse) SetHeaders(v map[string]*string) *GetAlertActionResponse {
	s.Headers = v
	return s
}

func (s *GetAlertActionResponse) SetStatusCode(v int32) *GetAlertActionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertActionResponse) SetBody(v *GetAlertActionResponseBody) *GetAlertActionResponse {
	s.Body = v
	return s
}

func (s *GetAlertActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
