// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultimodalDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultimodalDatasetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultimodalDatasetResponseBody) *UpdateMultimodalDatasetResponse
	GetBody() *UpdateMultimodalDatasetResponseBody
}

type UpdateMultimodalDatasetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultimodalDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultimodalDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalDatasetResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultimodalDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultimodalDatasetResponse) GetBody() *UpdateMultimodalDatasetResponseBody {
	return s.Body
}

func (s *UpdateMultimodalDatasetResponse) SetHeaders(v map[string]*string) *UpdateMultimodalDatasetResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultimodalDatasetResponse) SetStatusCode(v int32) *UpdateMultimodalDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultimodalDatasetResponse) SetBody(v *UpdateMultimodalDatasetResponseBody) *UpdateMultimodalDatasetResponse {
	s.Body = v
	return s
}

func (s *UpdateMultimodalDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
