// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceLogResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceLogResponseBody) *CreateResourceLogResponse
	GetBody() *CreateResourceLogResponseBody
}

type CreateResourceLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceLogResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceLogResponse) GetBody() *CreateResourceLogResponseBody {
	return s.Body
}

func (s *CreateResourceLogResponse) SetHeaders(v map[string]*string) *CreateResourceLogResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceLogResponse) SetStatusCode(v int32) *CreateResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceLogResponse) SetBody(v *CreateResourceLogResponseBody) *CreateResourceLogResponse {
	s.Body = v
	return s
}

func (s *CreateResourceLogResponse) Validate() error {
	return dara.Validate(s)
}
