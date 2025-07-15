// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePullToPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivePullToPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivePullToPushResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivePullToPushResponseBody) *DeleteLivePullToPushResponse
	GetBody() *DeleteLivePullToPushResponseBody
}

type DeleteLivePullToPushResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivePullToPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivePullToPushResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePullToPushResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePullToPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivePullToPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivePullToPushResponse) GetBody() *DeleteLivePullToPushResponseBody {
	return s.Body
}

func (s *DeleteLivePullToPushResponse) SetHeaders(v map[string]*string) *DeleteLivePullToPushResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivePullToPushResponse) SetStatusCode(v int32) *DeleteLivePullToPushResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivePullToPushResponse) SetBody(v *DeleteLivePullToPushResponseBody) *DeleteLivePullToPushResponse {
	s.Body = v
	return s
}

func (s *DeleteLivePullToPushResponse) Validate() error {
	return dara.Validate(s)
}
