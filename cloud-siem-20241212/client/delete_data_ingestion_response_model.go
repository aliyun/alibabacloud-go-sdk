// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataIngestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataIngestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataIngestionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataIngestionResponseBody) *DeleteDataIngestionResponse
	GetBody() *DeleteDataIngestionResponseBody
}

type DeleteDataIngestionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataIngestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataIngestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataIngestionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataIngestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataIngestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataIngestionResponse) GetBody() *DeleteDataIngestionResponseBody {
	return s.Body
}

func (s *DeleteDataIngestionResponse) SetHeaders(v map[string]*string) *DeleteDataIngestionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataIngestionResponse) SetStatusCode(v int32) *DeleteDataIngestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataIngestionResponse) SetBody(v *DeleteDataIngestionResponseBody) *DeleteDataIngestionResponse {
	s.Body = v
	return s
}

func (s *DeleteDataIngestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
