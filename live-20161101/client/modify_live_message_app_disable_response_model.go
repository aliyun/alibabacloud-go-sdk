// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppDisableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveMessageAppDisableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveMessageAppDisableResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveMessageAppDisableResponseBody) *ModifyLiveMessageAppDisableResponse
	GetBody() *ModifyLiveMessageAppDisableResponseBody
}

type ModifyLiveMessageAppDisableResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveMessageAppDisableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveMessageAppDisableResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppDisableResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppDisableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveMessageAppDisableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveMessageAppDisableResponse) GetBody() *ModifyLiveMessageAppDisableResponseBody {
	return s.Body
}

func (s *ModifyLiveMessageAppDisableResponse) SetHeaders(v map[string]*string) *ModifyLiveMessageAppDisableResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveMessageAppDisableResponse) SetStatusCode(v int32) *ModifyLiveMessageAppDisableResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveMessageAppDisableResponse) SetBody(v *ModifyLiveMessageAppDisableResponseBody) *ModifyLiveMessageAppDisableResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveMessageAppDisableResponse) Validate() error {
	return dara.Validate(s)
}
