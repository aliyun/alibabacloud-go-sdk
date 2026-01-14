// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonAreasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommonAreasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommonAreasResponse
	GetStatusCode() *int32
	SetBody(v *ListCommonAreasResponseBody) *ListCommonAreasResponse
	GetBody() *ListCommonAreasResponseBody
}

type ListCommonAreasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommonAreasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommonAreasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommonAreasResponse) GoString() string {
	return s.String()
}

func (s *ListCommonAreasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommonAreasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommonAreasResponse) GetBody() *ListCommonAreasResponseBody {
	return s.Body
}

func (s *ListCommonAreasResponse) SetHeaders(v map[string]*string) *ListCommonAreasResponse {
	s.Headers = v
	return s
}

func (s *ListCommonAreasResponse) SetStatusCode(v int32) *ListCommonAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonAreasResponse) SetBody(v *ListCommonAreasResponseBody) *ListCommonAreasResponse {
	s.Body = v
	return s
}

func (s *ListCommonAreasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
