// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageGroupMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveMessageGroupMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveMessageGroupMessageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveMessageGroupMessageResponseBody) *DeleteLiveMessageGroupMessageResponse
	GetBody() *DeleteLiveMessageGroupMessageResponseBody
}

type DeleteLiveMessageGroupMessageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveMessageGroupMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveMessageGroupMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageGroupMessageResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageGroupMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveMessageGroupMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveMessageGroupMessageResponse) GetBody() *DeleteLiveMessageGroupMessageResponseBody {
	return s.Body
}

func (s *DeleteLiveMessageGroupMessageResponse) SetHeaders(v map[string]*string) *DeleteLiveMessageGroupMessageResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveMessageGroupMessageResponse) SetStatusCode(v int32) *DeleteLiveMessageGroupMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageResponse) SetBody(v *DeleteLiveMessageGroupMessageResponseBody) *DeleteLiveMessageGroupMessageResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveMessageGroupMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
