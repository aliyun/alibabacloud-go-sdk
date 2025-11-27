// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCInstanceAttributeResponseBody) *ModifyRCInstanceAttributeResponse
	GetBody() *ModifyRCInstanceAttributeResponseBody
}

type ModifyRCInstanceAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCInstanceAttributeResponse) GetBody() *ModifyRCInstanceAttributeResponseBody {
	return s.Body
}

func (s *ModifyRCInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyRCInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCInstanceAttributeResponse) SetStatusCode(v int32) *ModifyRCInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCInstanceAttributeResponse) SetBody(v *ModifyRCInstanceAttributeResponseBody) *ModifyRCInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyRCInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
