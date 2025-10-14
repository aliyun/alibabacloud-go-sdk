// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnloadRegionSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnloadRegionSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnloadRegionSDGResponse
	GetStatusCode() *int32
	SetBody(v *UnloadRegionSDGResponseBody) *UnloadRegionSDGResponse
	GetBody() *UnloadRegionSDGResponseBody
}

type UnloadRegionSDGResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnloadRegionSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnloadRegionSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s UnloadRegionSDGResponse) GoString() string {
	return s.String()
}

func (s *UnloadRegionSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnloadRegionSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnloadRegionSDGResponse) GetBody() *UnloadRegionSDGResponseBody {
	return s.Body
}

func (s *UnloadRegionSDGResponse) SetHeaders(v map[string]*string) *UnloadRegionSDGResponse {
	s.Headers = v
	return s
}

func (s *UnloadRegionSDGResponse) SetStatusCode(v int32) *UnloadRegionSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *UnloadRegionSDGResponse) SetBody(v *UnloadRegionSDGResponseBody) *UnloadRegionSDGResponse {
	s.Body = v
	return s
}

func (s *UnloadRegionSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
