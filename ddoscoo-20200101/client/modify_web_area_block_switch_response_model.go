// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAreaBlockSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebAreaBlockSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebAreaBlockSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebAreaBlockSwitchResponseBody) *ModifyWebAreaBlockSwitchResponse
	GetBody() *ModifyWebAreaBlockSwitchResponseBody
}

type ModifyWebAreaBlockSwitchResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebAreaBlockSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebAreaBlockSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAreaBlockSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebAreaBlockSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebAreaBlockSwitchResponse) GetBody() *ModifyWebAreaBlockSwitchResponseBody {
	return s.Body
}

func (s *ModifyWebAreaBlockSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebAreaBlockSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAreaBlockSwitchResponse) SetStatusCode(v int32) *ModifyWebAreaBlockSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchResponse) SetBody(v *ModifyWebAreaBlockSwitchResponseBody) *ModifyWebAreaBlockSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyWebAreaBlockSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
