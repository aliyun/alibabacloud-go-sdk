// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetFileMetaResponseBody) *GetFileMetaResponse
	GetBody() *GetFileMetaResponseBody
}

type GetFileMetaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileMetaResponse) GoString() string {
	return s.String()
}

func (s *GetFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileMetaResponse) GetBody() *GetFileMetaResponseBody {
	return s.Body
}

func (s *GetFileMetaResponse) SetHeaders(v map[string]*string) *GetFileMetaResponse {
	s.Headers = v
	return s
}

func (s *GetFileMetaResponse) SetStatusCode(v int32) *GetFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileMetaResponse) SetBody(v *GetFileMetaResponseBody) *GetFileMetaResponse {
	s.Body = v
	return s
}

func (s *GetFileMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
