// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableUnderstandingResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableUnderstandingResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableUnderstandingResultResponse
	GetStatusCode() *int32
	SetBody(v *GetTableUnderstandingResultResponseBody) *GetTableUnderstandingResultResponse
	GetBody() *GetTableUnderstandingResultResponseBody
}

type GetTableUnderstandingResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableUnderstandingResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableUnderstandingResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableUnderstandingResultResponse) GoString() string {
	return s.String()
}

func (s *GetTableUnderstandingResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableUnderstandingResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableUnderstandingResultResponse) GetBody() *GetTableUnderstandingResultResponseBody {
	return s.Body
}

func (s *GetTableUnderstandingResultResponse) SetHeaders(v map[string]*string) *GetTableUnderstandingResultResponse {
	s.Headers = v
	return s
}

func (s *GetTableUnderstandingResultResponse) SetStatusCode(v int32) *GetTableUnderstandingResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableUnderstandingResultResponse) SetBody(v *GetTableUnderstandingResultResponseBody) *GetTableUnderstandingResultResponse {
	s.Body = v
	return s
}

func (s *GetTableUnderstandingResultResponse) Validate() error {
	return dara.Validate(s)
}
