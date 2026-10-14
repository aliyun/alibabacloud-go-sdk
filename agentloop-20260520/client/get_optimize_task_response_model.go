// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptimizeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOptimizeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOptimizeTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetOptimizeTaskResponseBody) *GetOptimizeTaskResponse
	GetBody() *GetOptimizeTaskResponseBody
}

type GetOptimizeTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOptimizeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOptimizeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOptimizeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetOptimizeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOptimizeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOptimizeTaskResponse) GetBody() *GetOptimizeTaskResponseBody {
	return s.Body
}

func (s *GetOptimizeTaskResponse) SetHeaders(v map[string]*string) *GetOptimizeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetOptimizeTaskResponse) SetStatusCode(v int32) *GetOptimizeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOptimizeTaskResponse) SetBody(v *GetOptimizeTaskResponseBody) *GetOptimizeTaskResponse {
	s.Body = v
	return s
}

func (s *GetOptimizeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
