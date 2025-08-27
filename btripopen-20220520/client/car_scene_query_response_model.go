// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarSceneQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CarSceneQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CarSceneQueryResponse
	GetStatusCode() *int32
	SetBody(v *CarSceneQueryResponseBody) *CarSceneQueryResponse
	GetBody() *CarSceneQueryResponseBody
}

type CarSceneQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CarSceneQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CarSceneQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CarSceneQueryResponse) GoString() string {
	return s.String()
}

func (s *CarSceneQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CarSceneQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CarSceneQueryResponse) GetBody() *CarSceneQueryResponseBody {
	return s.Body
}

func (s *CarSceneQueryResponse) SetHeaders(v map[string]*string) *CarSceneQueryResponse {
	s.Headers = v
	return s
}

func (s *CarSceneQueryResponse) SetStatusCode(v int32) *CarSceneQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CarSceneQueryResponse) SetBody(v *CarSceneQueryResponseBody) *CarSceneQueryResponse {
	s.Body = v
	return s
}

func (s *CarSceneQueryResponse) Validate() error {
	return dara.Validate(s)
}
