// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySavingsPlansInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySavingsPlansInstanceResponse
	GetStatusCode() *int32
	SetBody(v *QuerySavingsPlansInstanceResponseBody) *QuerySavingsPlansInstanceResponse
	GetBody() *QuerySavingsPlansInstanceResponseBody
}

type QuerySavingsPlansInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySavingsPlansInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySavingsPlansInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponse) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySavingsPlansInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySavingsPlansInstanceResponse) GetBody() *QuerySavingsPlansInstanceResponseBody {
	return s.Body
}

func (s *QuerySavingsPlansInstanceResponse) SetHeaders(v map[string]*string) *QuerySavingsPlansInstanceResponse {
	s.Headers = v
	return s
}

func (s *QuerySavingsPlansInstanceResponse) SetStatusCode(v int32) *QuerySavingsPlansInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponse) SetBody(v *QuerySavingsPlansInstanceResponseBody) *QuerySavingsPlansInstanceResponse {
	s.Body = v
	return s
}

func (s *QuerySavingsPlansInstanceResponse) Validate() error {
	return dara.Validate(s)
}
