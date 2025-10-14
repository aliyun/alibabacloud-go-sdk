// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanDistDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CleanDistDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CleanDistDataResponse
	GetStatusCode() *int32
	SetBody(v *CleanDistDataResponseBody) *CleanDistDataResponse
	GetBody() *CleanDistDataResponseBody
}

type CleanDistDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CleanDistDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CleanDistDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CleanDistDataResponse) GoString() string {
	return s.String()
}

func (s *CleanDistDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CleanDistDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CleanDistDataResponse) GetBody() *CleanDistDataResponseBody {
	return s.Body
}

func (s *CleanDistDataResponse) SetHeaders(v map[string]*string) *CleanDistDataResponse {
	s.Headers = v
	return s
}

func (s *CleanDistDataResponse) SetStatusCode(v int32) *CleanDistDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CleanDistDataResponse) SetBody(v *CleanDistDataResponseBody) *CleanDistDataResponse {
	s.Body = v
	return s
}

func (s *CleanDistDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
