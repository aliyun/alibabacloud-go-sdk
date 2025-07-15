// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorFilterRuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrafficMirrorFilterRuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrafficMirrorFilterRuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrafficMirrorFilterRuleAttributeResponseBody) *UpdateTrafficMirrorFilterRuleAttributeResponse
	GetBody() *UpdateTrafficMirrorFilterRuleAttributeResponseBody
}

type UpdateTrafficMirrorFilterRuleAttributeResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficMirrorFilterRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficMirrorFilterRuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorFilterRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponse) GetBody() *UpdateTrafficMirrorFilterRuleAttributeResponseBody {
	return s.Body
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateTrafficMirrorFilterRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponse) SetStatusCode(v int32) *UpdateTrafficMirrorFilterRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponse) SetBody(v *UpdateTrafficMirrorFilterRuleAttributeResponseBody) *UpdateTrafficMirrorFilterRuleAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponse) Validate() error {
	return dara.Validate(s)
}
