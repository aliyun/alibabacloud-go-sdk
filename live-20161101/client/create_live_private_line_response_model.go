// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePrivateLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLivePrivateLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLivePrivateLineResponse
	GetStatusCode() *int32
	SetBody(v *CreateLivePrivateLineResponseBody) *CreateLivePrivateLineResponse
	GetBody() *CreateLivePrivateLineResponseBody
}

type CreateLivePrivateLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivePrivateLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivePrivateLineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePrivateLineResponse) GoString() string {
	return s.String()
}

func (s *CreateLivePrivateLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLivePrivateLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLivePrivateLineResponse) GetBody() *CreateLivePrivateLineResponseBody {
	return s.Body
}

func (s *CreateLivePrivateLineResponse) SetHeaders(v map[string]*string) *CreateLivePrivateLineResponse {
	s.Headers = v
	return s
}

func (s *CreateLivePrivateLineResponse) SetStatusCode(v int32) *CreateLivePrivateLineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivePrivateLineResponse) SetBody(v *CreateLivePrivateLineResponseBody) *CreateLivePrivateLineResponse {
	s.Body = v
	return s
}

func (s *CreateLivePrivateLineResponse) Validate() error {
	return dara.Validate(s)
}
