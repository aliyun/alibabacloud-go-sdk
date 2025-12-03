// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPluginResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPluginResponseBody) *ModifyPluginResponse
	GetBody() *ModifyPluginResponseBody
}

type ModifyPluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPluginResponse) GoString() string {
	return s.String()
}

func (s *ModifyPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPluginResponse) GetBody() *ModifyPluginResponseBody {
	return s.Body
}

func (s *ModifyPluginResponse) SetHeaders(v map[string]*string) *ModifyPluginResponse {
	s.Headers = v
	return s
}

func (s *ModifyPluginResponse) SetStatusCode(v int32) *ModifyPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPluginResponse) SetBody(v *ModifyPluginResponseBody) *ModifyPluginResponse {
	s.Body = v
	return s
}

func (s *ModifyPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
