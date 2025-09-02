// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertMessageResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertMessageResponseBody) *GetAlertMessageResponse
	GetBody() *GetAlertMessageResponseBody
}

type GetAlertMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageResponse) GoString() string {
	return s.String()
}

func (s *GetAlertMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertMessageResponse) GetBody() *GetAlertMessageResponseBody {
	return s.Body
}

func (s *GetAlertMessageResponse) SetHeaders(v map[string]*string) *GetAlertMessageResponse {
	s.Headers = v
	return s
}

func (s *GetAlertMessageResponse) SetStatusCode(v int32) *GetAlertMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertMessageResponse) SetBody(v *GetAlertMessageResponseBody) *GetAlertMessageResponse {
	s.Body = v
	return s
}

func (s *GetAlertMessageResponse) Validate() error {
	return dara.Validate(s)
}
