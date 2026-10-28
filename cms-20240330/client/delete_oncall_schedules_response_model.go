// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOncallSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOncallSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOncallSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOncallSchedulesResponseBody) *DeleteOncallSchedulesResponse
	GetBody() *DeleteOncallSchedulesResponseBody
}

type DeleteOncallSchedulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOncallSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOncallSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOncallSchedulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteOncallSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOncallSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOncallSchedulesResponse) GetBody() *DeleteOncallSchedulesResponseBody {
	return s.Body
}

func (s *DeleteOncallSchedulesResponse) SetHeaders(v map[string]*string) *DeleteOncallSchedulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteOncallSchedulesResponse) SetStatusCode(v int32) *DeleteOncallSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOncallSchedulesResponse) SetBody(v *DeleteOncallSchedulesResponseBody) *DeleteOncallSchedulesResponse {
	s.Body = v
	return s
}

func (s *DeleteOncallSchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
