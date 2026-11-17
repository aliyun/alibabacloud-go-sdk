// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertDirectConversationToWorkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertDirectConversationToWorkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertDirectConversationToWorkResponse
	GetStatusCode() *int32
	SetBody(v *ConvertDirectConversationToWorkResponseBody) *ConvertDirectConversationToWorkResponse
	GetBody() *ConvertDirectConversationToWorkResponseBody
}

type ConvertDirectConversationToWorkResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertDirectConversationToWorkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertDirectConversationToWorkResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertDirectConversationToWorkResponse) GoString() string {
	return s.String()
}

func (s *ConvertDirectConversationToWorkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertDirectConversationToWorkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertDirectConversationToWorkResponse) GetBody() *ConvertDirectConversationToWorkResponseBody {
	return s.Body
}

func (s *ConvertDirectConversationToWorkResponse) SetHeaders(v map[string]*string) *ConvertDirectConversationToWorkResponse {
	s.Headers = v
	return s
}

func (s *ConvertDirectConversationToWorkResponse) SetStatusCode(v int32) *ConvertDirectConversationToWorkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponse) SetBody(v *ConvertDirectConversationToWorkResponseBody) *ConvertDirectConversationToWorkResponse {
	s.Body = v
	return s
}

func (s *ConvertDirectConversationToWorkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
