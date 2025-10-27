// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVulConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVulConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVulConfigResponseBody) *ModifyVulConfigResponse
	GetBody() *ModifyVulConfigResponseBody
}

type ModifyVulConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVulConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVulConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyVulConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVulConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVulConfigResponse) GetBody() *ModifyVulConfigResponseBody {
	return s.Body
}

func (s *ModifyVulConfigResponse) SetHeaders(v map[string]*string) *ModifyVulConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyVulConfigResponse) SetStatusCode(v int32) *ModifyVulConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVulConfigResponse) SetBody(v *ModifyVulConfigResponseBody) *ModifyVulConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyVulConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
