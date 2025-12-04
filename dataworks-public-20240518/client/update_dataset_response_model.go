// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDatasetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse
	GetBody() *UpdateDatasetResponseBody
}

type UpdateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDatasetResponse) GetBody() *UpdateDatasetResponseBody {
	return s.Body
}

func (s *UpdateDatasetResponse) SetHeaders(v map[string]*string) *UpdateDatasetResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetResponse) SetStatusCode(v int32) *UpdateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetResponse) SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse {
	s.Body = v
	return s
}

func (s *UpdateDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
