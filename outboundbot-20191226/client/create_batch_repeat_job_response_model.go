// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchRepeatJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBatchRepeatJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBatchRepeatJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateBatchRepeatJobResponseBody) *CreateBatchRepeatJobResponse
	GetBody() *CreateBatchRepeatJobResponseBody
}

type CreateBatchRepeatJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchRepeatJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchRepeatJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchRepeatJobResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchRepeatJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBatchRepeatJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBatchRepeatJobResponse) GetBody() *CreateBatchRepeatJobResponseBody {
	return s.Body
}

func (s *CreateBatchRepeatJobResponse) SetHeaders(v map[string]*string) *CreateBatchRepeatJobResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchRepeatJobResponse) SetStatusCode(v int32) *CreateBatchRepeatJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchRepeatJobResponse) SetBody(v *CreateBatchRepeatJobResponseBody) *CreateBatchRepeatJobResponse {
	s.Body = v
	return s
}

func (s *CreateBatchRepeatJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
