// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSandboxRebuildOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSandboxRebuildOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSandboxRebuildOperationResponse
	GetStatusCode() *int32
	SetBody(v *GetSandboxRebuildOperationResponseBody) *GetSandboxRebuildOperationResponse
	GetBody() *GetSandboxRebuildOperationResponseBody
}

type GetSandboxRebuildOperationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSandboxRebuildOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSandboxRebuildOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSandboxRebuildOperationResponse) GoString() string {
	return s.String()
}

func (s *GetSandboxRebuildOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSandboxRebuildOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSandboxRebuildOperationResponse) GetBody() *GetSandboxRebuildOperationResponseBody {
	return s.Body
}

func (s *GetSandboxRebuildOperationResponse) SetHeaders(v map[string]*string) *GetSandboxRebuildOperationResponse {
	s.Headers = v
	return s
}

func (s *GetSandboxRebuildOperationResponse) SetStatusCode(v int32) *GetSandboxRebuildOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSandboxRebuildOperationResponse) SetBody(v *GetSandboxRebuildOperationResponseBody) *GetSandboxRebuildOperationResponse {
	s.Body = v
	return s
}

func (s *GetSandboxRebuildOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
