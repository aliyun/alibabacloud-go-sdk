// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationSubmitStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationSubmitStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationSubmitStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationSubmitStatusResponseBody) *GetOperationSubmitStatusResponse
	GetBody() *GetOperationSubmitStatusResponseBody
}

type GetOperationSubmitStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationSubmitStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationSubmitStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationSubmitStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationSubmitStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationSubmitStatusResponse) GetBody() *GetOperationSubmitStatusResponseBody {
	return s.Body
}

func (s *GetOperationSubmitStatusResponse) SetHeaders(v map[string]*string) *GetOperationSubmitStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOperationSubmitStatusResponse) SetStatusCode(v int32) *GetOperationSubmitStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationSubmitStatusResponse) SetBody(v *GetOperationSubmitStatusResponseBody) *GetOperationSubmitStatusResponse {
	s.Body = v
	return s
}

func (s *GetOperationSubmitStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
