// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamRecordIndexFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamRecordIndexFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamRecordIndexFilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamRecordIndexFilesResponseBody) *DeleteLiveStreamRecordIndexFilesResponse
	GetBody() *DeleteLiveStreamRecordIndexFilesResponseBody
}

type DeleteLiveStreamRecordIndexFilesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamRecordIndexFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamRecordIndexFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamRecordIndexFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamRecordIndexFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamRecordIndexFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamRecordIndexFilesResponse) GetBody() *DeleteLiveStreamRecordIndexFilesResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamRecordIndexFilesResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamRecordIndexFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponse) SetStatusCode(v int32) *DeleteLiveStreamRecordIndexFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponse) SetBody(v *DeleteLiveStreamRecordIndexFilesResponseBody) *DeleteLiveStreamRecordIndexFilesResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamRecordIndexFilesResponse) Validate() error {
	return dara.Validate(s)
}
