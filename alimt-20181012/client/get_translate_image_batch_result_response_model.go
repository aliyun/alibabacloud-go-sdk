// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslateImageBatchResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranslateImageBatchResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranslateImageBatchResultResponse
	GetStatusCode() *int32
	SetBody(v *GetTranslateImageBatchResultResponseBody) *GetTranslateImageBatchResultResponse
	GetBody() *GetTranslateImageBatchResultResponseBody
}

type GetTranslateImageBatchResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranslateImageBatchResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranslateImageBatchResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranslateImageBatchResultResponse) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranslateImageBatchResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranslateImageBatchResultResponse) GetBody() *GetTranslateImageBatchResultResponseBody {
	return s.Body
}

func (s *GetTranslateImageBatchResultResponse) SetHeaders(v map[string]*string) *GetTranslateImageBatchResultResponse {
	s.Headers = v
	return s
}

func (s *GetTranslateImageBatchResultResponse) SetStatusCode(v int32) *GetTranslateImageBatchResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranslateImageBatchResultResponse) SetBody(v *GetTranslateImageBatchResultResponseBody) *GetTranslateImageBatchResultResponse {
	s.Body = v
	return s
}

func (s *GetTranslateImageBatchResultResponse) Validate() error {
	return dara.Validate(s)
}
