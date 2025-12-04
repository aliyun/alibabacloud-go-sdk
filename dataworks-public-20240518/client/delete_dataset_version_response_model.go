// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetVersionResponseBody) *DeleteDatasetVersionResponse
	GetBody() *DeleteDatasetVersionResponseBody
}

type DeleteDatasetVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetVersionResponse) GetBody() *DeleteDatasetVersionResponseBody {
	return s.Body
}

func (s *DeleteDatasetVersionResponse) SetHeaders(v map[string]*string) *DeleteDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetVersionResponse) SetStatusCode(v int32) *DeleteDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetVersionResponse) SetBody(v *DeleteDatasetVersionResponseBody) *DeleteDatasetVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
