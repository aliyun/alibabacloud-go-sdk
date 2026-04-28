// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMovePolarFsObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MovePolarFsObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MovePolarFsObjectsResponse
	GetStatusCode() *int32
	SetBody(v *MovePolarFsObjectsResponseBody) *MovePolarFsObjectsResponse
	GetBody() *MovePolarFsObjectsResponseBody
}

type MovePolarFsObjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MovePolarFsObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MovePolarFsObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s MovePolarFsObjectsResponse) GoString() string {
	return s.String()
}

func (s *MovePolarFsObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MovePolarFsObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MovePolarFsObjectsResponse) GetBody() *MovePolarFsObjectsResponseBody {
	return s.Body
}

func (s *MovePolarFsObjectsResponse) SetHeaders(v map[string]*string) *MovePolarFsObjectsResponse {
	s.Headers = v
	return s
}

func (s *MovePolarFsObjectsResponse) SetStatusCode(v int32) *MovePolarFsObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *MovePolarFsObjectsResponse) SetBody(v *MovePolarFsObjectsResponseBody) *MovePolarFsObjectsResponse {
	s.Body = v
	return s
}

func (s *MovePolarFsObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
