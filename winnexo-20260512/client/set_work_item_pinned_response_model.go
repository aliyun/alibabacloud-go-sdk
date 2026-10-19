// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkItemPinnedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWorkItemPinnedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWorkItemPinnedResponse
	GetStatusCode() *int32
	SetBody(v *SetWorkItemPinnedResponseBody) *SetWorkItemPinnedResponse
	GetBody() *SetWorkItemPinnedResponseBody
}

type SetWorkItemPinnedResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWorkItemPinnedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWorkItemPinnedResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWorkItemPinnedResponse) GoString() string {
	return s.String()
}

func (s *SetWorkItemPinnedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWorkItemPinnedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWorkItemPinnedResponse) GetBody() *SetWorkItemPinnedResponseBody {
	return s.Body
}

func (s *SetWorkItemPinnedResponse) SetHeaders(v map[string]*string) *SetWorkItemPinnedResponse {
	s.Headers = v
	return s
}

func (s *SetWorkItemPinnedResponse) SetStatusCode(v int32) *SetWorkItemPinnedResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWorkItemPinnedResponse) SetBody(v *SetWorkItemPinnedResponseBody) *SetWorkItemPinnedResponse {
	s.Body = v
	return s
}

func (s *SetWorkItemPinnedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
