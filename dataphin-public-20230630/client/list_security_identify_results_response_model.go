// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityIdentifyResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityIdentifyResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityIdentifyResultsResponseBody) *ListSecurityIdentifyResultsResponse
	GetBody() *ListSecurityIdentifyResultsResponseBody
}

type ListSecurityIdentifyResultsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityIdentifyResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityIdentifyResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyResultsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityIdentifyResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityIdentifyResultsResponse) GetBody() *ListSecurityIdentifyResultsResponseBody {
	return s.Body
}

func (s *ListSecurityIdentifyResultsResponse) SetHeaders(v map[string]*string) *ListSecurityIdentifyResultsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityIdentifyResultsResponse) SetStatusCode(v int32) *ListSecurityIdentifyResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponse) SetBody(v *ListSecurityIdentifyResultsResponseBody) *ListSecurityIdentifyResultsResponse {
	s.Body = v
	return s
}

func (s *ListSecurityIdentifyResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
