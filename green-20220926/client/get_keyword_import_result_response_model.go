// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeywordImportResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKeywordImportResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKeywordImportResultResponse
	GetStatusCode() *int32
	SetBody(v *GetKeywordImportResultResponseBody) *GetKeywordImportResultResponse
	GetBody() *GetKeywordImportResultResponseBody
}

type GetKeywordImportResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeywordImportResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeywordImportResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordImportResultResponse) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKeywordImportResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKeywordImportResultResponse) GetBody() *GetKeywordImportResultResponseBody {
	return s.Body
}

func (s *GetKeywordImportResultResponse) SetHeaders(v map[string]*string) *GetKeywordImportResultResponse {
	s.Headers = v
	return s
}

func (s *GetKeywordImportResultResponse) SetStatusCode(v int32) *GetKeywordImportResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeywordImportResultResponse) SetBody(v *GetKeywordImportResultResponseBody) *GetKeywordImportResultResponse {
	s.Body = v
	return s
}

func (s *GetKeywordImportResultResponse) Validate() error {
	return dara.Validate(s)
}
