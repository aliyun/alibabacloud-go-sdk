// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkerResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkerResponseBody) *GetWorkerResponse
	GetBody() *GetWorkerResponseBody
}

type GetWorkerResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerResponse) GoString() string {
	return s.String()
}

func (s *GetWorkerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkerResponse) GetBody() *GetWorkerResponseBody {
	return s.Body
}

func (s *GetWorkerResponse) SetHeaders(v map[string]*string) *GetWorkerResponse {
	s.Headers = v
	return s
}

func (s *GetWorkerResponse) SetStatusCode(v int32) *GetWorkerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkerResponse) SetBody(v *GetWorkerResponseBody) *GetWorkerResponse {
	s.Body = v
	return s
}

func (s *GetWorkerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
