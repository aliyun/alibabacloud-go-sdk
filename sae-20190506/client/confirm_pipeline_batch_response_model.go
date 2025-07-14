// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPipelineBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmPipelineBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmPipelineBatchResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmPipelineBatchResponseBody) *ConfirmPipelineBatchResponse
	GetBody() *ConfirmPipelineBatchResponseBody
}

type ConfirmPipelineBatchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmPipelineBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmPipelineBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPipelineBatchResponse) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmPipelineBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmPipelineBatchResponse) GetBody() *ConfirmPipelineBatchResponseBody {
	return s.Body
}

func (s *ConfirmPipelineBatchResponse) SetHeaders(v map[string]*string) *ConfirmPipelineBatchResponse {
	s.Headers = v
	return s
}

func (s *ConfirmPipelineBatchResponse) SetStatusCode(v int32) *ConfirmPipelineBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmPipelineBatchResponse) SetBody(v *ConfirmPipelineBatchResponseBody) *ConfirmPipelineBatchResponse {
	s.Body = v
	return s
}

func (s *ConfirmPipelineBatchResponse) Validate() error {
	return dara.Validate(s)
}
