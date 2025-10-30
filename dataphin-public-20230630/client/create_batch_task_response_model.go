// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateBatchTaskResponseBody) *CreateBatchTaskResponse
	GetBody() *CreateBatchTaskResponseBody
}

type CreateBatchTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBatchTaskResponse) GetBody() *CreateBatchTaskResponseBody {
	return s.Body
}

func (s *CreateBatchTaskResponse) SetHeaders(v map[string]*string) *CreateBatchTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchTaskResponse) SetStatusCode(v int32) *CreateBatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchTaskResponse) SetBody(v *CreateBatchTaskResponseBody) *CreateBatchTaskResponse {
	s.Body = v
	return s
}

func (s *CreateBatchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
