// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAlertHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAlertHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *SearchAlertHistoriesResponseBody) *SearchAlertHistoriesResponse
	GetBody() *SearchAlertHistoriesResponseBody
}

type SearchAlertHistoriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAlertHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAlertHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertHistoriesResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAlertHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAlertHistoriesResponse) GetBody() *SearchAlertHistoriesResponseBody {
	return s.Body
}

func (s *SearchAlertHistoriesResponse) SetHeaders(v map[string]*string) *SearchAlertHistoriesResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertHistoriesResponse) SetStatusCode(v int32) *SearchAlertHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertHistoriesResponse) SetBody(v *SearchAlertHistoriesResponseBody) *SearchAlertHistoriesResponse {
	s.Body = v
	return s
}

func (s *SearchAlertHistoriesResponse) Validate() error {
	return dara.Validate(s)
}
