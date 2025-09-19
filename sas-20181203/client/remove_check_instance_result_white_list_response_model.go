// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCheckInstanceResultWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveCheckInstanceResultWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveCheckInstanceResultWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *RemoveCheckInstanceResultWhiteListResponseBody) *RemoveCheckInstanceResultWhiteListResponse
	GetBody() *RemoveCheckInstanceResultWhiteListResponseBody
}

type RemoveCheckInstanceResultWhiteListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCheckInstanceResultWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCheckInstanceResultWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveCheckInstanceResultWhiteListResponse) GoString() string {
	return s.String()
}

func (s *RemoveCheckInstanceResultWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveCheckInstanceResultWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveCheckInstanceResultWhiteListResponse) GetBody() *RemoveCheckInstanceResultWhiteListResponseBody {
	return s.Body
}

func (s *RemoveCheckInstanceResultWhiteListResponse) SetHeaders(v map[string]*string) *RemoveCheckInstanceResultWhiteListResponse {
	s.Headers = v
	return s
}

func (s *RemoveCheckInstanceResultWhiteListResponse) SetStatusCode(v int32) *RemoveCheckInstanceResultWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCheckInstanceResultWhiteListResponse) SetBody(v *RemoveCheckInstanceResultWhiteListResponseBody) *RemoveCheckInstanceResultWhiteListResponse {
	s.Body = v
	return s
}

func (s *RemoveCheckInstanceResultWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
