// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAWSRegionInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAWSRegionInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAWSRegionInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListAWSRegionInfosResponseBody) *ListAWSRegionInfosResponse
	GetBody() *ListAWSRegionInfosResponseBody
}

type ListAWSRegionInfosResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAWSRegionInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAWSRegionInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAWSRegionInfosResponse) GoString() string {
	return s.String()
}

func (s *ListAWSRegionInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAWSRegionInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAWSRegionInfosResponse) GetBody() *ListAWSRegionInfosResponseBody {
	return s.Body
}

func (s *ListAWSRegionInfosResponse) SetHeaders(v map[string]*string) *ListAWSRegionInfosResponse {
	s.Headers = v
	return s
}

func (s *ListAWSRegionInfosResponse) SetStatusCode(v int32) *ListAWSRegionInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAWSRegionInfosResponse) SetBody(v *ListAWSRegionInfosResponseBody) *ListAWSRegionInfosResponse {
	s.Body = v
	return s
}

func (s *ListAWSRegionInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
