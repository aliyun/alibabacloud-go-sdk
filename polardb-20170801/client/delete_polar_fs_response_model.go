// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarFsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarFsResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarFsResponseBody) *DeletePolarFsResponse
	GetBody() *DeletePolarFsResponseBody
}

type DeletePolarFsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarFsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarFsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarFsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarFsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarFsResponse) GetBody() *DeletePolarFsResponseBody {
	return s.Body
}

func (s *DeletePolarFsResponse) SetHeaders(v map[string]*string) *DeletePolarFsResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarFsResponse) SetStatusCode(v int32) *DeletePolarFsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarFsResponse) SetBody(v *DeletePolarFsResponseBody) *DeletePolarFsResponse {
	s.Body = v
	return s
}

func (s *DeletePolarFsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
