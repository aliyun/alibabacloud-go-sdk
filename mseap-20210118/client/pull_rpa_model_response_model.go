// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullRpaModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PullRpaModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PullRpaModelResponse
	GetStatusCode() *int32
	SetBody(v *PullRpaModelResponseBody) *PullRpaModelResponse
	GetBody() *PullRpaModelResponseBody
}

type PullRpaModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PullRpaModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PullRpaModelResponse) String() string {
	return dara.Prettify(s)
}

func (s PullRpaModelResponse) GoString() string {
	return s.String()
}

func (s *PullRpaModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PullRpaModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PullRpaModelResponse) GetBody() *PullRpaModelResponseBody {
	return s.Body
}

func (s *PullRpaModelResponse) SetHeaders(v map[string]*string) *PullRpaModelResponse {
	s.Headers = v
	return s
}

func (s *PullRpaModelResponse) SetStatusCode(v int32) *PullRpaModelResponse {
	s.StatusCode = &v
	return s
}

func (s *PullRpaModelResponse) SetBody(v *PullRpaModelResponseBody) *PullRpaModelResponse {
	s.Body = v
	return s
}

func (s *PullRpaModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
