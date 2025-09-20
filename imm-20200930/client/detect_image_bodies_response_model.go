// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBodiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectImageBodiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectImageBodiesResponse
	GetStatusCode() *int32
	SetBody(v *DetectImageBodiesResponseBody) *DetectImageBodiesResponse
	GetBody() *DetectImageBodiesResponseBody
}

type DetectImageBodiesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectImageBodiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectImageBodiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBodiesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectImageBodiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectImageBodiesResponse) GetBody() *DetectImageBodiesResponseBody {
	return s.Body
}

func (s *DetectImageBodiesResponse) SetHeaders(v map[string]*string) *DetectImageBodiesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageBodiesResponse) SetStatusCode(v int32) *DetectImageBodiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageBodiesResponse) SetBody(v *DetectImageBodiesResponseBody) *DetectImageBodiesResponse {
	s.Body = v
	return s
}

func (s *DetectImageBodiesResponse) Validate() error {
	return dara.Validate(s)
}
