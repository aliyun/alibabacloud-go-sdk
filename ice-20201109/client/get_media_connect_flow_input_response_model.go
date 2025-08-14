// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaConnectFlowInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaConnectFlowInputResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaConnectFlowInputResponseBody) *GetMediaConnectFlowInputResponse
	GetBody() *GetMediaConnectFlowInputResponseBody
}

type GetMediaConnectFlowInputResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaConnectFlowInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaConnectFlowInputResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowInputResponse) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaConnectFlowInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaConnectFlowInputResponse) GetBody() *GetMediaConnectFlowInputResponseBody {
	return s.Body
}

func (s *GetMediaConnectFlowInputResponse) SetHeaders(v map[string]*string) *GetMediaConnectFlowInputResponse {
	s.Headers = v
	return s
}

func (s *GetMediaConnectFlowInputResponse) SetStatusCode(v int32) *GetMediaConnectFlowInputResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaConnectFlowInputResponse) SetBody(v *GetMediaConnectFlowInputResponseBody) *GetMediaConnectFlowInputResponse {
	s.Body = v
	return s
}

func (s *GetMediaConnectFlowInputResponse) Validate() error {
	return dara.Validate(s)
}
