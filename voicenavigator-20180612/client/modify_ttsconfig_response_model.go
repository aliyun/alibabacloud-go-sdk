// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTTSConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTTSConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTTSConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTTSConfigResponseBody) *ModifyTTSConfigResponse
	GetBody() *ModifyTTSConfigResponseBody
}

type ModifyTTSConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTTSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTTSConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTTSConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTTSConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTTSConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTTSConfigResponse) GetBody() *ModifyTTSConfigResponseBody {
	return s.Body
}

func (s *ModifyTTSConfigResponse) SetHeaders(v map[string]*string) *ModifyTTSConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTTSConfigResponse) SetStatusCode(v int32) *ModifyTTSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTTSConfigResponse) SetBody(v *ModifyTTSConfigResponseBody) *ModifyTTSConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyTTSConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
