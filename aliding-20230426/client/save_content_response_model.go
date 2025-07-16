// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveContentResponse
	GetStatusCode() *int32
	SetBody(v *SaveContentResponseBody) *SaveContentResponse
	GetBody() *SaveContentResponseBody
}

type SaveContentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveContentResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveContentResponse) GoString() string {
	return s.String()
}

func (s *SaveContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveContentResponse) GetBody() *SaveContentResponseBody {
	return s.Body
}

func (s *SaveContentResponse) SetHeaders(v map[string]*string) *SaveContentResponse {
	s.Headers = v
	return s
}

func (s *SaveContentResponse) SetStatusCode(v int32) *SaveContentResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveContentResponse) SetBody(v *SaveContentResponseBody) *SaveContentResponse {
	s.Body = v
	return s
}

func (s *SaveContentResponse) Validate() error {
	return dara.Validate(s)
}
