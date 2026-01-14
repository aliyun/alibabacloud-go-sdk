// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBlackWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBlackWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBlackWhiteListResponseBody) *DeleteBlackWhiteListResponse
	GetBody() *DeleteBlackWhiteListResponseBody
}

type DeleteBlackWhiteListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBlackWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBlackWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBlackWhiteListResponse) GetBody() *DeleteBlackWhiteListResponseBody {
	return s.Body
}

func (s *DeleteBlackWhiteListResponse) SetHeaders(v map[string]*string) *DeleteBlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DeleteBlackWhiteListResponse) SetStatusCode(v int32) *DeleteBlackWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBlackWhiteListResponse) SetBody(v *DeleteBlackWhiteListResponseBody) *DeleteBlackWhiteListResponse {
	s.Body = v
	return s
}

func (s *DeleteBlackWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
