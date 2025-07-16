// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveResponseBody) *DeleteLiveResponse
	GetBody() *DeleteLiveResponseBody
}

type DeleteLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveResponse) GetBody() *DeleteLiveResponseBody {
	return s.Body
}

func (s *DeleteLiveResponse) SetHeaders(v map[string]*string) *DeleteLiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveResponse) SetStatusCode(v int32) *DeleteLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveResponse) SetBody(v *DeleteLiveResponseBody) *DeleteLiveResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveResponse) Validate() error {
	return dara.Validate(s)
}
