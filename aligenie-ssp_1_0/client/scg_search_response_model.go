// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScgSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScgSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScgSearchResponse
	GetStatusCode() *int32
	SetBody(v *ScgSearchResponseBody) *ScgSearchResponse
	GetBody() *ScgSearchResponseBody
}

type ScgSearchResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScgSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScgSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchResponse) GoString() string {
	return s.String()
}

func (s *ScgSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScgSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScgSearchResponse) GetBody() *ScgSearchResponseBody {
	return s.Body
}

func (s *ScgSearchResponse) SetHeaders(v map[string]*string) *ScgSearchResponse {
	s.Headers = v
	return s
}

func (s *ScgSearchResponse) SetStatusCode(v int32) *ScgSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ScgSearchResponse) SetBody(v *ScgSearchResponseBody) *ScgSearchResponse {
	s.Body = v
	return s
}

func (s *ScgSearchResponse) Validate() error {
	return dara.Validate(s)
}
