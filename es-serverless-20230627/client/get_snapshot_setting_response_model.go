// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSnapshotSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSnapshotSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetSnapshotSettingResponseBody) *GetSnapshotSettingResponse
	GetBody() *GetSnapshotSettingResponseBody
}

type GetSnapshotSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSnapshotSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSnapshotSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotSettingResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSnapshotSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSnapshotSettingResponse) GetBody() *GetSnapshotSettingResponseBody {
	return s.Body
}

func (s *GetSnapshotSettingResponse) SetHeaders(v map[string]*string) *GetSnapshotSettingResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotSettingResponse) SetStatusCode(v int32) *GetSnapshotSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSnapshotSettingResponse) SetBody(v *GetSnapshotSettingResponseBody) *GetSnapshotSettingResponse {
	s.Body = v
	return s
}

func (s *GetSnapshotSettingResponse) Validate() error {
	return dara.Validate(s)
}
