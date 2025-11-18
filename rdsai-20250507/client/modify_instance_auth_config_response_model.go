// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAuthConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceAuthConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceAuthConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceAuthConfigResponseBody) *ModifyInstanceAuthConfigResponse
	GetBody() *ModifyInstanceAuthConfigResponseBody
}

type ModifyInstanceAuthConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceAuthConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAuthConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceAuthConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceAuthConfigResponse) GetBody() *ModifyInstanceAuthConfigResponseBody {
	return s.Body
}

func (s *ModifyInstanceAuthConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAuthConfigResponse) SetStatusCode(v int32) *ModifyInstanceAuthConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAuthConfigResponse) SetBody(v *ModifyInstanceAuthConfigResponseBody) *ModifyInstanceAuthConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceAuthConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
