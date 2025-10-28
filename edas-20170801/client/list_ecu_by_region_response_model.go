// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcuByRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEcuByRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEcuByRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListEcuByRegionResponseBody) *ListEcuByRegionResponse
	GetBody() *ListEcuByRegionResponseBody
}

type ListEcuByRegionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEcuByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEcuByRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEcuByRegionResponse) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEcuByRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEcuByRegionResponse) GetBody() *ListEcuByRegionResponseBody {
	return s.Body
}

func (s *ListEcuByRegionResponse) SetHeaders(v map[string]*string) *ListEcuByRegionResponse {
	s.Headers = v
	return s
}

func (s *ListEcuByRegionResponse) SetStatusCode(v int32) *ListEcuByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEcuByRegionResponse) SetBody(v *ListEcuByRegionResponseBody) *ListEcuByRegionResponse {
	s.Body = v
	return s
}

func (s *ListEcuByRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
