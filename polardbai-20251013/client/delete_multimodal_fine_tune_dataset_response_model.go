// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalFineTuneDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultimodalFineTuneDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultimodalFineTuneDatasetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultimodalFineTuneDatasetResponseBody) *DeleteMultimodalFineTuneDatasetResponse
	GetBody() *DeleteMultimodalFineTuneDatasetResponseBody
}

type DeleteMultimodalFineTuneDatasetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultimodalFineTuneDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultimodalFineTuneDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalFineTuneDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalFineTuneDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultimodalFineTuneDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultimodalFineTuneDatasetResponse) GetBody() *DeleteMultimodalFineTuneDatasetResponseBody {
	return s.Body
}

func (s *DeleteMultimodalFineTuneDatasetResponse) SetHeaders(v map[string]*string) *DeleteMultimodalFineTuneDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponse) SetStatusCode(v int32) *DeleteMultimodalFineTuneDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponse) SetBody(v *DeleteMultimodalFineTuneDatasetResponseBody) *DeleteMultimodalFineTuneDatasetResponse {
	s.Body = v
	return s
}

func (s *DeleteMultimodalFineTuneDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
