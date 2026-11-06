// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRegionIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRegionIdResponse
	GetStatusCode() *int32
	SetBody(v *ListRegionIdResponseBody) *ListRegionIdResponse
	GetBody() *ListRegionIdResponseBody
}

type ListRegionIdResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRegionIdResponse) GoString() string {
	return s.String()
}

func (s *ListRegionIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRegionIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRegionIdResponse) GetBody() *ListRegionIdResponseBody {
	return s.Body
}

func (s *ListRegionIdResponse) SetHeaders(v map[string]*string) *ListRegionIdResponse {
	s.Headers = v
	return s
}

func (s *ListRegionIdResponse) SetStatusCode(v int32) *ListRegionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionIdResponse) SetBody(v *ListRegionIdResponseBody) *ListRegionIdResponse {
	s.Body = v
	return s
}

func (s *ListRegionIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
