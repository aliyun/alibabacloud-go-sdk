// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLabelTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLabelTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLabelTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateLabelTableResponseBody) *CreateLabelTableResponse
	GetBody() *CreateLabelTableResponseBody
}

type CreateLabelTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLabelTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLabelTableResponse) GoString() string {
	return s.String()
}

func (s *CreateLabelTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLabelTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLabelTableResponse) GetBody() *CreateLabelTableResponseBody {
	return s.Body
}

func (s *CreateLabelTableResponse) SetHeaders(v map[string]*string) *CreateLabelTableResponse {
	s.Headers = v
	return s
}

func (s *CreateLabelTableResponse) SetStatusCode(v int32) *CreateLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLabelTableResponse) SetBody(v *CreateLabelTableResponseBody) *CreateLabelTableResponse {
	s.Body = v
	return s
}

func (s *CreateLabelTableResponse) Validate() error {
	return dara.Validate(s)
}
