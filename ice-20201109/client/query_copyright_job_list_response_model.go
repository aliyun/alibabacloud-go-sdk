// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCopyrightJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCopyrightJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryCopyrightJobListResponseBody) *QueryCopyrightJobListResponse
	GetBody() *QueryCopyrightJobListResponseBody
}

type QueryCopyrightJobListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCopyrightJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCopyrightJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCopyrightJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCopyrightJobListResponse) GetBody() *QueryCopyrightJobListResponseBody {
	return s.Body
}

func (s *QueryCopyrightJobListResponse) SetHeaders(v map[string]*string) *QueryCopyrightJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryCopyrightJobListResponse) SetStatusCode(v int32) *QueryCopyrightJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightJobListResponse) SetBody(v *QueryCopyrightJobListResponseBody) *QueryCopyrightJobListResponse {
	s.Body = v
	return s
}

func (s *QueryCopyrightJobListResponse) Validate() error {
	return dara.Validate(s)
}
