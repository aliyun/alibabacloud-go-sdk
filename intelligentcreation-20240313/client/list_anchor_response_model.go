// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnchorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnchorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnchorResponse
	GetStatusCode() *int32
	SetBody(v *ListAnchorResponseBody) *ListAnchorResponse
	GetBody() *ListAnchorResponseBody
}

type ListAnchorResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnchorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnchorResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnchorResponse) GoString() string {
	return s.String()
}

func (s *ListAnchorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnchorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnchorResponse) GetBody() *ListAnchorResponseBody {
	return s.Body
}

func (s *ListAnchorResponse) SetHeaders(v map[string]*string) *ListAnchorResponse {
	s.Headers = v
	return s
}

func (s *ListAnchorResponse) SetStatusCode(v int32) *ListAnchorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnchorResponse) SetBody(v *ListAnchorResponseBody) *ListAnchorResponse {
	s.Body = v
	return s
}

func (s *ListAnchorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
