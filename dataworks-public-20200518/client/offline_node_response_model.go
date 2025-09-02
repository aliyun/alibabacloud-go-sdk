// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineNodeResponse
	GetStatusCode() *int32
	SetBody(v *OfflineNodeResponseBody) *OfflineNodeResponse
	GetBody() *OfflineNodeResponseBody
}

type OfflineNodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineNodeResponse) GoString() string {
	return s.String()
}

func (s *OfflineNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineNodeResponse) GetBody() *OfflineNodeResponseBody {
	return s.Body
}

func (s *OfflineNodeResponse) SetHeaders(v map[string]*string) *OfflineNodeResponse {
	s.Headers = v
	return s
}

func (s *OfflineNodeResponse) SetStatusCode(v int32) *OfflineNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineNodeResponse) SetBody(v *OfflineNodeResponseBody) *OfflineNodeResponse {
	s.Body = v
	return s
}

func (s *OfflineNodeResponse) Validate() error {
	return dara.Validate(s)
}
