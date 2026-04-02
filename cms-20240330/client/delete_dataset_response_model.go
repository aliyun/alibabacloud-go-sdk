// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetResponseBody) *DeleteDatasetResponse
	GetBody() *DeleteDatasetResponseBody
}

type DeleteDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetResponse) GetBody() *DeleteDatasetResponseBody {
	return s.Body
}

func (s *DeleteDatasetResponse) SetHeaders(v map[string]*string) *DeleteDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetResponse) SetStatusCode(v int32) *DeleteDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetResponse) SetBody(v *DeleteDatasetResponseBody) *DeleteDatasetResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
