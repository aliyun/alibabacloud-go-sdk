// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCrowdDatasetAndBindingDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveCrowdDatasetAndBindingDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveCrowdDatasetAndBindingDatasetResponse
	GetStatusCode() *int32
	SetBody(v *SaveCrowdDatasetAndBindingDatasetResponseBody) *SaveCrowdDatasetAndBindingDatasetResponse
	GetBody() *SaveCrowdDatasetAndBindingDatasetResponseBody
}

type SaveCrowdDatasetAndBindingDatasetResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveCrowdDatasetAndBindingDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveCrowdDatasetAndBindingDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveCrowdDatasetAndBindingDatasetResponse) GoString() string {
	return s.String()
}

func (s *SaveCrowdDatasetAndBindingDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveCrowdDatasetAndBindingDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveCrowdDatasetAndBindingDatasetResponse) GetBody() *SaveCrowdDatasetAndBindingDatasetResponseBody {
	return s.Body
}

func (s *SaveCrowdDatasetAndBindingDatasetResponse) SetHeaders(v map[string]*string) *SaveCrowdDatasetAndBindingDatasetResponse {
	s.Headers = v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponse) SetStatusCode(v int32) *SaveCrowdDatasetAndBindingDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponse) SetBody(v *SaveCrowdDatasetAndBindingDatasetResponseBody) *SaveCrowdDatasetAndBindingDatasetResponse {
	s.Body = v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
