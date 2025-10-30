// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleBatchMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSampleBatchMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSampleBatchMetaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSampleBatchMetaResponseBody) *DeleteSampleBatchMetaResponse
	GetBody() *DeleteSampleBatchMetaResponseBody
}

type DeleteSampleBatchMetaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSampleBatchMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSampleBatchMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleBatchMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteSampleBatchMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSampleBatchMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSampleBatchMetaResponse) GetBody() *DeleteSampleBatchMetaResponseBody {
	return s.Body
}

func (s *DeleteSampleBatchMetaResponse) SetHeaders(v map[string]*string) *DeleteSampleBatchMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteSampleBatchMetaResponse) SetStatusCode(v int32) *DeleteSampleBatchMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSampleBatchMetaResponse) SetBody(v *DeleteSampleBatchMetaResponseBody) *DeleteSampleBatchMetaResponse {
	s.Body = v
	return s
}

func (s *DeleteSampleBatchMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
