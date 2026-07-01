// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse
	GetBody() *CreateDatasetResponseBody
}

type CreateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetResponse) GetBody() *CreateDatasetResponseBody {
	return s.Body
}

func (s *CreateDatasetResponse) SetHeaders(v map[string]*string) *CreateDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetResponse) SetStatusCode(v int32) *CreateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
