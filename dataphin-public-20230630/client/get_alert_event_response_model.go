// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertEventResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertEventResponseBody) *GetAlertEventResponse
	GetBody() *GetAlertEventResponseBody
}

type GetAlertEventResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponse) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertEventResponse) GetBody() *GetAlertEventResponseBody {
	return s.Body
}

func (s *GetAlertEventResponse) SetHeaders(v map[string]*string) *GetAlertEventResponse {
	s.Headers = v
	return s
}

func (s *GetAlertEventResponse) SetStatusCode(v int32) *GetAlertEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertEventResponse) SetBody(v *GetAlertEventResponseBody) *GetAlertEventResponse {
	s.Body = v
	return s
}

func (s *GetAlertEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
