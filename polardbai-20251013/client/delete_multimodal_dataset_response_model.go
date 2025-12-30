// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultimodalDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultimodalDatasetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultimodalDatasetResponseBody) *DeleteMultimodalDatasetResponse
	GetBody() *DeleteMultimodalDatasetResponseBody
}

type DeleteMultimodalDatasetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultimodalDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultimodalDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultimodalDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultimodalDatasetResponse) GetBody() *DeleteMultimodalDatasetResponseBody {
	return s.Body
}

func (s *DeleteMultimodalDatasetResponse) SetHeaders(v map[string]*string) *DeleteMultimodalDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultimodalDatasetResponse) SetStatusCode(v int32) *DeleteMultimodalDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultimodalDatasetResponse) SetBody(v *DeleteMultimodalDatasetResponseBody) *DeleteMultimodalDatasetResponse {
	s.Body = v
	return s
}

func (s *DeleteMultimodalDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
