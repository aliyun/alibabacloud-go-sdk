// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkerResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkerResponseBody) *CreateWorkerResponse
	GetBody() *CreateWorkerResponseBody
}

type CreateWorkerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkerResponse) GetBody() *CreateWorkerResponseBody {
	return s.Body
}

func (s *CreateWorkerResponse) SetHeaders(v map[string]*string) *CreateWorkerResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkerResponse) SetStatusCode(v int32) *CreateWorkerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkerResponse) SetBody(v *CreateWorkerResponseBody) *CreateWorkerResponse {
	s.Body = v
	return s
}

func (s *CreateWorkerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
