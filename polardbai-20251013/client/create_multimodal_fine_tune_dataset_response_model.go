// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalFineTuneDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultimodalFineTuneDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultimodalFineTuneDatasetResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultimodalFineTuneDatasetResponseBody) *CreateMultimodalFineTuneDatasetResponse
	GetBody() *CreateMultimodalFineTuneDatasetResponseBody
}

type CreateMultimodalFineTuneDatasetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultimodalFineTuneDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultimodalFineTuneDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalFineTuneDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateMultimodalFineTuneDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultimodalFineTuneDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultimodalFineTuneDatasetResponse) GetBody() *CreateMultimodalFineTuneDatasetResponseBody {
	return s.Body
}

func (s *CreateMultimodalFineTuneDatasetResponse) SetHeaders(v map[string]*string) *CreateMultimodalFineTuneDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponse) SetStatusCode(v int32) *CreateMultimodalFineTuneDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponse) SetBody(v *CreateMultimodalFineTuneDatasetResponseBody) *CreateMultimodalFineTuneDatasetResponse {
	s.Body = v
	return s
}

func (s *CreateMultimodalFineTuneDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
