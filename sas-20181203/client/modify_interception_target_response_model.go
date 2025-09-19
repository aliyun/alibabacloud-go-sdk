// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInterceptionTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInterceptionTargetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInterceptionTargetResponseBody) *ModifyInterceptionTargetResponse
	GetBody() *ModifyInterceptionTargetResponseBody
}

type ModifyInterceptionTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInterceptionTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInterceptionTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInterceptionTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInterceptionTargetResponse) GetBody() *ModifyInterceptionTargetResponseBody {
	return s.Body
}

func (s *ModifyInterceptionTargetResponse) SetHeaders(v map[string]*string) *ModifyInterceptionTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyInterceptionTargetResponse) SetStatusCode(v int32) *ModifyInterceptionTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInterceptionTargetResponse) SetBody(v *ModifyInterceptionTargetResponseBody) *ModifyInterceptionTargetResponse {
	s.Body = v
	return s
}

func (s *ModifyInterceptionTargetResponse) Validate() error {
	return dara.Validate(s)
}
