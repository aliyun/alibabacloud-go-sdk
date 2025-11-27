// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBatchResponse
	GetStatusCode() *int32
	SetBody(v *CreateBatchResponseBody) *CreateBatchResponse
	GetBody() *CreateBatchResponseBody
}

type CreateBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBatchResponse) GetBody() *CreateBatchResponseBody {
	return s.Body
}

func (s *CreateBatchResponse) SetHeaders(v map[string]*string) *CreateBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchResponse) SetStatusCode(v int32) *CreateBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchResponse) SetBody(v *CreateBatchResponseBody) *CreateBatchResponse {
	s.Body = v
	return s
}

func (s *CreateBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
