// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertHdrVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertHdrVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertHdrVideoResponse
	GetStatusCode() *int32
	SetBody(v *ConvertHdrVideoResponseBody) *ConvertHdrVideoResponse
	GetBody() *ConvertHdrVideoResponseBody
}

type ConvertHdrVideoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertHdrVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertHdrVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertHdrVideoResponse) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertHdrVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertHdrVideoResponse) GetBody() *ConvertHdrVideoResponseBody {
	return s.Body
}

func (s *ConvertHdrVideoResponse) SetHeaders(v map[string]*string) *ConvertHdrVideoResponse {
	s.Headers = v
	return s
}

func (s *ConvertHdrVideoResponse) SetStatusCode(v int32) *ConvertHdrVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertHdrVideoResponse) SetBody(v *ConvertHdrVideoResponseBody) *ConvertHdrVideoResponse {
	s.Body = v
	return s
}

func (s *ConvertHdrVideoResponse) Validate() error {
	return dara.Validate(s)
}
