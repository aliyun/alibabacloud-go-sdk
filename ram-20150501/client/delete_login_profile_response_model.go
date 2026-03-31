// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoginProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLoginProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLoginProfileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLoginProfileResponseBody) *DeleteLoginProfileResponse
	GetBody() *DeleteLoginProfileResponseBody
}

type DeleteLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLoginProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLoginProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLoginProfileResponse) GetBody() *DeleteLoginProfileResponseBody {
	return s.Body
}

func (s *DeleteLoginProfileResponse) SetHeaders(v map[string]*string) *DeleteLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoginProfileResponse) SetStatusCode(v int32) *DeleteLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoginProfileResponse) SetBody(v *DeleteLoginProfileResponseBody) *DeleteLoginProfileResponse {
	s.Body = v
	return s
}

func (s *DeleteLoginProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
