// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishGraySwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublishGraySwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublishGraySwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublishGraySwitchResponseBody) *UpdatePublishGraySwitchResponse
	GetBody() *UpdatePublishGraySwitchResponseBody
}

type UpdatePublishGraySwitchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublishGraySwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublishGraySwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishGraySwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublishGraySwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublishGraySwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublishGraySwitchResponse) GetBody() *UpdatePublishGraySwitchResponseBody {
	return s.Body
}

func (s *UpdatePublishGraySwitchResponse) SetHeaders(v map[string]*string) *UpdatePublishGraySwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublishGraySwitchResponse) SetStatusCode(v int32) *UpdatePublishGraySwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublishGraySwitchResponse) SetBody(v *UpdatePublishGraySwitchResponseBody) *UpdatePublishGraySwitchResponse {
	s.Body = v
	return s
}

func (s *UpdatePublishGraySwitchResponse) Validate() error {
	return dara.Validate(s)
}
