// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstanceWorkerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceInstanceWorkerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceInstanceWorkerResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceInstanceWorkerResponseBody) *ListResourceInstanceWorkerResponse
	GetBody() *ListResourceInstanceWorkerResponseBody
}

type ListResourceInstanceWorkerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceInstanceWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceInstanceWorkerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceInstanceWorkerResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceInstanceWorkerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceInstanceWorkerResponse) GetBody() *ListResourceInstanceWorkerResponseBody {
	return s.Body
}

func (s *ListResourceInstanceWorkerResponse) SetHeaders(v map[string]*string) *ListResourceInstanceWorkerResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInstanceWorkerResponse) SetStatusCode(v int32) *ListResourceInstanceWorkerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceInstanceWorkerResponse) SetBody(v *ListResourceInstanceWorkerResponseBody) *ListResourceInstanceWorkerResponse {
	s.Body = v
	return s
}

func (s *ListResourceInstanceWorkerResponse) Validate() error {
	return dara.Validate(s)
}
