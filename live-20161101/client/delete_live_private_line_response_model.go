// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePrivateLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivePrivateLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivePrivateLineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivePrivateLineResponseBody) *DeleteLivePrivateLineResponse
	GetBody() *DeleteLivePrivateLineResponseBody
}

type DeleteLivePrivateLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivePrivateLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivePrivateLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePrivateLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePrivateLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivePrivateLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivePrivateLineResponse) GetBody() *DeleteLivePrivateLineResponseBody {
	return s.Body
}

func (s *DeleteLivePrivateLineResponse) SetHeaders(v map[string]*string) *DeleteLivePrivateLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivePrivateLineResponse) SetStatusCode(v int32) *DeleteLivePrivateLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivePrivateLineResponse) SetBody(v *DeleteLivePrivateLineResponseBody) *DeleteLivePrivateLineResponse {
	s.Body = v
	return s
}

func (s *DeleteLivePrivateLineResponse) Validate() error {
	return dara.Validate(s)
}
