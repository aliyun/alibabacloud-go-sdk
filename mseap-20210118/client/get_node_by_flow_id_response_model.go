// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByFlowIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeByFlowIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeByFlowIdResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeByFlowIdResponseBody) *GetNodeByFlowIdResponse
	GetBody() *GetNodeByFlowIdResponseBody
}

type GetNodeByFlowIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeByFlowIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeByFlowIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByFlowIdResponse) GoString() string {
	return s.String()
}

func (s *GetNodeByFlowIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeByFlowIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeByFlowIdResponse) GetBody() *GetNodeByFlowIdResponseBody {
	return s.Body
}

func (s *GetNodeByFlowIdResponse) SetHeaders(v map[string]*string) *GetNodeByFlowIdResponse {
	s.Headers = v
	return s
}

func (s *GetNodeByFlowIdResponse) SetStatusCode(v int32) *GetNodeByFlowIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeByFlowIdResponse) SetBody(v *GetNodeByFlowIdResponseBody) *GetNodeByFlowIdResponse {
	s.Body = v
	return s
}

func (s *GetNodeByFlowIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
