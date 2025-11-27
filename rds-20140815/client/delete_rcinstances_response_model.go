// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRCInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRCInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRCInstancesResponseBody) *DeleteRCInstancesResponse
	GetBody() *DeleteRCInstancesResponseBody
}

type DeleteRCInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRCInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRCInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRCInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRCInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRCInstancesResponse) GetBody() *DeleteRCInstancesResponseBody {
	return s.Body
}

func (s *DeleteRCInstancesResponse) SetHeaders(v map[string]*string) *DeleteRCInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRCInstancesResponse) SetStatusCode(v int32) *DeleteRCInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRCInstancesResponse) SetBody(v *DeleteRCInstancesResponseBody) *DeleteRCInstancesResponse {
	s.Body = v
	return s
}

func (s *DeleteRCInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
