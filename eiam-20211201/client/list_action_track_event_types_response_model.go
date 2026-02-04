// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionTrackEventTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListActionTrackEventTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListActionTrackEventTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListActionTrackEventTypesResponseBody) *ListActionTrackEventTypesResponse
	GetBody() *ListActionTrackEventTypesResponseBody
}

type ListActionTrackEventTypesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActionTrackEventTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActionTrackEventTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListActionTrackEventTypesResponse) GoString() string {
	return s.String()
}

func (s *ListActionTrackEventTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListActionTrackEventTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListActionTrackEventTypesResponse) GetBody() *ListActionTrackEventTypesResponseBody {
	return s.Body
}

func (s *ListActionTrackEventTypesResponse) SetHeaders(v map[string]*string) *ListActionTrackEventTypesResponse {
	s.Headers = v
	return s
}

func (s *ListActionTrackEventTypesResponse) SetStatusCode(v int32) *ListActionTrackEventTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActionTrackEventTypesResponse) SetBody(v *ListActionTrackEventTypesResponseBody) *ListActionTrackEventTypesResponse {
	s.Body = v
	return s
}

func (s *ListActionTrackEventTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
