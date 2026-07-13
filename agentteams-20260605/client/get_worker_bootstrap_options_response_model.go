// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerBootstrapOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkerBootstrapOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkerBootstrapOptionsResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkerBootstrapOptionsResponseBody) *GetWorkerBootstrapOptionsResponse
	GetBody() *GetWorkerBootstrapOptionsResponseBody
}

type GetWorkerBootstrapOptionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkerBootstrapOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkerBootstrapOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerBootstrapOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetWorkerBootstrapOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkerBootstrapOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkerBootstrapOptionsResponse) GetBody() *GetWorkerBootstrapOptionsResponseBody {
	return s.Body
}

func (s *GetWorkerBootstrapOptionsResponse) SetHeaders(v map[string]*string) *GetWorkerBootstrapOptionsResponse {
	s.Headers = v
	return s
}

func (s *GetWorkerBootstrapOptionsResponse) SetStatusCode(v int32) *GetWorkerBootstrapOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkerBootstrapOptionsResponse) SetBody(v *GetWorkerBootstrapOptionsResponseBody) *GetWorkerBootstrapOptionsResponse {
	s.Body = v
	return s
}

func (s *GetWorkerBootstrapOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
