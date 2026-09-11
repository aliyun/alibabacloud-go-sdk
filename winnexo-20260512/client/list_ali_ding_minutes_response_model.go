// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliDingMinutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAliDingMinutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAliDingMinutesResponse
	GetStatusCode() *int32
	SetBody(v *ListAliDingMinutesResponseBody) *ListAliDingMinutesResponse
	GetBody() *ListAliDingMinutesResponseBody
}

type ListAliDingMinutesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliDingMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliDingMinutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingMinutesResponse) GoString() string {
	return s.String()
}

func (s *ListAliDingMinutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAliDingMinutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAliDingMinutesResponse) GetBody() *ListAliDingMinutesResponseBody {
	return s.Body
}

func (s *ListAliDingMinutesResponse) SetHeaders(v map[string]*string) *ListAliDingMinutesResponse {
	s.Headers = v
	return s
}

func (s *ListAliDingMinutesResponse) SetStatusCode(v int32) *ListAliDingMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliDingMinutesResponse) SetBody(v *ListAliDingMinutesResponseBody) *ListAliDingMinutesResponse {
	s.Body = v
	return s
}

func (s *ListAliDingMinutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
