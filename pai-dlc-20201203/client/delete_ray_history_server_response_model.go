// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRayHistoryServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRayHistoryServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRayHistoryServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRayHistoryServerResponseBody) *DeleteRayHistoryServerResponse
	GetBody() *DeleteRayHistoryServerResponseBody
}

type DeleteRayHistoryServerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRayHistoryServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRayHistoryServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRayHistoryServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteRayHistoryServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRayHistoryServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRayHistoryServerResponse) GetBody() *DeleteRayHistoryServerResponseBody {
	return s.Body
}

func (s *DeleteRayHistoryServerResponse) SetHeaders(v map[string]*string) *DeleteRayHistoryServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteRayHistoryServerResponse) SetStatusCode(v int32) *DeleteRayHistoryServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRayHistoryServerResponse) SetBody(v *DeleteRayHistoryServerResponseBody) *DeleteRayHistoryServerResponse {
	s.Body = v
	return s
}

func (s *DeleteRayHistoryServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
