// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLabelTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLabelTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLabelTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLabelTableResponseBody) *UpdateLabelTableResponse
	GetBody() *UpdateLabelTableResponseBody
}

type UpdateLabelTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLabelTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLabelTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLabelTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLabelTableResponse) GetBody() *UpdateLabelTableResponseBody {
	return s.Body
}

func (s *UpdateLabelTableResponse) SetHeaders(v map[string]*string) *UpdateLabelTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateLabelTableResponse) SetStatusCode(v int32) *UpdateLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLabelTableResponse) SetBody(v *UpdateLabelTableResponseBody) *UpdateLabelTableResponse {
	s.Body = v
	return s
}

func (s *UpdateLabelTableResponse) Validate() error {
	return dara.Validate(s)
}
