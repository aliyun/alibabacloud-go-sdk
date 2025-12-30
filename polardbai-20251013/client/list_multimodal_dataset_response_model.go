// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalDatasetResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalDatasetResponseBody) *ListMultimodalDatasetResponse
	GetBody() *ListMultimodalDatasetResponseBody
}

type ListMultimodalDatasetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalDatasetResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalDatasetResponse) GetBody() *ListMultimodalDatasetResponseBody {
	return s.Body
}

func (s *ListMultimodalDatasetResponse) SetHeaders(v map[string]*string) *ListMultimodalDatasetResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalDatasetResponse) SetStatusCode(v int32) *ListMultimodalDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalDatasetResponse) SetBody(v *ListMultimodalDatasetResponseBody) *ListMultimodalDatasetResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
