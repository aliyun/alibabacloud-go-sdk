// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceProjectWhiteListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceProjectWhiteListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceProjectWhiteListsResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceProjectWhiteListsResponseBody) *ReplaceProjectWhiteListsResponse
	GetBody() *ReplaceProjectWhiteListsResponseBody
}

type ReplaceProjectWhiteListsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceProjectWhiteListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceProjectWhiteListsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceProjectWhiteListsResponse) GoString() string {
	return s.String()
}

func (s *ReplaceProjectWhiteListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceProjectWhiteListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceProjectWhiteListsResponse) GetBody() *ReplaceProjectWhiteListsResponseBody {
	return s.Body
}

func (s *ReplaceProjectWhiteListsResponse) SetHeaders(v map[string]*string) *ReplaceProjectWhiteListsResponse {
	s.Headers = v
	return s
}

func (s *ReplaceProjectWhiteListsResponse) SetStatusCode(v int32) *ReplaceProjectWhiteListsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceProjectWhiteListsResponse) SetBody(v *ReplaceProjectWhiteListsResponseBody) *ReplaceProjectWhiteListsResponse {
	s.Body = v
	return s
}

func (s *ReplaceProjectWhiteListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
