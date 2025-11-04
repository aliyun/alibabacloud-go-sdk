// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaConnectFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaConnectFlowResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaConnectFlowResponseBody) *GetMediaConnectFlowResponse
	GetBody() *GetMediaConnectFlowResponseBody
}

type GetMediaConnectFlowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaConnectFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaConnectFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowResponse) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaConnectFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaConnectFlowResponse) GetBody() *GetMediaConnectFlowResponseBody {
	return s.Body
}

func (s *GetMediaConnectFlowResponse) SetHeaders(v map[string]*string) *GetMediaConnectFlowResponse {
	s.Headers = v
	return s
}

func (s *GetMediaConnectFlowResponse) SetStatusCode(v int32) *GetMediaConnectFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaConnectFlowResponse) SetBody(v *GetMediaConnectFlowResponseBody) *GetMediaConnectFlowResponse {
	s.Body = v
	return s
}

func (s *GetMediaConnectFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
