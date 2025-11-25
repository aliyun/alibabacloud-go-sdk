// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlackholeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBlackholeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBlackholeStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBlackholeStatusResponseBody) *ModifyBlackholeStatusResponse
	GetBody() *ModifyBlackholeStatusResponseBody
}

type ModifyBlackholeStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBlackholeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBlackholeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackholeStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyBlackholeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBlackholeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBlackholeStatusResponse) GetBody() *ModifyBlackholeStatusResponseBody {
	return s.Body
}

func (s *ModifyBlackholeStatusResponse) SetHeaders(v map[string]*string) *ModifyBlackholeStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyBlackholeStatusResponse) SetStatusCode(v int32) *ModifyBlackholeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBlackholeStatusResponse) SetBody(v *ModifyBlackholeStatusResponseBody) *ModifyBlackholeStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyBlackholeStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
