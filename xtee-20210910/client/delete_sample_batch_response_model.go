// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSampleBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSampleBatchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSampleBatchResponseBody) *DeleteSampleBatchResponse
	GetBody() *DeleteSampleBatchResponseBody
}

type DeleteSampleBatchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSampleBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSampleBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleBatchResponse) GoString() string {
	return s.String()
}

func (s *DeleteSampleBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSampleBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSampleBatchResponse) GetBody() *DeleteSampleBatchResponseBody {
	return s.Body
}

func (s *DeleteSampleBatchResponse) SetHeaders(v map[string]*string) *DeleteSampleBatchResponse {
	s.Headers = v
	return s
}

func (s *DeleteSampleBatchResponse) SetStatusCode(v int32) *DeleteSampleBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSampleBatchResponse) SetBody(v *DeleteSampleBatchResponseBody) *DeleteSampleBatchResponse {
	s.Body = v
	return s
}

func (s *DeleteSampleBatchResponse) Validate() error {
	return dara.Validate(s)
}
