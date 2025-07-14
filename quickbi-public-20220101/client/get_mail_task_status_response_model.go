// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMailTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMailTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMailTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetMailTaskStatusResponseBody) *GetMailTaskStatusResponse
	GetBody() *GetMailTaskStatusResponseBody
}

type GetMailTaskStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMailTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMailTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMailTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMailTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMailTaskStatusResponse) GetBody() *GetMailTaskStatusResponseBody {
	return s.Body
}

func (s *GetMailTaskStatusResponse) SetHeaders(v map[string]*string) *GetMailTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMailTaskStatusResponse) SetStatusCode(v int32) *GetMailTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMailTaskStatusResponse) SetBody(v *GetMailTaskStatusResponseBody) *GetMailTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetMailTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
