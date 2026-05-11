// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAsrConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAsrConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAsrConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAsrConfigResponseBody) *ModifyAsrConfigResponse
	GetBody() *ModifyAsrConfigResponseBody
}

type ModifyAsrConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAsrConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAsrConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAsrConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAsrConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAsrConfigResponse) GetBody() *ModifyAsrConfigResponseBody {
	return s.Body
}

func (s *ModifyAsrConfigResponse) SetHeaders(v map[string]*string) *ModifyAsrConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAsrConfigResponse) SetStatusCode(v int32) *ModifyAsrConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAsrConfigResponse) SetBody(v *ModifyAsrConfigResponseBody) *ModifyAsrConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAsrConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
