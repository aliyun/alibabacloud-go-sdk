// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotImportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFpShotImportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFpShotImportJobResponse
	GetStatusCode() *int32
	SetBody(v *ListFpShotImportJobResponseBody) *ListFpShotImportJobResponse
	GetBody() *ListFpShotImportJobResponseBody
}

type ListFpShotImportJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFpShotImportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFpShotImportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotImportJobResponse) GoString() string {
	return s.String()
}

func (s *ListFpShotImportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFpShotImportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFpShotImportJobResponse) GetBody() *ListFpShotImportJobResponseBody {
	return s.Body
}

func (s *ListFpShotImportJobResponse) SetHeaders(v map[string]*string) *ListFpShotImportJobResponse {
	s.Headers = v
	return s
}

func (s *ListFpShotImportJobResponse) SetStatusCode(v int32) *ListFpShotImportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFpShotImportJobResponse) SetBody(v *ListFpShotImportJobResponseBody) *ListFpShotImportJobResponse {
	s.Body = v
	return s
}

func (s *ListFpShotImportJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
