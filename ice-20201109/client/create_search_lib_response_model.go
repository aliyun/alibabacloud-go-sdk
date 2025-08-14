// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *CreateSearchLibResponseBody) *CreateSearchLibResponse
	GetBody() *CreateSearchLibResponseBody
}

type CreateSearchLibResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchLibResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSearchLibResponse) GetBody() *CreateSearchLibResponseBody {
	return s.Body
}

func (s *CreateSearchLibResponse) SetHeaders(v map[string]*string) *CreateSearchLibResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchLibResponse) SetStatusCode(v int32) *CreateSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSearchLibResponse) SetBody(v *CreateSearchLibResponseBody) *CreateSearchLibResponse {
	s.Body = v
	return s
}

func (s *CreateSearchLibResponse) Validate() error {
	return dara.Validate(s)
}
