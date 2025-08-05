// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectBasicMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectBasicMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectBasicMetaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectBasicMetaResponseBody) *UpdateProjectBasicMetaResponse
	GetBody() *UpdateProjectBasicMetaResponseBody
}

type UpdateProjectBasicMetaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectBasicMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectBasicMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectBasicMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectBasicMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectBasicMetaResponse) GetBody() *UpdateProjectBasicMetaResponseBody {
	return s.Body
}

func (s *UpdateProjectBasicMetaResponse) SetHeaders(v map[string]*string) *UpdateProjectBasicMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectBasicMetaResponse) SetStatusCode(v int32) *UpdateProjectBasicMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectBasicMetaResponse) SetBody(v *UpdateProjectBasicMetaResponseBody) *UpdateProjectBasicMetaResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectBasicMetaResponse) Validate() error {
	return dara.Validate(s)
}
