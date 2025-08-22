// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadStagingRoutineCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadStagingRoutineCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadStagingRoutineCodeResponse
	GetStatusCode() *int32
	SetBody(v *UploadStagingRoutineCodeResponseBody) *UploadStagingRoutineCodeResponse
	GetBody() *UploadStagingRoutineCodeResponseBody
}

type UploadStagingRoutineCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadStagingRoutineCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadStagingRoutineCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadStagingRoutineCodeResponse) GoString() string {
	return s.String()
}

func (s *UploadStagingRoutineCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadStagingRoutineCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadStagingRoutineCodeResponse) GetBody() *UploadStagingRoutineCodeResponseBody {
	return s.Body
}

func (s *UploadStagingRoutineCodeResponse) SetHeaders(v map[string]*string) *UploadStagingRoutineCodeResponse {
	s.Headers = v
	return s
}

func (s *UploadStagingRoutineCodeResponse) SetStatusCode(v int32) *UploadStagingRoutineCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadStagingRoutineCodeResponse) SetBody(v *UploadStagingRoutineCodeResponseBody) *UploadStagingRoutineCodeResponse {
	s.Body = v
	return s
}

func (s *UploadStagingRoutineCodeResponse) Validate() error {
	return dara.Validate(s)
}
