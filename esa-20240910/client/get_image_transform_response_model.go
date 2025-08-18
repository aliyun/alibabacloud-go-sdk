// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTransformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageTransformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageTransformResponse
	GetStatusCode() *int32
	SetBody(v *GetImageTransformResponseBody) *GetImageTransformResponse
	GetBody() *GetImageTransformResponseBody
}

type GetImageTransformResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageTransformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageTransformResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageTransformResponse) GoString() string {
	return s.String()
}

func (s *GetImageTransformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageTransformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageTransformResponse) GetBody() *GetImageTransformResponseBody {
	return s.Body
}

func (s *GetImageTransformResponse) SetHeaders(v map[string]*string) *GetImageTransformResponse {
	s.Headers = v
	return s
}

func (s *GetImageTransformResponse) SetStatusCode(v int32) *GetImageTransformResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageTransformResponse) SetBody(v *GetImageTransformResponseBody) *GetImageTransformResponse {
	s.Body = v
	return s
}

func (s *GetImageTransformResponse) Validate() error {
	return dara.Validate(s)
}
