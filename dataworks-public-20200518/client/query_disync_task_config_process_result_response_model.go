// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDISyncTaskConfigProcessResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDISyncTaskConfigProcessResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDISyncTaskConfigProcessResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryDISyncTaskConfigProcessResultResponseBody) *QueryDISyncTaskConfigProcessResultResponse
	GetBody() *QueryDISyncTaskConfigProcessResultResponseBody
}

type QueryDISyncTaskConfigProcessResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDISyncTaskConfigProcessResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDISyncTaskConfigProcessResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDISyncTaskConfigProcessResultResponse) GoString() string {
	return s.String()
}

func (s *QueryDISyncTaskConfigProcessResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDISyncTaskConfigProcessResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDISyncTaskConfigProcessResultResponse) GetBody() *QueryDISyncTaskConfigProcessResultResponseBody {
	return s.Body
}

func (s *QueryDISyncTaskConfigProcessResultResponse) SetHeaders(v map[string]*string) *QueryDISyncTaskConfigProcessResultResponse {
	s.Headers = v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponse) SetStatusCode(v int32) *QueryDISyncTaskConfigProcessResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponse) SetBody(v *QueryDISyncTaskConfigProcessResultResponseBody) *QueryDISyncTaskConfigProcessResultResponse {
	s.Body = v
	return s
}

func (s *QueryDISyncTaskConfigProcessResultResponse) Validate() error {
	return dara.Validate(s)
}
