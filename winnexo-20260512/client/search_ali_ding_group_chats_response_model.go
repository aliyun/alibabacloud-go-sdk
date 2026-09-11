// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAliDingGroupChatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAliDingGroupChatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAliDingGroupChatsResponse
	GetStatusCode() *int32
	SetBody(v *SearchAliDingGroupChatsResponseBody) *SearchAliDingGroupChatsResponse
	GetBody() *SearchAliDingGroupChatsResponseBody
}

type SearchAliDingGroupChatsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAliDingGroupChatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAliDingGroupChatsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAliDingGroupChatsResponse) GoString() string {
	return s.String()
}

func (s *SearchAliDingGroupChatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAliDingGroupChatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAliDingGroupChatsResponse) GetBody() *SearchAliDingGroupChatsResponseBody {
	return s.Body
}

func (s *SearchAliDingGroupChatsResponse) SetHeaders(v map[string]*string) *SearchAliDingGroupChatsResponse {
	s.Headers = v
	return s
}

func (s *SearchAliDingGroupChatsResponse) SetStatusCode(v int32) *SearchAliDingGroupChatsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAliDingGroupChatsResponse) SetBody(v *SearchAliDingGroupChatsResponseBody) *SearchAliDingGroupChatsResponse {
	s.Body = v
	return s
}

func (s *SearchAliDingGroupChatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
