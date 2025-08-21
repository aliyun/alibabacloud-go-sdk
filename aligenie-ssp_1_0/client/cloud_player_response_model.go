// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPlayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudPlayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudPlayerResponse
	GetStatusCode() *int32
	SetBody(v *CloudPlayerResponseBody) *CloudPlayerResponse
	GetBody() *CloudPlayerResponseBody
}

type CloudPlayerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudPlayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudPlayerResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudPlayerResponse) GoString() string {
	return s.String()
}

func (s *CloudPlayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudPlayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudPlayerResponse) GetBody() *CloudPlayerResponseBody {
	return s.Body
}

func (s *CloudPlayerResponse) SetHeaders(v map[string]*string) *CloudPlayerResponse {
	s.Headers = v
	return s
}

func (s *CloudPlayerResponse) SetStatusCode(v int32) *CloudPlayerResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudPlayerResponse) SetBody(v *CloudPlayerResponseBody) *CloudPlayerResponse {
	s.Body = v
	return s
}

func (s *CloudPlayerResponse) Validate() error {
	return dara.Validate(s)
}
