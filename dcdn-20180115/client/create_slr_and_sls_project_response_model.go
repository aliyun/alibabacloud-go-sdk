// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrAndSlsProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSlrAndSlsProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSlrAndSlsProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateSlrAndSlsProjectResponseBody) *CreateSlrAndSlsProjectResponse
	GetBody() *CreateSlrAndSlsProjectResponseBody
}

type CreateSlrAndSlsProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSlrAndSlsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSlrAndSlsProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrAndSlsProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSlrAndSlsProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSlrAndSlsProjectResponse) GetBody() *CreateSlrAndSlsProjectResponseBody {
	return s.Body
}

func (s *CreateSlrAndSlsProjectResponse) SetHeaders(v map[string]*string) *CreateSlrAndSlsProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateSlrAndSlsProjectResponse) SetStatusCode(v int32) *CreateSlrAndSlsProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponse) SetBody(v *CreateSlrAndSlsProjectResponseBody) *CreateSlrAndSlsProjectResponse {
	s.Body = v
	return s
}

func (s *CreateSlrAndSlsProjectResponse) Validate() error {
	return dara.Validate(s)
}
