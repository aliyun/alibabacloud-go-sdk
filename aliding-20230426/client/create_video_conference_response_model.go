// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoConferenceResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoConferenceResponseBody) *CreateVideoConferenceResponse
	GetBody() *CreateVideoConferenceResponseBody
}

type CreateVideoConferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoConferenceResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoConferenceResponse) GetBody() *CreateVideoConferenceResponseBody {
	return s.Body
}

func (s *CreateVideoConferenceResponse) SetHeaders(v map[string]*string) *CreateVideoConferenceResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoConferenceResponse) SetStatusCode(v int32) *CreateVideoConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoConferenceResponse) SetBody(v *CreateVideoConferenceResponseBody) *CreateVideoConferenceResponse {
	s.Body = v
	return s
}

func (s *CreateVideoConferenceResponse) Validate() error {
	return dara.Validate(s)
}
