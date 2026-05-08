// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryImageToVideoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryImageToVideoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryImageToVideoTaskResponse
	GetStatusCode() *int32
	SetBody(v *QueryImageToVideoTaskResponseBody) *QueryImageToVideoTaskResponse
	GetBody() *QueryImageToVideoTaskResponseBody
}

type QueryImageToVideoTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryImageToVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryImageToVideoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryImageToVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryImageToVideoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryImageToVideoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryImageToVideoTaskResponse) GetBody() *QueryImageToVideoTaskResponseBody {
	return s.Body
}

func (s *QueryImageToVideoTaskResponse) SetHeaders(v map[string]*string) *QueryImageToVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryImageToVideoTaskResponse) SetStatusCode(v int32) *QueryImageToVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryImageToVideoTaskResponse) SetBody(v *QueryImageToVideoTaskResponseBody) *QueryImageToVideoTaskResponse {
	s.Body = v
	return s
}

func (s *QueryImageToVideoTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
