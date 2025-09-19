// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientAlertModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientAlertModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientAlertModeResponse
	GetStatusCode() *int32
	SetBody(v *ListClientAlertModeResponseBody) *ListClientAlertModeResponse
	GetBody() *ListClientAlertModeResponseBody
}

type ListClientAlertModeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientAlertModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientAlertModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientAlertModeResponse) GoString() string {
	return s.String()
}

func (s *ListClientAlertModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientAlertModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientAlertModeResponse) GetBody() *ListClientAlertModeResponseBody {
	return s.Body
}

func (s *ListClientAlertModeResponse) SetHeaders(v map[string]*string) *ListClientAlertModeResponse {
	s.Headers = v
	return s
}

func (s *ListClientAlertModeResponse) SetStatusCode(v int32) *ListClientAlertModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientAlertModeResponse) SetBody(v *ListClientAlertModeResponseBody) *ListClientAlertModeResponse {
	s.Body = v
	return s
}

func (s *ListClientAlertModeResponse) Validate() error {
	return dara.Validate(s)
}
