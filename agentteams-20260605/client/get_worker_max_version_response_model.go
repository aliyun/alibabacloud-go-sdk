// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerMaxVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkerMaxVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkerMaxVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkerMaxVersionResponseBody) *GetWorkerMaxVersionResponse
	GetBody() *GetWorkerMaxVersionResponseBody
}

type GetWorkerMaxVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkerMaxVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkerMaxVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerMaxVersionResponse) GoString() string {
	return s.String()
}

func (s *GetWorkerMaxVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkerMaxVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkerMaxVersionResponse) GetBody() *GetWorkerMaxVersionResponseBody {
	return s.Body
}

func (s *GetWorkerMaxVersionResponse) SetHeaders(v map[string]*string) *GetWorkerMaxVersionResponse {
	s.Headers = v
	return s
}

func (s *GetWorkerMaxVersionResponse) SetStatusCode(v int32) *GetWorkerMaxVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkerMaxVersionResponse) SetBody(v *GetWorkerMaxVersionResponseBody) *GetWorkerMaxVersionResponse {
	s.Body = v
	return s
}

func (s *GetWorkerMaxVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
