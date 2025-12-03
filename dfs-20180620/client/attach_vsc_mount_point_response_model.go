// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachVscMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachVscMountPointResponse
	GetStatusCode() *int32
	SetBody(v *AttachVscMountPointResponseBody) *AttachVscMountPointResponse
	GetBody() *AttachVscMountPointResponseBody
}

type AttachVscMountPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachVscMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachVscMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachVscMountPointResponse) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachVscMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachVscMountPointResponse) GetBody() *AttachVscMountPointResponseBody {
	return s.Body
}

func (s *AttachVscMountPointResponse) SetHeaders(v map[string]*string) *AttachVscMountPointResponse {
	s.Headers = v
	return s
}

func (s *AttachVscMountPointResponse) SetStatusCode(v int32) *AttachVscMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVscMountPointResponse) SetBody(v *AttachVscMountPointResponseBody) *AttachVscMountPointResponse {
	s.Body = v
	return s
}

func (s *AttachVscMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
