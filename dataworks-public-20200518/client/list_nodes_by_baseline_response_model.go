// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesByBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodesByBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodesByBaselineResponse
	GetStatusCode() *int32
	SetBody(v *ListNodesByBaselineResponseBody) *ListNodesByBaselineResponse
	GetBody() *ListNodesByBaselineResponseBody
}

type ListNodesByBaselineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesByBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesByBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByBaselineResponse) GoString() string {
	return s.String()
}

func (s *ListNodesByBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodesByBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodesByBaselineResponse) GetBody() *ListNodesByBaselineResponseBody {
	return s.Body
}

func (s *ListNodesByBaselineResponse) SetHeaders(v map[string]*string) *ListNodesByBaselineResponse {
	s.Headers = v
	return s
}

func (s *ListNodesByBaselineResponse) SetStatusCode(v int32) *ListNodesByBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesByBaselineResponse) SetBody(v *ListNodesByBaselineResponseBody) *ListNodesByBaselineResponse {
	s.Body = v
	return s
}

func (s *ListNodesByBaselineResponse) Validate() error {
	return dara.Validate(s)
}
