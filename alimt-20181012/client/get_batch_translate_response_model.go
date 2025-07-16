// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTranslateResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTranslateResponseBody) *GetBatchTranslateResponse
	GetBody() *GetBatchTranslateResponseBody
}

type GetBatchTranslateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTranslateResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTranslateResponse) GetBody() *GetBatchTranslateResponseBody {
	return s.Body
}

func (s *GetBatchTranslateResponse) SetHeaders(v map[string]*string) *GetBatchTranslateResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTranslateResponse) SetStatusCode(v int32) *GetBatchTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTranslateResponse) SetBody(v *GetBatchTranslateResponseBody) *GetBatchTranslateResponse {
	s.Body = v
	return s
}

func (s *GetBatchTranslateResponse) Validate() error {
	return dara.Validate(s)
}
