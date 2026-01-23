// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityIdentifyRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityIdentifyRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityIdentifyRecordsResponseBody) *ListSecurityIdentifyRecordsResponse
	GetBody() *ListSecurityIdentifyRecordsResponseBody
}

type ListSecurityIdentifyRecordsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityIdentifyRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityIdentifyRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityIdentifyRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityIdentifyRecordsResponse) GetBody() *ListSecurityIdentifyRecordsResponseBody {
	return s.Body
}

func (s *ListSecurityIdentifyRecordsResponse) SetHeaders(v map[string]*string) *ListSecurityIdentifyRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityIdentifyRecordsResponse) SetStatusCode(v int32) *ListSecurityIdentifyRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponse) SetBody(v *ListSecurityIdentifyRecordsResponseBody) *ListSecurityIdentifyRecordsResponse {
	s.Body = v
	return s
}

func (s *ListSecurityIdentifyRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
