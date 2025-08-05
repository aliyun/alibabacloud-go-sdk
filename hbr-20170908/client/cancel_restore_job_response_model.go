// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRestoreJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRestoreJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRestoreJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelRestoreJobResponseBody) *CancelRestoreJobResponse
	GetBody() *CancelRestoreJobResponseBody
}

type CancelRestoreJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRestoreJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRestoreJobResponse) GoString() string {
	return s.String()
}

func (s *CancelRestoreJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRestoreJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRestoreJobResponse) GetBody() *CancelRestoreJobResponseBody {
	return s.Body
}

func (s *CancelRestoreJobResponse) SetHeaders(v map[string]*string) *CancelRestoreJobResponse {
	s.Headers = v
	return s
}

func (s *CancelRestoreJobResponse) SetStatusCode(v int32) *CancelRestoreJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRestoreJobResponse) SetBody(v *CancelRestoreJobResponseBody) *CancelRestoreJobResponse {
	s.Body = v
	return s
}

func (s *CancelRestoreJobResponse) Validate() error {
	return dara.Validate(s)
}
