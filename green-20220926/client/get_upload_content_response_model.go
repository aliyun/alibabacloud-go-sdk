// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadContentResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadContentResponseBody) *GetUploadContentResponse
	GetBody() *GetUploadContentResponseBody
}

type GetUploadContentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadContentResponse) GoString() string {
	return s.String()
}

func (s *GetUploadContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadContentResponse) GetBody() *GetUploadContentResponseBody {
	return s.Body
}

func (s *GetUploadContentResponse) SetHeaders(v map[string]*string) *GetUploadContentResponse {
	s.Headers = v
	return s
}

func (s *GetUploadContentResponse) SetStatusCode(v int32) *GetUploadContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadContentResponse) SetBody(v *GetUploadContentResponseBody) *GetUploadContentResponse {
	s.Body = v
	return s
}

func (s *GetUploadContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
