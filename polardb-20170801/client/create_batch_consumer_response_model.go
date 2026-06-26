// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBatchConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBatchConsumerResponse
	GetStatusCode() *int32
	SetBody(v *CreateBatchConsumerResponseBody) *CreateBatchConsumerResponse
	GetBody() *CreateBatchConsumerResponseBody
}

type CreateBatchConsumerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchConsumerResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBatchConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBatchConsumerResponse) GetBody() *CreateBatchConsumerResponseBody {
	return s.Body
}

func (s *CreateBatchConsumerResponse) SetHeaders(v map[string]*string) *CreateBatchConsumerResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchConsumerResponse) SetStatusCode(v int32) *CreateBatchConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchConsumerResponse) SetBody(v *CreateBatchConsumerResponseBody) *CreateBatchConsumerResponse {
	s.Body = v
	return s
}

func (s *CreateBatchConsumerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
