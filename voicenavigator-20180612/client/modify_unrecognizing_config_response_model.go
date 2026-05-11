// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUnrecognizingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUnrecognizingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUnrecognizingConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUnrecognizingConfigResponseBody) *ModifyUnrecognizingConfigResponse
	GetBody() *ModifyUnrecognizingConfigResponseBody
}

type ModifyUnrecognizingConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUnrecognizingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUnrecognizingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUnrecognizingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyUnrecognizingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUnrecognizingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUnrecognizingConfigResponse) GetBody() *ModifyUnrecognizingConfigResponseBody {
	return s.Body
}

func (s *ModifyUnrecognizingConfigResponse) SetHeaders(v map[string]*string) *ModifyUnrecognizingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyUnrecognizingConfigResponse) SetStatusCode(v int32) *ModifyUnrecognizingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUnrecognizingConfigResponse) SetBody(v *ModifyUnrecognizingConfigResponseBody) *ModifyUnrecognizingConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyUnrecognizingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
