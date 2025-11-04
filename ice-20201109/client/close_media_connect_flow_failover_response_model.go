// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseMediaConnectFlowFailoverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseMediaConnectFlowFailoverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseMediaConnectFlowFailoverResponse
	GetStatusCode() *int32
	SetBody(v *CloseMediaConnectFlowFailoverResponseBody) *CloseMediaConnectFlowFailoverResponse
	GetBody() *CloseMediaConnectFlowFailoverResponseBody
}

type CloseMediaConnectFlowFailoverResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseMediaConnectFlowFailoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseMediaConnectFlowFailoverResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseMediaConnectFlowFailoverResponse) GoString() string {
	return s.String()
}

func (s *CloseMediaConnectFlowFailoverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseMediaConnectFlowFailoverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseMediaConnectFlowFailoverResponse) GetBody() *CloseMediaConnectFlowFailoverResponseBody {
	return s.Body
}

func (s *CloseMediaConnectFlowFailoverResponse) SetHeaders(v map[string]*string) *CloseMediaConnectFlowFailoverResponse {
	s.Headers = v
	return s
}

func (s *CloseMediaConnectFlowFailoverResponse) SetStatusCode(v int32) *CloseMediaConnectFlowFailoverResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseMediaConnectFlowFailoverResponse) SetBody(v *CloseMediaConnectFlowFailoverResponseBody) *CloseMediaConnectFlowFailoverResponse {
	s.Body = v
	return s
}

func (s *CloseMediaConnectFlowFailoverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
