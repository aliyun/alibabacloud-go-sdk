// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalFineTuneDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalFineTuneDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalFineTuneDatasetResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalFineTuneDatasetResponseBody) *ListMultimodalFineTuneDatasetResponse
	GetBody() *ListMultimodalFineTuneDatasetResponseBody
}

type ListMultimodalFineTuneDatasetResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalFineTuneDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalFineTuneDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalFineTuneDatasetResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalFineTuneDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalFineTuneDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalFineTuneDatasetResponse) GetBody() *ListMultimodalFineTuneDatasetResponseBody {
	return s.Body
}

func (s *ListMultimodalFineTuneDatasetResponse) SetHeaders(v map[string]*string) *ListMultimodalFineTuneDatasetResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponse) SetStatusCode(v int32) *ListMultimodalFineTuneDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponse) SetBody(v *ListMultimodalFineTuneDatasetResponseBody) *ListMultimodalFineTuneDatasetResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalFineTuneDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
