// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetSmartqStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDatasetSmartqStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDatasetSmartqStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryDatasetSmartqStatusResponseBody) *QueryDatasetSmartqStatusResponse
	GetBody() *QueryDatasetSmartqStatusResponseBody
}

type QueryDatasetSmartqStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetSmartqStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDatasetSmartqStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetSmartqStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetSmartqStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDatasetSmartqStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDatasetSmartqStatusResponse) GetBody() *QueryDatasetSmartqStatusResponseBody {
	return s.Body
}

func (s *QueryDatasetSmartqStatusResponse) SetHeaders(v map[string]*string) *QueryDatasetSmartqStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetSmartqStatusResponse) SetStatusCode(v int32) *QueryDatasetSmartqStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetSmartqStatusResponse) SetBody(v *QueryDatasetSmartqStatusResponseBody) *QueryDatasetSmartqStatusResponse {
	s.Body = v
	return s
}

func (s *QueryDatasetSmartqStatusResponse) Validate() error {
	return dara.Validate(s)
}
