// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetJobConfigResponseBody) *CreateDatasetJobConfigResponse
	GetBody() *CreateDatasetJobConfigResponseBody
}

type CreateDatasetJobConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetJobConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetJobConfigResponse) GetBody() *CreateDatasetJobConfigResponseBody {
	return s.Body
}

func (s *CreateDatasetJobConfigResponse) SetHeaders(v map[string]*string) *CreateDatasetJobConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetJobConfigResponse) SetStatusCode(v int32) *CreateDatasetJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetJobConfigResponse) SetBody(v *CreateDatasetJobConfigResponseBody) *CreateDatasetJobConfigResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetJobConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
