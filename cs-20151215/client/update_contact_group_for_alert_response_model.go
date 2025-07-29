// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactGroupForAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContactGroupForAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContactGroupForAlertResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContactGroupForAlertResponseBody) *UpdateContactGroupForAlertResponse
	GetBody() *UpdateContactGroupForAlertResponseBody
}

type UpdateContactGroupForAlertResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContactGroupForAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContactGroupForAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactGroupForAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupForAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContactGroupForAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContactGroupForAlertResponse) GetBody() *UpdateContactGroupForAlertResponseBody {
	return s.Body
}

func (s *UpdateContactGroupForAlertResponse) SetHeaders(v map[string]*string) *UpdateContactGroupForAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactGroupForAlertResponse) SetStatusCode(v int32) *UpdateContactGroupForAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactGroupForAlertResponse) SetBody(v *UpdateContactGroupForAlertResponseBody) *UpdateContactGroupForAlertResponse {
	s.Body = v
	return s
}

func (s *UpdateContactGroupForAlertResponse) Validate() error {
	return dara.Validate(s)
}
