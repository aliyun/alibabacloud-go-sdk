// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarFsObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarFsObjectsResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarFsObjectsResponseBody) *DeletePolarFsObjectsResponse
	GetBody() *DeletePolarFsObjectsResponseBody
}

type DeletePolarFsObjectsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarFsObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarFsObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsObjectsResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarFsObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarFsObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarFsObjectsResponse) GetBody() *DeletePolarFsObjectsResponseBody {
	return s.Body
}

func (s *DeletePolarFsObjectsResponse) SetHeaders(v map[string]*string) *DeletePolarFsObjectsResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarFsObjectsResponse) SetStatusCode(v int32) *DeletePolarFsObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarFsObjectsResponse) SetBody(v *DeletePolarFsObjectsResponseBody) *DeletePolarFsObjectsResponse {
	s.Body = v
	return s
}

func (s *DeletePolarFsObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
