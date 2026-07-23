// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventHouseRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventHouseRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventHouseRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventHouseRuntimeResponseBody) *DeleteEventHouseRuntimeResponse
	GetBody() *DeleteEventHouseRuntimeResponseBody
}

type DeleteEventHouseRuntimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventHouseRuntimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventHouseRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventHouseRuntimeResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventHouseRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventHouseRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventHouseRuntimeResponse) GetBody() *DeleteEventHouseRuntimeResponseBody {
	return s.Body
}

func (s *DeleteEventHouseRuntimeResponse) SetHeaders(v map[string]*string) *DeleteEventHouseRuntimeResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventHouseRuntimeResponse) SetStatusCode(v int32) *DeleteEventHouseRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventHouseRuntimeResponse) SetBody(v *DeleteEventHouseRuntimeResponseBody) *DeleteEventHouseRuntimeResponse {
	s.Body = v
	return s
}

func (s *DeleteEventHouseRuntimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
