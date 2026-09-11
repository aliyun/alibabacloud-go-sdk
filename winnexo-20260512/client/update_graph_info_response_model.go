// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGraphInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGraphInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGraphInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGraphInfoResponseBody) *UpdateGraphInfoResponse
	GetBody() *UpdateGraphInfoResponseBody
}

type UpdateGraphInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGraphInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGraphInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGraphInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateGraphInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGraphInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGraphInfoResponse) GetBody() *UpdateGraphInfoResponseBody {
	return s.Body
}

func (s *UpdateGraphInfoResponse) SetHeaders(v map[string]*string) *UpdateGraphInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateGraphInfoResponse) SetStatusCode(v int32) *UpdateGraphInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGraphInfoResponse) SetBody(v *UpdateGraphInfoResponseBody) *UpdateGraphInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateGraphInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
