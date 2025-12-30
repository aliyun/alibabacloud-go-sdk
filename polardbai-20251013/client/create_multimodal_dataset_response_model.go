// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultimodalDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultimodalDatasetResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultimodalDatasetResponseBody) *CreateMultimodalDatasetResponse
	GetBody() *CreateMultimodalDatasetResponseBody
}

type CreateMultimodalDatasetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultimodalDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultimodalDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateMultimodalDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultimodalDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultimodalDatasetResponse) GetBody() *CreateMultimodalDatasetResponseBody {
	return s.Body
}

func (s *CreateMultimodalDatasetResponse) SetHeaders(v map[string]*string) *CreateMultimodalDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateMultimodalDatasetResponse) SetStatusCode(v int32) *CreateMultimodalDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultimodalDatasetResponse) SetBody(v *CreateMultimodalDatasetResponseBody) *CreateMultimodalDatasetResponse {
	s.Body = v
	return s
}

func (s *CreateMultimodalDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
