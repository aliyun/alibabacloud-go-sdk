// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserEniVswitchOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserEniVswitchOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserEniVswitchOptionsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserEniVswitchOptionsResponseBody) *ModifyUserEniVswitchOptionsResponse
	GetBody() *ModifyUserEniVswitchOptionsResponseBody
}

type ModifyUserEniVswitchOptionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserEniVswitchOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserEniVswitchOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserEniVswitchOptionsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserEniVswitchOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserEniVswitchOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserEniVswitchOptionsResponse) GetBody() *ModifyUserEniVswitchOptionsResponseBody {
	return s.Body
}

func (s *ModifyUserEniVswitchOptionsResponse) SetHeaders(v map[string]*string) *ModifyUserEniVswitchOptionsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponse) SetStatusCode(v int32) *ModifyUserEniVswitchOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponse) SetBody(v *ModifyUserEniVswitchOptionsResponseBody) *ModifyUserEniVswitchOptionsResponse {
	s.Body = v
	return s
}

func (s *ModifyUserEniVswitchOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
