// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveSnapshotTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveSnapshotTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveSnapshotTemplatesResponseBody) *ListLiveSnapshotTemplatesResponse
	GetBody() *ListLiveSnapshotTemplatesResponseBody
}

type ListLiveSnapshotTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveSnapshotTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveSnapshotTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveSnapshotTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveSnapshotTemplatesResponse) GetBody() *ListLiveSnapshotTemplatesResponseBody {
	return s.Body
}

func (s *ListLiveSnapshotTemplatesResponse) SetHeaders(v map[string]*string) *ListLiveSnapshotTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveSnapshotTemplatesResponse) SetStatusCode(v int32) *ListLiveSnapshotTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponse) SetBody(v *ListLiveSnapshotTemplatesResponseBody) *ListLiveSnapshotTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListLiveSnapshotTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
