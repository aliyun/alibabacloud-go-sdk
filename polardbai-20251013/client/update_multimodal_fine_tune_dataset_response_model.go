// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalFineTuneDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultimodalFineTuneDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultimodalFineTuneDatasetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultimodalFineTuneDatasetResponseBody) *UpdateMultimodalFineTuneDatasetResponse
	GetBody() *UpdateMultimodalFineTuneDatasetResponseBody
}

type UpdateMultimodalFineTuneDatasetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultimodalFineTuneDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultimodalFineTuneDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalFineTuneDatasetResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalFineTuneDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultimodalFineTuneDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultimodalFineTuneDatasetResponse) GetBody() *UpdateMultimodalFineTuneDatasetResponseBody {
	return s.Body
}

func (s *UpdateMultimodalFineTuneDatasetResponse) SetHeaders(v map[string]*string) *UpdateMultimodalFineTuneDatasetResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponse) SetStatusCode(v int32) *UpdateMultimodalFineTuneDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponse) SetBody(v *UpdateMultimodalFineTuneDatasetResponseBody) *UpdateMultimodalFineTuneDatasetResponse {
	s.Body = v
	return s
}

func (s *UpdateMultimodalFineTuneDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
