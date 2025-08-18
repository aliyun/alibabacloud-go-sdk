// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTransformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageTransformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageTransformResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageTransformResponseBody) *CreateImageTransformResponse
	GetBody() *CreateImageTransformResponseBody
}

type CreateImageTransformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageTransformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageTransformResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTransformResponse) GoString() string {
	return s.String()
}

func (s *CreateImageTransformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageTransformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageTransformResponse) GetBody() *CreateImageTransformResponseBody {
	return s.Body
}

func (s *CreateImageTransformResponse) SetHeaders(v map[string]*string) *CreateImageTransformResponse {
	s.Headers = v
	return s
}

func (s *CreateImageTransformResponse) SetStatusCode(v int32) *CreateImageTransformResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageTransformResponse) SetBody(v *CreateImageTransformResponseBody) *CreateImageTransformResponse {
	s.Body = v
	return s
}

func (s *CreateImageTransformResponse) Validate() error {
	return dara.Validate(s)
}
