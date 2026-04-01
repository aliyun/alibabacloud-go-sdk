// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFalconSeekResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFalconSeekResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFalconSeekResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFalconSeekResponseBody) *UpdateFalconSeekResponse
	GetBody() *UpdateFalconSeekResponseBody
}

type UpdateFalconSeekResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFalconSeekResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFalconSeekResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFalconSeekResponse) GoString() string {
	return s.String()
}

func (s *UpdateFalconSeekResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFalconSeekResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFalconSeekResponse) GetBody() *UpdateFalconSeekResponseBody {
	return s.Body
}

func (s *UpdateFalconSeekResponse) SetHeaders(v map[string]*string) *UpdateFalconSeekResponse {
	s.Headers = v
	return s
}

func (s *UpdateFalconSeekResponse) SetStatusCode(v int32) *UpdateFalconSeekResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFalconSeekResponse) SetBody(v *UpdateFalconSeekResponseBody) *UpdateFalconSeekResponse {
	s.Body = v
	return s
}

func (s *UpdateFalconSeekResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
