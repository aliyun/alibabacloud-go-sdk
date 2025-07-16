// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveResponseBody) *UpdateLiveResponse
	GetBody() *UpdateLiveResponseBody
}

type UpdateLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveResponse) GetBody() *UpdateLiveResponseBody {
	return s.Body
}

func (s *UpdateLiveResponse) SetHeaders(v map[string]*string) *UpdateLiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveResponse) SetStatusCode(v int32) *UpdateLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveResponse) SetBody(v *UpdateLiveResponseBody) *UpdateLiveResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveResponse) Validate() error {
	return dara.Validate(s)
}
