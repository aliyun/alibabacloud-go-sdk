// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetFileMetaResponseBody) *GetDatasetFileMetaResponse
	GetBody() *GetDatasetFileMetaResponseBody
}

type GetDatasetFileMetaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetFileMetaResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetFileMetaResponse) GetBody() *GetDatasetFileMetaResponseBody {
	return s.Body
}

func (s *GetDatasetFileMetaResponse) SetHeaders(v map[string]*string) *GetDatasetFileMetaResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetFileMetaResponse) SetStatusCode(v int32) *GetDatasetFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetFileMetaResponse) SetBody(v *GetDatasetFileMetaResponseBody) *GetDatasetFileMetaResponse {
	s.Body = v
	return s
}

func (s *GetDatasetFileMetaResponse) Validate() error {
	return dara.Validate(s)
}
