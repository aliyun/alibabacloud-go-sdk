// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreArchivedRecordingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreArchivedRecordingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreArchivedRecordingsResponse
	GetStatusCode() *int32
	SetBody(v *RestoreArchivedRecordingsResponseBody) *RestoreArchivedRecordingsResponse
	GetBody() *RestoreArchivedRecordingsResponseBody
}

type RestoreArchivedRecordingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreArchivedRecordingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreArchivedRecordingsResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreArchivedRecordingsResponse) GoString() string {
	return s.String()
}

func (s *RestoreArchivedRecordingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreArchivedRecordingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreArchivedRecordingsResponse) GetBody() *RestoreArchivedRecordingsResponseBody {
	return s.Body
}

func (s *RestoreArchivedRecordingsResponse) SetHeaders(v map[string]*string) *RestoreArchivedRecordingsResponse {
	s.Headers = v
	return s
}

func (s *RestoreArchivedRecordingsResponse) SetStatusCode(v int32) *RestoreArchivedRecordingsResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreArchivedRecordingsResponse) SetBody(v *RestoreArchivedRecordingsResponseBody) *RestoreArchivedRecordingsResponse {
	s.Body = v
	return s
}

func (s *RestoreArchivedRecordingsResponse) Validate() error {
	return dara.Validate(s)
}
