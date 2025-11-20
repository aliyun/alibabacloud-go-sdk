// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversaionSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConversaionSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConversaionSpaceResponse
	GetStatusCode() *int32
	SetBody(v *GetConversaionSpaceResponseBody) *GetConversaionSpaceResponse
	GetBody() *GetConversaionSpaceResponseBody
}

type GetConversaionSpaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversaionSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversaionSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConversaionSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConversaionSpaceResponse) GetBody() *GetConversaionSpaceResponseBody {
	return s.Body
}

func (s *GetConversaionSpaceResponse) SetHeaders(v map[string]*string) *GetConversaionSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetConversaionSpaceResponse) SetStatusCode(v int32) *GetConversaionSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversaionSpaceResponse) SetBody(v *GetConversaionSpaceResponseBody) *GetConversaionSpaceResponse {
	s.Body = v
	return s
}

func (s *GetConversaionSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
