// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMountInstanceSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MountInstanceSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MountInstanceSDGResponse
	GetStatusCode() *int32
	SetBody(v *MountInstanceSDGResponseBody) *MountInstanceSDGResponse
	GetBody() *MountInstanceSDGResponseBody
}

type MountInstanceSDGResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MountInstanceSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MountInstanceSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s MountInstanceSDGResponse) GoString() string {
	return s.String()
}

func (s *MountInstanceSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MountInstanceSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MountInstanceSDGResponse) GetBody() *MountInstanceSDGResponseBody {
	return s.Body
}

func (s *MountInstanceSDGResponse) SetHeaders(v map[string]*string) *MountInstanceSDGResponse {
	s.Headers = v
	return s
}

func (s *MountInstanceSDGResponse) SetStatusCode(v int32) *MountInstanceSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *MountInstanceSDGResponse) SetBody(v *MountInstanceSDGResponseBody) *MountInstanceSDGResponse {
	s.Body = v
	return s
}

func (s *MountInstanceSDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
