// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallSyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNatFirewallSyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNatFirewallSyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateNatFirewallSyncTaskResponseBody) *CreateNatFirewallSyncTaskResponse
	GetBody() *CreateNatFirewallSyncTaskResponseBody
}

type CreateNatFirewallSyncTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNatFirewallSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNatFirewallSyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallSyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNatFirewallSyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNatFirewallSyncTaskResponse) GetBody() *CreateNatFirewallSyncTaskResponseBody {
	return s.Body
}

func (s *CreateNatFirewallSyncTaskResponse) SetHeaders(v map[string]*string) *CreateNatFirewallSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateNatFirewallSyncTaskResponse) SetStatusCode(v int32) *CreateNatFirewallSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNatFirewallSyncTaskResponse) SetBody(v *CreateNatFirewallSyncTaskResponseBody) *CreateNatFirewallSyncTaskResponse {
	s.Body = v
	return s
}

func (s *CreateNatFirewallSyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
