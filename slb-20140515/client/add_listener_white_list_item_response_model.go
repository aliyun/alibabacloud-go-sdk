// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddListenerWhiteListItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddListenerWhiteListItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddListenerWhiteListItemResponse
	GetStatusCode() *int32
	SetBody(v *AddListenerWhiteListItemResponseBody) *AddListenerWhiteListItemResponse
	GetBody() *AddListenerWhiteListItemResponseBody
}

type AddListenerWhiteListItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddListenerWhiteListItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddListenerWhiteListItemResponse) String() string {
	return dara.Prettify(s)
}

func (s AddListenerWhiteListItemResponse) GoString() string {
	return s.String()
}

func (s *AddListenerWhiteListItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddListenerWhiteListItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddListenerWhiteListItemResponse) GetBody() *AddListenerWhiteListItemResponseBody {
	return s.Body
}

func (s *AddListenerWhiteListItemResponse) SetHeaders(v map[string]*string) *AddListenerWhiteListItemResponse {
	s.Headers = v
	return s
}

func (s *AddListenerWhiteListItemResponse) SetStatusCode(v int32) *AddListenerWhiteListItemResponse {
	s.StatusCode = &v
	return s
}

func (s *AddListenerWhiteListItemResponse) SetBody(v *AddListenerWhiteListItemResponseBody) *AddListenerWhiteListItemResponse {
	s.Body = v
	return s
}

func (s *AddListenerWhiteListItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
