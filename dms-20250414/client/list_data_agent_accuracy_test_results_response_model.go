// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentAccuracyTestResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentAccuracyTestResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentAccuracyTestResultsResponseBody) *ListDataAgentAccuracyTestResultsResponse
	GetBody() *ListDataAgentAccuracyTestResultsResponseBody
}

type ListDataAgentAccuracyTestResultsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentAccuracyTestResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentAccuracyTestResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentAccuracyTestResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentAccuracyTestResultsResponse) GetBody() *ListDataAgentAccuracyTestResultsResponseBody {
	return s.Body
}

func (s *ListDataAgentAccuracyTestResultsResponse) SetHeaders(v map[string]*string) *ListDataAgentAccuracyTestResultsResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponse) SetStatusCode(v int32) *ListDataAgentAccuracyTestResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponse) SetBody(v *ListDataAgentAccuracyTestResultsResponseBody) *ListDataAgentAccuracyTestResultsResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
