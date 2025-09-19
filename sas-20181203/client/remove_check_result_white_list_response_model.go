// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCheckResultWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveCheckResultWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveCheckResultWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *RemoveCheckResultWhiteListResponseBody) *RemoveCheckResultWhiteListResponse
	GetBody() *RemoveCheckResultWhiteListResponseBody
}

type RemoveCheckResultWhiteListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCheckResultWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCheckResultWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveCheckResultWhiteListResponse) GoString() string {
	return s.String()
}

func (s *RemoveCheckResultWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveCheckResultWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveCheckResultWhiteListResponse) GetBody() *RemoveCheckResultWhiteListResponseBody {
	return s.Body
}

func (s *RemoveCheckResultWhiteListResponse) SetHeaders(v map[string]*string) *RemoveCheckResultWhiteListResponse {
	s.Headers = v
	return s
}

func (s *RemoveCheckResultWhiteListResponse) SetStatusCode(v int32) *RemoveCheckResultWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCheckResultWhiteListResponse) SetBody(v *RemoveCheckResultWhiteListResponseBody) *RemoveCheckResultWhiteListResponse {
	s.Body = v
	return s
}

func (s *RemoveCheckResultWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
