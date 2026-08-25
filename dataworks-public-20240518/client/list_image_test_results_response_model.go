// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTestResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageTestResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageTestResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListImageTestResultsResponseBody) *ListImageTestResultsResponse
	GetBody() *ListImageTestResultsResponseBody
}

type ListImageTestResultsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageTestResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageTestResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageTestResultsResponse) GoString() string {
	return s.String()
}

func (s *ListImageTestResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageTestResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageTestResultsResponse) GetBody() *ListImageTestResultsResponseBody {
	return s.Body
}

func (s *ListImageTestResultsResponse) SetHeaders(v map[string]*string) *ListImageTestResultsResponse {
	s.Headers = v
	return s
}

func (s *ListImageTestResultsResponse) SetStatusCode(v int32) *ListImageTestResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageTestResultsResponse) SetBody(v *ListImageTestResultsResponseBody) *ListImageTestResultsResponse {
	s.Body = v
	return s
}

func (s *ListImageTestResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
