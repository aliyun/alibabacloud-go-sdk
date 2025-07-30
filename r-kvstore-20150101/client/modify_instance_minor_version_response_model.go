// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMinorVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceMinorVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceMinorVersionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceMinorVersionResponseBody) *ModifyInstanceMinorVersionResponse
	GetBody() *ModifyInstanceMinorVersionResponseBody
}

type ModifyInstanceMinorVersionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMinorVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMinorVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceMinorVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceMinorVersionResponse) GetBody() *ModifyInstanceMinorVersionResponseBody {
	return s.Body
}

func (s *ModifyInstanceMinorVersionResponse) SetHeaders(v map[string]*string) *ModifyInstanceMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMinorVersionResponse) SetStatusCode(v int32) *ModifyInstanceMinorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMinorVersionResponse) SetBody(v *ModifyInstanceMinorVersionResponseBody) *ModifyInstanceMinorVersionResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceMinorVersionResponse) Validate() error {
	return dara.Validate(s)
}
