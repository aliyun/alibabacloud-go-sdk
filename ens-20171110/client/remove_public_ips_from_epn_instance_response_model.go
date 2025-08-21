// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePublicIpsFromEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePublicIpsFromEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePublicIpsFromEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RemovePublicIpsFromEpnInstanceResponseBody) *RemovePublicIpsFromEpnInstanceResponse
	GetBody() *RemovePublicIpsFromEpnInstanceResponseBody
}

type RemovePublicIpsFromEpnInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePublicIpsFromEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePublicIpsFromEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePublicIpsFromEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePublicIpsFromEpnInstanceResponse) GetBody() *RemovePublicIpsFromEpnInstanceResponseBody {
	return s.Body
}

func (s *RemovePublicIpsFromEpnInstanceResponse) SetHeaders(v map[string]*string) *RemovePublicIpsFromEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceResponse) SetStatusCode(v int32) *RemovePublicIpsFromEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceResponse) SetBody(v *RemovePublicIpsFromEpnInstanceResponseBody) *RemovePublicIpsFromEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceResponse) Validate() error {
	return dara.Validate(s)
}
