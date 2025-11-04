// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaConnectFlowOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaConnectFlowOutputResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaConnectFlowOutputResponseBody) *GetMediaConnectFlowOutputResponse
	GetBody() *GetMediaConnectFlowOutputResponseBody
}

type GetMediaConnectFlowOutputResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaConnectFlowOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaConnectFlowOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowOutputResponse) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaConnectFlowOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaConnectFlowOutputResponse) GetBody() *GetMediaConnectFlowOutputResponseBody {
	return s.Body
}

func (s *GetMediaConnectFlowOutputResponse) SetHeaders(v map[string]*string) *GetMediaConnectFlowOutputResponse {
	s.Headers = v
	return s
}

func (s *GetMediaConnectFlowOutputResponse) SetStatusCode(v int32) *GetMediaConnectFlowOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaConnectFlowOutputResponse) SetBody(v *GetMediaConnectFlowOutputResponseBody) *GetMediaConnectFlowOutputResponse {
	s.Body = v
	return s
}

func (s *GetMediaConnectFlowOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
