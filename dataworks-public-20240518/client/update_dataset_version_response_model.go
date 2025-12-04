// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDatasetVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDatasetVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDatasetVersionResponseBody) *UpdateDatasetVersionResponse
	GetBody() *UpdateDatasetVersionResponseBody
}

type UpdateDatasetVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDatasetVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDatasetVersionResponse) GetBody() *UpdateDatasetVersionResponseBody {
	return s.Body
}

func (s *UpdateDatasetVersionResponse) SetHeaders(v map[string]*string) *UpdateDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetVersionResponse) SetStatusCode(v int32) *UpdateDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetVersionResponse) SetBody(v *UpdateDatasetVersionResponseBody) *UpdateDatasetVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateDatasetVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
