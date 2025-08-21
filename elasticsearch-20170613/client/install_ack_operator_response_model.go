// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAckOperatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAckOperatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAckOperatorResponse
	GetStatusCode() *int32
	SetBody(v *InstallAckOperatorResponseBody) *InstallAckOperatorResponse
	GetBody() *InstallAckOperatorResponseBody
}

type InstallAckOperatorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAckOperatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAckOperatorResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAckOperatorResponse) GoString() string {
	return s.String()
}

func (s *InstallAckOperatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAckOperatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAckOperatorResponse) GetBody() *InstallAckOperatorResponseBody {
	return s.Body
}

func (s *InstallAckOperatorResponse) SetHeaders(v map[string]*string) *InstallAckOperatorResponse {
	s.Headers = v
	return s
}

func (s *InstallAckOperatorResponse) SetStatusCode(v int32) *InstallAckOperatorResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAckOperatorResponse) SetBody(v *InstallAckOperatorResponseBody) *InstallAckOperatorResponse {
	s.Body = v
	return s
}

func (s *InstallAckOperatorResponse) Validate() error {
	return dara.Validate(s)
}
