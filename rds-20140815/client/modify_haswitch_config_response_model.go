// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHASwitchConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHASwitchConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHASwitchConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHASwitchConfigResponseBody) *ModifyHASwitchConfigResponse
	GetBody() *ModifyHASwitchConfigResponseBody
}

type ModifyHASwitchConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHASwitchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHASwitchConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHASwitchConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyHASwitchConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHASwitchConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHASwitchConfigResponse) GetBody() *ModifyHASwitchConfigResponseBody {
	return s.Body
}

func (s *ModifyHASwitchConfigResponse) SetHeaders(v map[string]*string) *ModifyHASwitchConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyHASwitchConfigResponse) SetStatusCode(v int32) *ModifyHASwitchConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHASwitchConfigResponse) SetBody(v *ModifyHASwitchConfigResponseBody) *ModifyHASwitchConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyHASwitchConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
