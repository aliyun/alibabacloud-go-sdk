// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetVersionResponseBody) *GetDatasetVersionResponse
	GetBody() *GetDatasetVersionResponseBody
}

type GetDatasetVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetVersionResponse) GetBody() *GetDatasetVersionResponseBody {
	return s.Body
}

func (s *GetDatasetVersionResponse) SetHeaders(v map[string]*string) *GetDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetVersionResponse) SetStatusCode(v int32) *GetDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetVersionResponse) SetBody(v *GetDatasetVersionResponseBody) *GetDatasetVersionResponse {
	s.Body = v
	return s
}

func (s *GetDatasetVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
