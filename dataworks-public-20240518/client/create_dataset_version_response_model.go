// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasetVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasetVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasetVersionResponseBody) *CreateDatasetVersionResponse
	GetBody() *CreateDatasetVersionResponseBody
}

type CreateDatasetVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasetVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasetVersionResponse) GetBody() *CreateDatasetVersionResponseBody {
	return s.Body
}

func (s *CreateDatasetVersionResponse) SetHeaders(v map[string]*string) *CreateDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetVersionResponse) SetStatusCode(v int32) *CreateDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetVersionResponse) SetBody(v *CreateDatasetVersionResponseBody) *CreateDatasetVersionResponse {
	s.Body = v
	return s
}

func (s *CreateDatasetVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
