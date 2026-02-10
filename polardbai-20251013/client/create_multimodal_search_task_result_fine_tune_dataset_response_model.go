// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskResultFineTuneDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultimodalSearchTaskResultFineTuneDatasetResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) *CreateMultimodalSearchTaskResultFineTuneDatasetResponse
	GetBody() *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody
}

type CreateMultimodalSearchTaskResultFineTuneDatasetResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultimodalSearchTaskResultFineTuneDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskResultFineTuneDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponse) GetBody() *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody {
	return s.Body
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponse) SetHeaders(v map[string]*string) *CreateMultimodalSearchTaskResultFineTuneDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponse) SetStatusCode(v int32) *CreateMultimodalSearchTaskResultFineTuneDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponse) SetBody(v *CreateMultimodalSearchTaskResultFineTuneDatasetResponseBody) *CreateMultimodalSearchTaskResultFineTuneDatasetResponse {
	s.Body = v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
