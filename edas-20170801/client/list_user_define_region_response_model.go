// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDefineRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserDefineRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserDefineRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListUserDefineRegionResponseBody) *ListUserDefineRegionResponse
	GetBody() *ListUserDefineRegionResponseBody
}

type ListUserDefineRegionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDefineRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDefineRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefineRegionResponse) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserDefineRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserDefineRegionResponse) GetBody() *ListUserDefineRegionResponseBody {
	return s.Body
}

func (s *ListUserDefineRegionResponse) SetHeaders(v map[string]*string) *ListUserDefineRegionResponse {
	s.Headers = v
	return s
}

func (s *ListUserDefineRegionResponse) SetStatusCode(v int32) *ListUserDefineRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDefineRegionResponse) SetBody(v *ListUserDefineRegionResponseBody) *ListUserDefineRegionResponse {
	s.Body = v
	return s
}

func (s *ListUserDefineRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
