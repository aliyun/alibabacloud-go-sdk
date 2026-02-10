// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOSSMultimodalFineTuneDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOSSMultimodalFineTuneDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOSSMultimodalFineTuneDatasetResponse
	GetStatusCode() *int32
	SetBody(v *AddOSSMultimodalFineTuneDatasetResponseBody) *AddOSSMultimodalFineTuneDatasetResponse
	GetBody() *AddOSSMultimodalFineTuneDatasetResponseBody
}

type AddOSSMultimodalFineTuneDatasetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOSSMultimodalFineTuneDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOSSMultimodalFineTuneDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOSSMultimodalFineTuneDatasetResponse) GoString() string {
	return s.String()
}

func (s *AddOSSMultimodalFineTuneDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOSSMultimodalFineTuneDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOSSMultimodalFineTuneDatasetResponse) GetBody() *AddOSSMultimodalFineTuneDatasetResponseBody {
	return s.Body
}

func (s *AddOSSMultimodalFineTuneDatasetResponse) SetHeaders(v map[string]*string) *AddOSSMultimodalFineTuneDatasetResponse {
	s.Headers = v
	return s
}

func (s *AddOSSMultimodalFineTuneDatasetResponse) SetStatusCode(v int32) *AddOSSMultimodalFineTuneDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOSSMultimodalFineTuneDatasetResponse) SetBody(v *AddOSSMultimodalFineTuneDatasetResponseBody) *AddOSSMultimodalFineTuneDatasetResponse {
	s.Body = v
	return s
}

func (s *AddOSSMultimodalFineTuneDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
