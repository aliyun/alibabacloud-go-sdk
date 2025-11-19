// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoInfoResponseBody) *UpdateVideoInfoResponse
	GetBody() *UpdateVideoInfoResponseBody
}

type UpdateVideoInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoInfoResponse) GetBody() *UpdateVideoInfoResponseBody {
	return s.Body
}

func (s *UpdateVideoInfoResponse) SetHeaders(v map[string]*string) *UpdateVideoInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoInfoResponse) SetStatusCode(v int32) *UpdateVideoInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoInfoResponse) SetBody(v *UpdateVideoInfoResponseBody) *UpdateVideoInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
