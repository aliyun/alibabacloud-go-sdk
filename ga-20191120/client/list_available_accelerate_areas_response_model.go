// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableAccelerateAreasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableAccelerateAreasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableAccelerateAreasResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableAccelerateAreasResponseBody) *ListAvailableAccelerateAreasResponse
	GetBody() *ListAvailableAccelerateAreasResponseBody
}

type ListAvailableAccelerateAreasResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableAccelerateAreasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableAccelerateAreasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableAccelerateAreasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableAccelerateAreasResponse) GetBody() *ListAvailableAccelerateAreasResponseBody {
	return s.Body
}

func (s *ListAvailableAccelerateAreasResponse) SetHeaders(v map[string]*string) *ListAvailableAccelerateAreasResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableAccelerateAreasResponse) SetStatusCode(v int32) *ListAvailableAccelerateAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponse) SetBody(v *ListAvailableAccelerateAreasResponseBody) *ListAvailableAccelerateAreasResponse {
	s.Body = v
	return s
}

func (s *ListAvailableAccelerateAreasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
