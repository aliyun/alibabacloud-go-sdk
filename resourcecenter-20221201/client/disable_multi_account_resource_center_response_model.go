// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMultiAccountResourceCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableMultiAccountResourceCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableMultiAccountResourceCenterResponse
	GetStatusCode() *int32
	SetBody(v *DisableMultiAccountResourceCenterResponseBody) *DisableMultiAccountResourceCenterResponse
	GetBody() *DisableMultiAccountResourceCenterResponseBody
}

type DisableMultiAccountResourceCenterResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableMultiAccountResourceCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableMultiAccountResourceCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableMultiAccountResourceCenterResponse) GetBody() *DisableMultiAccountResourceCenterResponseBody {
	return s.Body
}

func (s *DisableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *DisableMultiAccountResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *DisableMultiAccountResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetBody(v *DisableMultiAccountResourceCenterResponseBody) *DisableMultiAccountResourceCenterResponse {
	s.Body = v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) Validate() error {
	return dara.Validate(s)
}
