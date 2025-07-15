// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveStreamRecordIndexFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveStreamRecordIndexFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveStreamRecordIndexFilesResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveStreamRecordIndexFilesResponseBody) *CreateLiveStreamRecordIndexFilesResponse
	GetBody() *CreateLiveStreamRecordIndexFilesResponseBody
}

type CreateLiveStreamRecordIndexFilesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveStreamRecordIndexFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveStreamRecordIndexFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveStreamRecordIndexFilesResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamRecordIndexFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveStreamRecordIndexFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveStreamRecordIndexFilesResponse) GetBody() *CreateLiveStreamRecordIndexFilesResponseBody {
	return s.Body
}

func (s *CreateLiveStreamRecordIndexFilesResponse) SetHeaders(v map[string]*string) *CreateLiveStreamRecordIndexFilesResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponse) SetStatusCode(v int32) *CreateLiveStreamRecordIndexFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponse) SetBody(v *CreateLiveStreamRecordIndexFilesResponseBody) *CreateLiveStreamRecordIndexFilesResponse {
	s.Body = v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponse) Validate() error {
	return dara.Validate(s)
}
