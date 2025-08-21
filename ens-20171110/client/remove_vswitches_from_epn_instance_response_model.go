// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVSwitchesFromEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveVSwitchesFromEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveVSwitchesFromEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveVSwitchesFromEpnInstanceResponseBody) *RemoveVSwitchesFromEpnInstanceResponse
	GetBody() *RemoveVSwitchesFromEpnInstanceResponseBody
}

type RemoveVSwitchesFromEpnInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveVSwitchesFromEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveVSwitchesFromEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) GetBody() *RemoveVSwitchesFromEpnInstanceResponseBody {
	return s.Body
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) SetHeaders(v map[string]*string) *RemoveVSwitchesFromEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) SetStatusCode(v int32) *RemoveVSwitchesFromEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) SetBody(v *RemoveVSwitchesFromEpnInstanceResponseBody) *RemoveVSwitchesFromEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) Validate() error {
	return dara.Validate(s)
}
