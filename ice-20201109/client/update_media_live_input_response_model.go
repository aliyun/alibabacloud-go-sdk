// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaLiveInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaLiveInputResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaLiveInputResponseBody) *UpdateMediaLiveInputResponse
	GetBody() *UpdateMediaLiveInputResponseBody
}

type UpdateMediaLiveInputResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaLiveInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaLiveInputResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaLiveInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaLiveInputResponse) GetBody() *UpdateMediaLiveInputResponseBody {
	return s.Body
}

func (s *UpdateMediaLiveInputResponse) SetHeaders(v map[string]*string) *UpdateMediaLiveInputResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaLiveInputResponse) SetStatusCode(v int32) *UpdateMediaLiveInputResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaLiveInputResponse) SetBody(v *UpdateMediaLiveInputResponseBody) *UpdateMediaLiveInputResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaLiveInputResponse) Validate() error {
	return dara.Validate(s)
}
