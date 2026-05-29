// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAutomaticWriteOffChangeRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAutomaticWriteOffChangeRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAutomaticWriteOffChangeRecordsResponse
	GetStatusCode() *int32
	SetBody(v *QueryAutomaticWriteOffChangeRecordsResponseBody) *QueryAutomaticWriteOffChangeRecordsResponse
	GetBody() *QueryAutomaticWriteOffChangeRecordsResponseBody
}

type QueryAutomaticWriteOffChangeRecordsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAutomaticWriteOffChangeRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAutomaticWriteOffChangeRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAutomaticWriteOffChangeRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryAutomaticWriteOffChangeRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAutomaticWriteOffChangeRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAutomaticWriteOffChangeRecordsResponse) GetBody() *QueryAutomaticWriteOffChangeRecordsResponseBody {
	return s.Body
}

func (s *QueryAutomaticWriteOffChangeRecordsResponse) SetHeaders(v map[string]*string) *QueryAutomaticWriteOffChangeRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponse) SetStatusCode(v int32) *QueryAutomaticWriteOffChangeRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponse) SetBody(v *QueryAutomaticWriteOffChangeRecordsResponseBody) *QueryAutomaticWriteOffChangeRecordsResponse {
	s.Body = v
	return s
}

func (s *QueryAutomaticWriteOffChangeRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
