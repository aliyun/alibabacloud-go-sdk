// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewApgroupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NewApgroupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NewApgroupConfigResponse
	GetStatusCode() *int32
	SetBody(v *NewApgroupConfigResponseBody) *NewApgroupConfigResponse
	GetBody() *NewApgroupConfigResponseBody
}

type NewApgroupConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NewApgroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NewApgroupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s NewApgroupConfigResponse) GoString() string {
	return s.String()
}

func (s *NewApgroupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NewApgroupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NewApgroupConfigResponse) GetBody() *NewApgroupConfigResponseBody {
	return s.Body
}

func (s *NewApgroupConfigResponse) SetHeaders(v map[string]*string) *NewApgroupConfigResponse {
	s.Headers = v
	return s
}

func (s *NewApgroupConfigResponse) SetStatusCode(v int32) *NewApgroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *NewApgroupConfigResponse) SetBody(v *NewApgroupConfigResponseBody) *NewApgroupConfigResponse {
	s.Body = v
	return s
}

func (s *NewApgroupConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
