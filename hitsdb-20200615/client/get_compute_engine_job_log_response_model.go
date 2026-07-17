// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEngineJobLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeEngineJobLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeEngineJobLogResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeEngineJobLogResponseBody) *GetComputeEngineJobLogResponse
	GetBody() *GetComputeEngineJobLogResponseBody
}

type GetComputeEngineJobLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeEngineJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeEngineJobLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEngineJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetComputeEngineJobLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeEngineJobLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeEngineJobLogResponse) GetBody() *GetComputeEngineJobLogResponseBody {
	return s.Body
}

func (s *GetComputeEngineJobLogResponse) SetHeaders(v map[string]*string) *GetComputeEngineJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetComputeEngineJobLogResponse) SetStatusCode(v int32) *GetComputeEngineJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeEngineJobLogResponse) SetBody(v *GetComputeEngineJobLogResponseBody) *GetComputeEngineJobLogResponse {
	s.Body = v
	return s
}

func (s *GetComputeEngineJobLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
