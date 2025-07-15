// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySnapshotCallbackAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySnapshotCallbackAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySnapshotCallbackAuthResponse
	GetStatusCode() *int32
	SetBody(v *QuerySnapshotCallbackAuthResponseBody) *QuerySnapshotCallbackAuthResponse
	GetBody() *QuerySnapshotCallbackAuthResponseBody
}

type QuerySnapshotCallbackAuthResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySnapshotCallbackAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySnapshotCallbackAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotCallbackAuthResponse) GoString() string {
	return s.String()
}

func (s *QuerySnapshotCallbackAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySnapshotCallbackAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySnapshotCallbackAuthResponse) GetBody() *QuerySnapshotCallbackAuthResponseBody {
	return s.Body
}

func (s *QuerySnapshotCallbackAuthResponse) SetHeaders(v map[string]*string) *QuerySnapshotCallbackAuthResponse {
	s.Headers = v
	return s
}

func (s *QuerySnapshotCallbackAuthResponse) SetStatusCode(v int32) *QuerySnapshotCallbackAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySnapshotCallbackAuthResponse) SetBody(v *QuerySnapshotCallbackAuthResponseBody) *QuerySnapshotCallbackAuthResponse {
	s.Body = v
	return s
}

func (s *QuerySnapshotCallbackAuthResponse) Validate() error {
	return dara.Validate(s)
}
