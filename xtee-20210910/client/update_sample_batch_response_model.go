// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSampleBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSampleBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSampleBatchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSampleBatchResponseBody) *UpdateSampleBatchResponse
	GetBody() *UpdateSampleBatchResponseBody
}

type UpdateSampleBatchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSampleBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSampleBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSampleBatchResponse) GoString() string {
	return s.String()
}

func (s *UpdateSampleBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSampleBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSampleBatchResponse) GetBody() *UpdateSampleBatchResponseBody {
	return s.Body
}

func (s *UpdateSampleBatchResponse) SetHeaders(v map[string]*string) *UpdateSampleBatchResponse {
	s.Headers = v
	return s
}

func (s *UpdateSampleBatchResponse) SetStatusCode(v int32) *UpdateSampleBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSampleBatchResponse) SetBody(v *UpdateSampleBatchResponseBody) *UpdateSampleBatchResponse {
	s.Body = v
	return s
}

func (s *UpdateSampleBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
