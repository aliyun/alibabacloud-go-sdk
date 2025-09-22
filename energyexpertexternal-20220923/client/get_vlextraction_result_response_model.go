// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVLExtractionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVLExtractionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVLExtractionResultResponse
	GetStatusCode() *int32
	SetBody(v *GetVLExtractionResultResponseBody) *GetVLExtractionResultResponse
	GetBody() *GetVLExtractionResultResponseBody
}

type GetVLExtractionResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVLExtractionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVLExtractionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVLExtractionResultResponse) GoString() string {
	return s.String()
}

func (s *GetVLExtractionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVLExtractionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVLExtractionResultResponse) GetBody() *GetVLExtractionResultResponseBody {
	return s.Body
}

func (s *GetVLExtractionResultResponse) SetHeaders(v map[string]*string) *GetVLExtractionResultResponse {
	s.Headers = v
	return s
}

func (s *GetVLExtractionResultResponse) SetStatusCode(v int32) *GetVLExtractionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVLExtractionResultResponse) SetBody(v *GetVLExtractionResultResponseBody) *GetVLExtractionResultResponse {
	s.Body = v
	return s
}

func (s *GetVLExtractionResultResponse) Validate() error {
	return dara.Validate(s)
}
