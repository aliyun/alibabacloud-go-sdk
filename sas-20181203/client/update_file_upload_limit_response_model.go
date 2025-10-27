// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileUploadLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileUploadLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileUploadLimitResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileUploadLimitResponseBody) *UpdateFileUploadLimitResponse
	GetBody() *UpdateFileUploadLimitResponseBody
}

type UpdateFileUploadLimitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileUploadLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileUploadLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileUploadLimitResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileUploadLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileUploadLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileUploadLimitResponse) GetBody() *UpdateFileUploadLimitResponseBody {
	return s.Body
}

func (s *UpdateFileUploadLimitResponse) SetHeaders(v map[string]*string) *UpdateFileUploadLimitResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileUploadLimitResponse) SetStatusCode(v int32) *UpdateFileUploadLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileUploadLimitResponse) SetBody(v *UpdateFileUploadLimitResponseBody) *UpdateFileUploadLimitResponse {
	s.Body = v
	return s
}

func (s *UpdateFileUploadLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
