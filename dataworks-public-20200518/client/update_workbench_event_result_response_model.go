// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkbenchEventResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkbenchEventResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkbenchEventResultResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkbenchEventResultResponseBody) *UpdateWorkbenchEventResultResponse
	GetBody() *UpdateWorkbenchEventResultResponseBody
}

type UpdateWorkbenchEventResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkbenchEventResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkbenchEventResultResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkbenchEventResultResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkbenchEventResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkbenchEventResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkbenchEventResultResponse) GetBody() *UpdateWorkbenchEventResultResponseBody {
	return s.Body
}

func (s *UpdateWorkbenchEventResultResponse) SetHeaders(v map[string]*string) *UpdateWorkbenchEventResultResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkbenchEventResultResponse) SetStatusCode(v int32) *UpdateWorkbenchEventResultResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkbenchEventResultResponse) SetBody(v *UpdateWorkbenchEventResultResponseBody) *UpdateWorkbenchEventResultResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkbenchEventResultResponse) Validate() error {
	return dara.Validate(s)
}
