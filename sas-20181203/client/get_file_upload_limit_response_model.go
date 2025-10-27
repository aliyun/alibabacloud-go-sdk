// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileUploadLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileUploadLimitResponse
	GetStatusCode() *int32
	SetBody(v *GetFileUploadLimitResponseBody) *GetFileUploadLimitResponse
	GetBody() *GetFileUploadLimitResponseBody
}

type GetFileUploadLimitResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileUploadLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileUploadLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadLimitResponse) GoString() string {
	return s.String()
}

func (s *GetFileUploadLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileUploadLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileUploadLimitResponse) GetBody() *GetFileUploadLimitResponseBody {
	return s.Body
}

func (s *GetFileUploadLimitResponse) SetHeaders(v map[string]*string) *GetFileUploadLimitResponse {
	s.Headers = v
	return s
}

func (s *GetFileUploadLimitResponse) SetStatusCode(v int32) *GetFileUploadLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileUploadLimitResponse) SetBody(v *GetFileUploadLimitResponseBody) *GetFileUploadLimitResponse {
	s.Body = v
	return s
}

func (s *GetFileUploadLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
