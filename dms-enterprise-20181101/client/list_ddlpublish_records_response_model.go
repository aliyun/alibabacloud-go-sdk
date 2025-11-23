// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDDLPublishRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDDLPublishRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDDLPublishRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListDDLPublishRecordsResponseBody) *ListDDLPublishRecordsResponse
	GetBody() *ListDDLPublishRecordsResponseBody
}

type ListDDLPublishRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDDLPublishRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDDLPublishRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDDLPublishRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDDLPublishRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDDLPublishRecordsResponse) GetBody() *ListDDLPublishRecordsResponseBody {
	return s.Body
}

func (s *ListDDLPublishRecordsResponse) SetHeaders(v map[string]*string) *ListDDLPublishRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListDDLPublishRecordsResponse) SetStatusCode(v int32) *ListDDLPublishRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDDLPublishRecordsResponse) SetBody(v *ListDDLPublishRecordsResponseBody) *ListDDLPublishRecordsResponse {
	s.Body = v
	return s
}

func (s *ListDDLPublishRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
