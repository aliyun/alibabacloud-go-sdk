// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMajorVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceMajorVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceMajorVersionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceMajorVersionResponseBody) *ModifyInstanceMajorVersionResponse
	GetBody() *ModifyInstanceMajorVersionResponseBody
}

type ModifyInstanceMajorVersionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMajorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMajorVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMajorVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMajorVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceMajorVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceMajorVersionResponse) GetBody() *ModifyInstanceMajorVersionResponseBody {
	return s.Body
}

func (s *ModifyInstanceMajorVersionResponse) SetHeaders(v map[string]*string) *ModifyInstanceMajorVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMajorVersionResponse) SetStatusCode(v int32) *ModifyInstanceMajorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMajorVersionResponse) SetBody(v *ModifyInstanceMajorVersionResponseBody) *ModifyInstanceMajorVersionResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceMajorVersionResponse) Validate() error {
	return dara.Validate(s)
}
