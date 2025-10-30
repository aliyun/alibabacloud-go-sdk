// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubmitRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubmitRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubmitRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListSubmitRecordsResponseBody) *ListSubmitRecordsResponse
	GetBody() *ListSubmitRecordsResponseBody
}

type ListSubmitRecordsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubmitRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubmitRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubmitRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubmitRecordsResponse) GetBody() *ListSubmitRecordsResponseBody {
	return s.Body
}

func (s *ListSubmitRecordsResponse) SetHeaders(v map[string]*string) *ListSubmitRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListSubmitRecordsResponse) SetStatusCode(v int32) *ListSubmitRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubmitRecordsResponse) SetBody(v *ListSubmitRecordsResponseBody) *ListSubmitRecordsResponse {
	s.Body = v
	return s
}

func (s *ListSubmitRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
