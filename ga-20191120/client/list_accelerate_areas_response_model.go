// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccelerateAreasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccelerateAreasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccelerateAreasResponse
	GetStatusCode() *int32
	SetBody(v *ListAccelerateAreasResponseBody) *ListAccelerateAreasResponse
	GetBody() *ListAccelerateAreasResponseBody
}

type ListAccelerateAreasResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccelerateAreasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccelerateAreasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerateAreasResponse) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccelerateAreasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccelerateAreasResponse) GetBody() *ListAccelerateAreasResponseBody {
	return s.Body
}

func (s *ListAccelerateAreasResponse) SetHeaders(v map[string]*string) *ListAccelerateAreasResponse {
	s.Headers = v
	return s
}

func (s *ListAccelerateAreasResponse) SetStatusCode(v int32) *ListAccelerateAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccelerateAreasResponse) SetBody(v *ListAccelerateAreasResponseBody) *ListAccelerateAreasResponse {
	s.Body = v
	return s
}

func (s *ListAccelerateAreasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
