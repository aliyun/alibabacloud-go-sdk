// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMountPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMountPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMountPointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMountPointResponseBody) *ModifyMountPointResponse
	GetBody() *ModifyMountPointResponseBody
}

type ModifyMountPointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMountPointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMountPointResponse) GoString() string {
	return s.String()
}

func (s *ModifyMountPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMountPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMountPointResponse) GetBody() *ModifyMountPointResponseBody {
	return s.Body
}

func (s *ModifyMountPointResponse) SetHeaders(v map[string]*string) *ModifyMountPointResponse {
	s.Headers = v
	return s
}

func (s *ModifyMountPointResponse) SetStatusCode(v int32) *ModifyMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMountPointResponse) SetBody(v *ModifyMountPointResponseBody) *ModifyMountPointResponse {
	s.Body = v
	return s
}

func (s *ModifyMountPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
