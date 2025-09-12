// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancePayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstancePayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstancePayTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstancePayTypeResponseBody) *ModifyInstancePayTypeResponse
	GetBody() *ModifyInstancePayTypeResponseBody
}

type ModifyInstancePayTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstancePayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstancePayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePayTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstancePayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstancePayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstancePayTypeResponse) GetBody() *ModifyInstancePayTypeResponseBody {
	return s.Body
}

func (s *ModifyInstancePayTypeResponse) SetHeaders(v map[string]*string) *ModifyInstancePayTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstancePayTypeResponse) SetStatusCode(v int32) *ModifyInstancePayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstancePayTypeResponse) SetBody(v *ModifyInstancePayTypeResponseBody) *ModifyInstancePayTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstancePayTypeResponse) Validate() error {
	return dara.Validate(s)
}
