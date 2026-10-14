// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptimizeRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOptimizeRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOptimizeRunResponse
	GetStatusCode() *int32
	SetBody(v *GetOptimizeRunResponseBody) *GetOptimizeRunResponse
	GetBody() *GetOptimizeRunResponseBody
}

type GetOptimizeRunResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOptimizeRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOptimizeRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOptimizeRunResponse) GoString() string {
	return s.String()
}

func (s *GetOptimizeRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOptimizeRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOptimizeRunResponse) GetBody() *GetOptimizeRunResponseBody {
	return s.Body
}

func (s *GetOptimizeRunResponse) SetHeaders(v map[string]*string) *GetOptimizeRunResponse {
	s.Headers = v
	return s
}

func (s *GetOptimizeRunResponse) SetStatusCode(v int32) *GetOptimizeRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOptimizeRunResponse) SetBody(v *GetOptimizeRunResponseBody) *GetOptimizeRunResponse {
	s.Body = v
	return s
}

func (s *GetOptimizeRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
