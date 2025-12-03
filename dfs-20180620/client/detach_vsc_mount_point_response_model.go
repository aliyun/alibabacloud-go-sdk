// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachVscMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachVscMountPointResponse
	GetStatusCode() *int32
	SetBody(v *DetachVscMountPointResponseBody) *DetachVscMountPointResponse
	GetBody() *DetachVscMountPointResponseBody
}

type DetachVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachVscMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachVscMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachVscMountPointResponse) GetBody() *DetachVscMountPointResponseBody {
	return s.Body
}

func (s *DetachVscMountPointResponse) SetHeaders(v map[string]*string) *DetachVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *DetachVscMountPointResponse) SetStatusCode(v int32) *DetachVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachVscMountPointResponse) SetBody(v *DetachVscMountPointResponseBody) *DetachVscMountPointResponse {
	s.Body = v
	return s
}

func (s *DetachVscMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
