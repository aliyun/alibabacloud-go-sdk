// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOncallSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOncallSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOncallSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *ListOncallSchedulesResponseBody) *ListOncallSchedulesResponse
	GetBody() *ListOncallSchedulesResponseBody
}

type ListOncallSchedulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOncallSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOncallSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOncallSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListOncallSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOncallSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOncallSchedulesResponse) GetBody() *ListOncallSchedulesResponseBody {
	return s.Body
}

func (s *ListOncallSchedulesResponse) SetHeaders(v map[string]*string) *ListOncallSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListOncallSchedulesResponse) SetStatusCode(v int32) *ListOncallSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOncallSchedulesResponse) SetBody(v *ListOncallSchedulesResponseBody) *ListOncallSchedulesResponse {
	s.Body = v
	return s
}

func (s *ListOncallSchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
