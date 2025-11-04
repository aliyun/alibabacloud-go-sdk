// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowAllOutputNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaConnectFlowAllOutputNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaConnectFlowAllOutputNameResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaConnectFlowAllOutputNameResponseBody) *GetMediaConnectFlowAllOutputNameResponse
	GetBody() *GetMediaConnectFlowAllOutputNameResponseBody
}

type GetMediaConnectFlowAllOutputNameResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaConnectFlowAllOutputNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaConnectFlowAllOutputNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowAllOutputNameResponse) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowAllOutputNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaConnectFlowAllOutputNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaConnectFlowAllOutputNameResponse) GetBody() *GetMediaConnectFlowAllOutputNameResponseBody {
	return s.Body
}

func (s *GetMediaConnectFlowAllOutputNameResponse) SetHeaders(v map[string]*string) *GetMediaConnectFlowAllOutputNameResponse {
	s.Headers = v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameResponse) SetStatusCode(v int32) *GetMediaConnectFlowAllOutputNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameResponse) SetBody(v *GetMediaConnectFlowAllOutputNameResponseBody) *GetMediaConnectFlowAllOutputNameResponse {
	s.Body = v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
