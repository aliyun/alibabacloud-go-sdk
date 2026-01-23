// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityIdentifyResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityIdentifyResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityIdentifyResultsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityIdentifyResultsResponseBody) *DeleteSecurityIdentifyResultsResponse
	GetBody() *DeleteSecurityIdentifyResultsResponseBody
}

type DeleteSecurityIdentifyResultsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityIdentifyResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityIdentifyResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIdentifyResultsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIdentifyResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityIdentifyResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityIdentifyResultsResponse) GetBody() *DeleteSecurityIdentifyResultsResponseBody {
	return s.Body
}

func (s *DeleteSecurityIdentifyResultsResponse) SetHeaders(v map[string]*string) *DeleteSecurityIdentifyResultsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponse) SetStatusCode(v int32) *DeleteSecurityIdentifyResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponse) SetBody(v *DeleteSecurityIdentifyResultsResponseBody) *DeleteSecurityIdentifyResultsResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityIdentifyResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
