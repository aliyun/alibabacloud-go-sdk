// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortAutoCcStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPortAutoCcStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPortAutoCcStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPortAutoCcStatusResponseBody) *ModifyPortAutoCcStatusResponse
	GetBody() *ModifyPortAutoCcStatusResponseBody
}

type ModifyPortAutoCcStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPortAutoCcStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPortAutoCcStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortAutoCcStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyPortAutoCcStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPortAutoCcStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPortAutoCcStatusResponse) GetBody() *ModifyPortAutoCcStatusResponseBody {
	return s.Body
}

func (s *ModifyPortAutoCcStatusResponse) SetHeaders(v map[string]*string) *ModifyPortAutoCcStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyPortAutoCcStatusResponse) SetStatusCode(v int32) *ModifyPortAutoCcStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPortAutoCcStatusResponse) SetBody(v *ModifyPortAutoCcStatusResponseBody) *ModifyPortAutoCcStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyPortAutoCcStatusResponse) Validate() error {
	return dara.Validate(s)
}
