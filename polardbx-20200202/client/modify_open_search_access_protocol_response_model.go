// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchAccessProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOpenSearchAccessProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOpenSearchAccessProtocolResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOpenSearchAccessProtocolResponseBody) *ModifyOpenSearchAccessProtocolResponse
	GetBody() *ModifyOpenSearchAccessProtocolResponseBody
}

type ModifyOpenSearchAccessProtocolResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOpenSearchAccessProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOpenSearchAccessProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchAccessProtocolResponse) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchAccessProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOpenSearchAccessProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOpenSearchAccessProtocolResponse) GetBody() *ModifyOpenSearchAccessProtocolResponseBody {
	return s.Body
}

func (s *ModifyOpenSearchAccessProtocolResponse) SetHeaders(v map[string]*string) *ModifyOpenSearchAccessProtocolResponse {
	s.Headers = v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponse) SetStatusCode(v int32) *ModifyOpenSearchAccessProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponse) SetBody(v *ModifyOpenSearchAccessProtocolResponseBody) *ModifyOpenSearchAccessProtocolResponse {
	s.Body = v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
