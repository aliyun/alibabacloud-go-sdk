// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableModelInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableModelInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableModelInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableModelInfoResponseBody) *UpdateTableModelInfoResponse
	GetBody() *UpdateTableModelInfoResponseBody
}

type UpdateTableModelInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableModelInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableModelInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableModelInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableModelInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableModelInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableModelInfoResponse) GetBody() *UpdateTableModelInfoResponseBody {
	return s.Body
}

func (s *UpdateTableModelInfoResponse) SetHeaders(v map[string]*string) *UpdateTableModelInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableModelInfoResponse) SetStatusCode(v int32) *UpdateTableModelInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableModelInfoResponse) SetBody(v *UpdateTableModelInfoResponseBody) *UpdateTableModelInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateTableModelInfoResponse) Validate() error {
	return dara.Validate(s)
}
