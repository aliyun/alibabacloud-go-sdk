// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectDatasetResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectDatasetResponseBody) *GetProjectDatasetResponse
	GetBody() *GetProjectDatasetResponseBody
}

type GetProjectDatasetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetProjectDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectDatasetResponse) GetBody() *GetProjectDatasetResponseBody {
	return s.Body
}

func (s *GetProjectDatasetResponse) SetHeaders(v map[string]*string) *GetProjectDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetProjectDatasetResponse) SetStatusCode(v int32) *GetProjectDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectDatasetResponse) SetBody(v *GetProjectDatasetResponseBody) *GetProjectDatasetResponse {
	s.Body = v
	return s
}

func (s *GetProjectDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
