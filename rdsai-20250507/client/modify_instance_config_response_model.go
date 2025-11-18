// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceConfigResponseBody) *ModifyInstanceConfigResponse
	GetBody() *ModifyInstanceConfigResponseBody
}

type ModifyInstanceConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceConfigResponse) GetBody() *ModifyInstanceConfigResponseBody {
	return s.Body
}

func (s *ModifyInstanceConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceConfigResponse) SetStatusCode(v int32) *ModifyInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceConfigResponse) SetBody(v *ModifyInstanceConfigResponseBody) *ModifyInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
