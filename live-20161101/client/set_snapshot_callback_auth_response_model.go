// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSnapshotCallbackAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSnapshotCallbackAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSnapshotCallbackAuthResponse
	GetStatusCode() *int32
	SetBody(v *SetSnapshotCallbackAuthResponseBody) *SetSnapshotCallbackAuthResponse
	GetBody() *SetSnapshotCallbackAuthResponseBody
}

type SetSnapshotCallbackAuthResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSnapshotCallbackAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSnapshotCallbackAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSnapshotCallbackAuthResponse) GoString() string {
	return s.String()
}

func (s *SetSnapshotCallbackAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSnapshotCallbackAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSnapshotCallbackAuthResponse) GetBody() *SetSnapshotCallbackAuthResponseBody {
	return s.Body
}

func (s *SetSnapshotCallbackAuthResponse) SetHeaders(v map[string]*string) *SetSnapshotCallbackAuthResponse {
	s.Headers = v
	return s
}

func (s *SetSnapshotCallbackAuthResponse) SetStatusCode(v int32) *SetSnapshotCallbackAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSnapshotCallbackAuthResponse) SetBody(v *SetSnapshotCallbackAuthResponseBody) *SetSnapshotCallbackAuthResponse {
	s.Body = v
	return s
}

func (s *SetSnapshotCallbackAuthResponse) Validate() error {
	return dara.Validate(s)
}
