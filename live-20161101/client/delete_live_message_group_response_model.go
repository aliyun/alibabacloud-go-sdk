// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveMessageGroupResponseBody) *DeleteLiveMessageGroupResponse
	GetBody() *DeleteLiveMessageGroupResponseBody
}

type DeleteLiveMessageGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveMessageGroupResponse) GetBody() *DeleteLiveMessageGroupResponseBody {
	return s.Body
}

func (s *DeleteLiveMessageGroupResponse) SetHeaders(v map[string]*string) *DeleteLiveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveMessageGroupResponse) SetStatusCode(v int32) *DeleteLiveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveMessageGroupResponse) SetBody(v *DeleteLiveMessageGroupResponseBody) *DeleteLiveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
