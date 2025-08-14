// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMediaLiveInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMediaLiveInputResponse
	GetStatusCode() *int32
	SetBody(v *CreateMediaLiveInputResponseBody) *CreateMediaLiveInputResponse
	GetBody() *CreateMediaLiveInputResponseBody
}

type CreateMediaLiveInputResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMediaLiveInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMediaLiveInputResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMediaLiveInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMediaLiveInputResponse) GetBody() *CreateMediaLiveInputResponseBody {
	return s.Body
}

func (s *CreateMediaLiveInputResponse) SetHeaders(v map[string]*string) *CreateMediaLiveInputResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaLiveInputResponse) SetStatusCode(v int32) *CreateMediaLiveInputResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaLiveInputResponse) SetBody(v *CreateMediaLiveInputResponseBody) *CreateMediaLiveInputResponse {
	s.Body = v
	return s
}

func (s *CreateMediaLiveInputResponse) Validate() error {
	return dara.Validate(s)
}
