// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamBatchJobMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamBatchJobMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamBatchJobMappingResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamBatchJobMappingResponseBody) *CreateStreamBatchJobMappingResponse
	GetBody() *CreateStreamBatchJobMappingResponseBody
}

type CreateStreamBatchJobMappingResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamBatchJobMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamBatchJobMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamBatchJobMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamBatchJobMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamBatchJobMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamBatchJobMappingResponse) GetBody() *CreateStreamBatchJobMappingResponseBody {
	return s.Body
}

func (s *CreateStreamBatchJobMappingResponse) SetHeaders(v map[string]*string) *CreateStreamBatchJobMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamBatchJobMappingResponse) SetStatusCode(v int32) *CreateStreamBatchJobMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamBatchJobMappingResponse) SetBody(v *CreateStreamBatchJobMappingResponseBody) *CreateStreamBatchJobMappingResponse {
	s.Body = v
	return s
}

func (s *CreateStreamBatchJobMappingResponse) Validate() error {
	return dara.Validate(s)
}
