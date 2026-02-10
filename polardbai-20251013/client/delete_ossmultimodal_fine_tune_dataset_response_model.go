// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOSSMultimodalFineTuneDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOSSMultimodalFineTuneDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOSSMultimodalFineTuneDatasetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOSSMultimodalFineTuneDatasetResponseBody) *DeleteOSSMultimodalFineTuneDatasetResponse
	GetBody() *DeleteOSSMultimodalFineTuneDatasetResponseBody
}

type DeleteOSSMultimodalFineTuneDatasetResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOSSMultimodalFineTuneDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOSSMultimodalFineTuneDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOSSMultimodalFineTuneDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponse) GetBody() *DeleteOSSMultimodalFineTuneDatasetResponseBody {
	return s.Body
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponse) SetHeaders(v map[string]*string) *DeleteOSSMultimodalFineTuneDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponse) SetStatusCode(v int32) *DeleteOSSMultimodalFineTuneDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponse) SetBody(v *DeleteOSSMultimodalFineTuneDatasetResponseBody) *DeleteOSSMultimodalFineTuneDatasetResponse {
	s.Body = v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
