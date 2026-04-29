// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePolarClawMCPServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePolarClawMCPServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePolarClawMCPServerResponse
	GetStatusCode() *int32
	SetBody(v *RemovePolarClawMCPServerResponseBody) *RemovePolarClawMCPServerResponse
	GetBody() *RemovePolarClawMCPServerResponseBody
}

type RemovePolarClawMCPServerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePolarClawMCPServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePolarClawMCPServerResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePolarClawMCPServerResponse) GoString() string {
	return s.String()
}

func (s *RemovePolarClawMCPServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePolarClawMCPServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePolarClawMCPServerResponse) GetBody() *RemovePolarClawMCPServerResponseBody {
	return s.Body
}

func (s *RemovePolarClawMCPServerResponse) SetHeaders(v map[string]*string) *RemovePolarClawMCPServerResponse {
	s.Headers = v
	return s
}

func (s *RemovePolarClawMCPServerResponse) SetStatusCode(v int32) *RemovePolarClawMCPServerResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePolarClawMCPServerResponse) SetBody(v *RemovePolarClawMCPServerResponseBody) *RemovePolarClawMCPServerResponse {
	s.Body = v
	return s
}

func (s *RemovePolarClawMCPServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
