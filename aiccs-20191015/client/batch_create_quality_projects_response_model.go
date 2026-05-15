// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateQualityProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateQualityProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateQualityProjectsResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateQualityProjectsResponseBody) *BatchCreateQualityProjectsResponse
	GetBody() *BatchCreateQualityProjectsResponseBody
}

type BatchCreateQualityProjectsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateQualityProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateQualityProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateQualityProjectsResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateQualityProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateQualityProjectsResponse) GetBody() *BatchCreateQualityProjectsResponseBody {
	return s.Body
}

func (s *BatchCreateQualityProjectsResponse) SetHeaders(v map[string]*string) *BatchCreateQualityProjectsResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateQualityProjectsResponse) SetStatusCode(v int32) *BatchCreateQualityProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateQualityProjectsResponse) SetBody(v *BatchCreateQualityProjectsResponseBody) *BatchCreateQualityProjectsResponse {
	s.Body = v
	return s
}

func (s *BatchCreateQualityProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
