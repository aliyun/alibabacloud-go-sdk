// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentSubAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataAgentSubAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataAgentSubAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDataAgentSubAccountInfoResponseBody) *GetDataAgentSubAccountInfoResponse
	GetBody() *GetDataAgentSubAccountInfoResponseBody
}

type GetDataAgentSubAccountInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataAgentSubAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataAgentSubAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentSubAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDataAgentSubAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataAgentSubAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataAgentSubAccountInfoResponse) GetBody() *GetDataAgentSubAccountInfoResponseBody {
	return s.Body
}

func (s *GetDataAgentSubAccountInfoResponse) SetHeaders(v map[string]*string) *GetDataAgentSubAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDataAgentSubAccountInfoResponse) SetStatusCode(v int32) *GetDataAgentSubAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponse) SetBody(v *GetDataAgentSubAccountInfoResponseBody) *GetDataAgentSubAccountInfoResponse {
	s.Body = v
	return s
}

func (s *GetDataAgentSubAccountInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
