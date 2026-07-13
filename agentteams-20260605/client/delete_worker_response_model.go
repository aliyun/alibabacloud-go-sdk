// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkerResponseBody) *DeleteWorkerResponse
	GetBody() *DeleteWorkerResponseBody
}

type DeleteWorkerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkerResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkerResponse) GetBody() *DeleteWorkerResponseBody {
	return s.Body
}

func (s *DeleteWorkerResponse) SetHeaders(v map[string]*string) *DeleteWorkerResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkerResponse) SetStatusCode(v int32) *DeleteWorkerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkerResponse) SetBody(v *DeleteWorkerResponseBody) *DeleteWorkerResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
