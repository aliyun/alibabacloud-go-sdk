// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveDelayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveDelayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveDelayConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveDelayConfigResponseBody) *CreateLiveDelayConfigResponse
	GetBody() *CreateLiveDelayConfigResponseBody
}

type CreateLiveDelayConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveDelayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveDelayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveDelayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveDelayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveDelayConfigResponse) GetBody() *CreateLiveDelayConfigResponseBody {
	return s.Body
}

func (s *CreateLiveDelayConfigResponse) SetHeaders(v map[string]*string) *CreateLiveDelayConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveDelayConfigResponse) SetStatusCode(v int32) *CreateLiveDelayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveDelayConfigResponse) SetBody(v *CreateLiveDelayConfigResponseBody) *CreateLiveDelayConfigResponse {
	s.Body = v
	return s
}

func (s *CreateLiveDelayConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
