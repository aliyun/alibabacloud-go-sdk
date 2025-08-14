// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveRecordFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveRecordFilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveRecordFilesResponseBody) *DeleteLiveRecordFilesResponse
	GetBody() *DeleteLiveRecordFilesResponseBody
}

type DeleteLiveRecordFilesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveRecordFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveRecordFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveRecordFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveRecordFilesResponse) GetBody() *DeleteLiveRecordFilesResponseBody {
	return s.Body
}

func (s *DeleteLiveRecordFilesResponse) SetHeaders(v map[string]*string) *DeleteLiveRecordFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRecordFilesResponse) SetStatusCode(v int32) *DeleteLiveRecordFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveRecordFilesResponse) SetBody(v *DeleteLiveRecordFilesResponseBody) *DeleteLiveRecordFilesResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveRecordFilesResponse) Validate() error {
	return dara.Validate(s)
}
