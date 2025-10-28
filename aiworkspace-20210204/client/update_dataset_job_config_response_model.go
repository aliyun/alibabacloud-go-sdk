// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDatasetJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDatasetJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDatasetJobConfigResponseBody) *UpdateDatasetJobConfigResponse
	GetBody() *UpdateDatasetJobConfigResponseBody
}

type UpdateDatasetJobConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetJobConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDatasetJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDatasetJobConfigResponse) GetBody() *UpdateDatasetJobConfigResponseBody {
	return s.Body
}

func (s *UpdateDatasetJobConfigResponse) SetHeaders(v map[string]*string) *UpdateDatasetJobConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetJobConfigResponse) SetStatusCode(v int32) *UpdateDatasetJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetJobConfigResponse) SetBody(v *UpdateDatasetJobConfigResponseBody) *UpdateDatasetJobConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateDatasetJobConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
