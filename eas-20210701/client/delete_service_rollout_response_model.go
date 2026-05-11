// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRolloutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceRolloutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceRolloutResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceRolloutResponseBody) *DeleteServiceRolloutResponse
	GetBody() *DeleteServiceRolloutResponseBody
}

type DeleteServiceRolloutResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceRolloutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceRolloutResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRolloutResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceRolloutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceRolloutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceRolloutResponse) GetBody() *DeleteServiceRolloutResponseBody {
	return s.Body
}

func (s *DeleteServiceRolloutResponse) SetHeaders(v map[string]*string) *DeleteServiceRolloutResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceRolloutResponse) SetStatusCode(v int32) *DeleteServiceRolloutResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceRolloutResponse) SetBody(v *DeleteServiceRolloutResponseBody) *DeleteServiceRolloutResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceRolloutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
