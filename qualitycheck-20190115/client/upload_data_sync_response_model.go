// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDataSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDataSyncResponse
	GetStatusCode() *int32
	SetBody(v *UploadDataSyncResponseBody) *UploadDataSyncResponse
	GetBody() *UploadDataSyncResponseBody
}

type UploadDataSyncResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncResponse) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDataSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDataSyncResponse) GetBody() *UploadDataSyncResponseBody {
	return s.Body
}

func (s *UploadDataSyncResponse) SetHeaders(v map[string]*string) *UploadDataSyncResponse {
	s.Headers = v
	return s
}

func (s *UploadDataSyncResponse) SetStatusCode(v int32) *UploadDataSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataSyncResponse) SetBody(v *UploadDataSyncResponseBody) *UploadDataSyncResponse {
	s.Body = v
	return s
}

func (s *UploadDataSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
