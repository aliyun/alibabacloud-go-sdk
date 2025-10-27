// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogHubStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLogHubStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLogHubStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLogHubStatusResponseBody) *ModifyLogHubStatusResponse
	GetBody() *ModifyLogHubStatusResponseBody
}

type ModifyLogHubStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLogHubStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLogHubStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogHubStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogHubStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLogHubStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLogHubStatusResponse) GetBody() *ModifyLogHubStatusResponseBody {
	return s.Body
}

func (s *ModifyLogHubStatusResponse) SetHeaders(v map[string]*string) *ModifyLogHubStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogHubStatusResponse) SetStatusCode(v int32) *ModifyLogHubStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogHubStatusResponse) SetBody(v *ModifyLogHubStatusResponseBody) *ModifyLogHubStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyLogHubStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
