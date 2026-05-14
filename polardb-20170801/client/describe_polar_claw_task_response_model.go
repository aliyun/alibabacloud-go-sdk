// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawTaskResponseBody) *DescribePolarClawTaskResponse
	GetBody() *DescribePolarClawTaskResponseBody
}

type DescribePolarClawTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawTaskResponse) GetBody() *DescribePolarClawTaskResponseBody {
	return s.Body
}

func (s *DescribePolarClawTaskResponse) SetHeaders(v map[string]*string) *DescribePolarClawTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawTaskResponse) SetStatusCode(v int32) *DescribePolarClawTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawTaskResponse) SetBody(v *DescribePolarClawTaskResponseBody) *DescribePolarClawTaskResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
