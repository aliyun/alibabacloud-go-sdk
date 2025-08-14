// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSampleBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSampleBatchResponse
	GetStatusCode() *int32
	SetBody(v *CreateSampleBatchResponseBody) *CreateSampleBatchResponse
	GetBody() *CreateSampleBatchResponseBody
}

type CreateSampleBatchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSampleBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSampleBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateSampleBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSampleBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSampleBatchResponse) GetBody() *CreateSampleBatchResponseBody {
	return s.Body
}

func (s *CreateSampleBatchResponse) SetHeaders(v map[string]*string) *CreateSampleBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateSampleBatchResponse) SetStatusCode(v int32) *CreateSampleBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSampleBatchResponse) SetBody(v *CreateSampleBatchResponseBody) *CreateSampleBatchResponse {
	s.Body = v
	return s
}

func (s *CreateSampleBatchResponse) Validate() error {
	return dara.Validate(s)
}
