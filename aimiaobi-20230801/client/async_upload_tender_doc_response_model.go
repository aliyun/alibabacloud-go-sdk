// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadTenderDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncUploadTenderDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncUploadTenderDocResponse
	GetStatusCode() *int32
	SetBody(v *AsyncUploadTenderDocResponseBody) *AsyncUploadTenderDocResponse
	GetBody() *AsyncUploadTenderDocResponseBody
}

type AsyncUploadTenderDocResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncUploadTenderDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncUploadTenderDocResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadTenderDocResponse) GoString() string {
	return s.String()
}

func (s *AsyncUploadTenderDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncUploadTenderDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncUploadTenderDocResponse) GetBody() *AsyncUploadTenderDocResponseBody {
	return s.Body
}

func (s *AsyncUploadTenderDocResponse) SetHeaders(v map[string]*string) *AsyncUploadTenderDocResponse {
	s.Headers = v
	return s
}

func (s *AsyncUploadTenderDocResponse) SetStatusCode(v int32) *AsyncUploadTenderDocResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncUploadTenderDocResponse) SetBody(v *AsyncUploadTenderDocResponseBody) *AsyncUploadTenderDocResponse {
	s.Body = v
	return s
}

func (s *AsyncUploadTenderDocResponse) Validate() error {
	return dara.Validate(s)
}
