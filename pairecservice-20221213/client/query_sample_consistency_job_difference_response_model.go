// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySampleConsistencyJobDifferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySampleConsistencyJobDifferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySampleConsistencyJobDifferenceResponse
	GetStatusCode() *int32
	SetBody(v *QuerySampleConsistencyJobDifferenceResponseBody) *QuerySampleConsistencyJobDifferenceResponse
	GetBody() *QuerySampleConsistencyJobDifferenceResponseBody
}

type QuerySampleConsistencyJobDifferenceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySampleConsistencyJobDifferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySampleConsistencyJobDifferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySampleConsistencyJobDifferenceResponse) GoString() string {
	return s.String()
}

func (s *QuerySampleConsistencyJobDifferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySampleConsistencyJobDifferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySampleConsistencyJobDifferenceResponse) GetBody() *QuerySampleConsistencyJobDifferenceResponseBody {
	return s.Body
}

func (s *QuerySampleConsistencyJobDifferenceResponse) SetHeaders(v map[string]*string) *QuerySampleConsistencyJobDifferenceResponse {
	s.Headers = v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponse) SetStatusCode(v int32) *QuerySampleConsistencyJobDifferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponse) SetBody(v *QuerySampleConsistencyJobDifferenceResponseBody) *QuerySampleConsistencyJobDifferenceResponse {
	s.Body = v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceResponse) Validate() error {
	return dara.Validate(s)
}
