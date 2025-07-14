// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDatasetInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDatasetInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryDatasetInfoResponseBody) *QueryDatasetInfoResponse
	GetBody() *QueryDatasetInfoResponseBody
}

type QueryDatasetInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDatasetInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDatasetInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDatasetInfoResponse) GetBody() *QueryDatasetInfoResponseBody {
	return s.Body
}

func (s *QueryDatasetInfoResponse) SetHeaders(v map[string]*string) *QueryDatasetInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetInfoResponse) SetStatusCode(v int32) *QueryDatasetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetInfoResponse) SetBody(v *QueryDatasetInfoResponseBody) *QueryDatasetInfoResponse {
	s.Body = v
	return s
}

func (s *QueryDatasetInfoResponse) Validate() error {
	return dara.Validate(s)
}
