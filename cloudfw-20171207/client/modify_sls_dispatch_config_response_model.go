// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySlsDispatchConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySlsDispatchConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySlsDispatchConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifySlsDispatchConfigResponseBody) *ModifySlsDispatchConfigResponse
	GetBody() *ModifySlsDispatchConfigResponseBody
}

type ModifySlsDispatchConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySlsDispatchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySlsDispatchConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySlsDispatchConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySlsDispatchConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySlsDispatchConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySlsDispatchConfigResponse) GetBody() *ModifySlsDispatchConfigResponseBody {
	return s.Body
}

func (s *ModifySlsDispatchConfigResponse) SetHeaders(v map[string]*string) *ModifySlsDispatchConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySlsDispatchConfigResponse) SetStatusCode(v int32) *ModifySlsDispatchConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySlsDispatchConfigResponse) SetBody(v *ModifySlsDispatchConfigResponseBody) *ModifySlsDispatchConfigResponse {
	s.Body = v
	return s
}

func (s *ModifySlsDispatchConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
