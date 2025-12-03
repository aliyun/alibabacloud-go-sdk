// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveListenerWhiteListItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveListenerWhiteListItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveListenerWhiteListItemResponse
	GetStatusCode() *int32
	SetBody(v *RemoveListenerWhiteListItemResponseBody) *RemoveListenerWhiteListItemResponse
	GetBody() *RemoveListenerWhiteListItemResponseBody
}

type RemoveListenerWhiteListItemResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveListenerWhiteListItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveListenerWhiteListItemResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveListenerWhiteListItemResponse) GoString() string {
	return s.String()
}

func (s *RemoveListenerWhiteListItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveListenerWhiteListItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveListenerWhiteListItemResponse) GetBody() *RemoveListenerWhiteListItemResponseBody {
	return s.Body
}

func (s *RemoveListenerWhiteListItemResponse) SetHeaders(v map[string]*string) *RemoveListenerWhiteListItemResponse {
	s.Headers = v
	return s
}

func (s *RemoveListenerWhiteListItemResponse) SetStatusCode(v int32) *RemoveListenerWhiteListItemResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveListenerWhiteListItemResponse) SetBody(v *RemoveListenerWhiteListItemResponseBody) *RemoveListenerWhiteListItemResponse {
	s.Body = v
	return s
}

func (s *RemoveListenerWhiteListItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
