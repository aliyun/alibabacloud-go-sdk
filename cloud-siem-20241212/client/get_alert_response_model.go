// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertResponseBody) *GetAlertResponse
	GetBody() *GetAlertResponseBody
}

type GetAlertResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertResponse) GoString() string {
	return s.String()
}

func (s *GetAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertResponse) GetBody() *GetAlertResponseBody {
	return s.Body
}

func (s *GetAlertResponse) SetHeaders(v map[string]*string) *GetAlertResponse {
	s.Headers = v
	return s
}

func (s *GetAlertResponse) SetStatusCode(v int32) *GetAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertResponse) SetBody(v *GetAlertResponseBody) *GetAlertResponse {
	s.Body = v
	return s
}

func (s *GetAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
