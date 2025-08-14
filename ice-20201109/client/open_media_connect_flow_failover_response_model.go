// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenMediaConnectFlowFailoverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenMediaConnectFlowFailoverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenMediaConnectFlowFailoverResponse
	GetStatusCode() *int32
	SetBody(v *OpenMediaConnectFlowFailoverResponseBody) *OpenMediaConnectFlowFailoverResponse
	GetBody() *OpenMediaConnectFlowFailoverResponseBody
}

type OpenMediaConnectFlowFailoverResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenMediaConnectFlowFailoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenMediaConnectFlowFailoverResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenMediaConnectFlowFailoverResponse) GoString() string {
	return s.String()
}

func (s *OpenMediaConnectFlowFailoverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenMediaConnectFlowFailoverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenMediaConnectFlowFailoverResponse) GetBody() *OpenMediaConnectFlowFailoverResponseBody {
	return s.Body
}

func (s *OpenMediaConnectFlowFailoverResponse) SetHeaders(v map[string]*string) *OpenMediaConnectFlowFailoverResponse {
	s.Headers = v
	return s
}

func (s *OpenMediaConnectFlowFailoverResponse) SetStatusCode(v int32) *OpenMediaConnectFlowFailoverResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenMediaConnectFlowFailoverResponse) SetBody(v *OpenMediaConnectFlowFailoverResponseBody) *OpenMediaConnectFlowFailoverResponse {
	s.Body = v
	return s
}

func (s *OpenMediaConnectFlowFailoverResponse) Validate() error {
	return dara.Validate(s)
}
