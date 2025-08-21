// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSnapshotSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSnapshotSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSnapshotSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSnapshotSettingResponseBody) *UpdateSnapshotSettingResponse
	GetBody() *UpdateSnapshotSettingResponseBody
}

type UpdateSnapshotSettingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSnapshotSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSnapshotSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSnapshotSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSnapshotSettingResponse) GetBody() *UpdateSnapshotSettingResponseBody {
	return s.Body
}

func (s *UpdateSnapshotSettingResponse) SetHeaders(v map[string]*string) *UpdateSnapshotSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotSettingResponse) SetStatusCode(v int32) *UpdateSnapshotSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSnapshotSettingResponse) SetBody(v *UpdateSnapshotSettingResponseBody) *UpdateSnapshotSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateSnapshotSettingResponse) Validate() error {
	return dara.Validate(s)
}
