// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasetJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasetJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasetJobConfigResponseBody) *DeleteDatasetJobConfigResponse
	GetBody() *DeleteDatasetJobConfigResponseBody
}

type DeleteDatasetJobConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetJobConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasetJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasetJobConfigResponse) GetBody() *DeleteDatasetJobConfigResponseBody {
	return s.Body
}

func (s *DeleteDatasetJobConfigResponse) SetHeaders(v map[string]*string) *DeleteDatasetJobConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetJobConfigResponse) SetStatusCode(v int32) *DeleteDatasetJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetJobConfigResponse) SetBody(v *DeleteDatasetJobConfigResponseBody) *DeleteDatasetJobConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasetJobConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
