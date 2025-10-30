// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflinePipelineByAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflinePipelineByAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflinePipelineByAsyncResponse
	GetStatusCode() *int32
	SetBody(v *OfflinePipelineByAsyncResponseBody) *OfflinePipelineByAsyncResponse
	GetBody() *OfflinePipelineByAsyncResponseBody
}

type OfflinePipelineByAsyncResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflinePipelineByAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflinePipelineByAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineByAsyncResponse) GoString() string {
	return s.String()
}

func (s *OfflinePipelineByAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflinePipelineByAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflinePipelineByAsyncResponse) GetBody() *OfflinePipelineByAsyncResponseBody {
	return s.Body
}

func (s *OfflinePipelineByAsyncResponse) SetHeaders(v map[string]*string) *OfflinePipelineByAsyncResponse {
	s.Headers = v
	return s
}

func (s *OfflinePipelineByAsyncResponse) SetStatusCode(v int32) *OfflinePipelineByAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflinePipelineByAsyncResponse) SetBody(v *OfflinePipelineByAsyncResponseBody) *OfflinePipelineByAsyncResponse {
	s.Body = v
	return s
}

func (s *OfflinePipelineByAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
