// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceVpcAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCInstanceVpcAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCInstanceVpcAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCInstanceVpcAttributeResponseBody) *ModifyRCInstanceVpcAttributeResponse
	GetBody() *ModifyRCInstanceVpcAttributeResponseBody
}

type ModifyRCInstanceVpcAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCInstanceVpcAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCInstanceVpcAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceVpcAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceVpcAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCInstanceVpcAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCInstanceVpcAttributeResponse) GetBody() *ModifyRCInstanceVpcAttributeResponseBody {
	return s.Body
}

func (s *ModifyRCInstanceVpcAttributeResponse) SetHeaders(v map[string]*string) *ModifyRCInstanceVpcAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCInstanceVpcAttributeResponse) SetStatusCode(v int32) *ModifyRCInstanceVpcAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCInstanceVpcAttributeResponse) SetBody(v *ModifyRCInstanceVpcAttributeResponseBody) *ModifyRCInstanceVpcAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyRCInstanceVpcAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
