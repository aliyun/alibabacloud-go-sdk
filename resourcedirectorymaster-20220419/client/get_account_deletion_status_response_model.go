// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountDeletionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountDeletionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountDeletionStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountDeletionStatusResponseBody) *GetAccountDeletionStatusResponse
	GetBody() *GetAccountDeletionStatusResponseBody
}

type GetAccountDeletionStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountDeletionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountDeletionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountDeletionStatusResponse) GetBody() *GetAccountDeletionStatusResponseBody {
	return s.Body
}

func (s *GetAccountDeletionStatusResponse) SetHeaders(v map[string]*string) *GetAccountDeletionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAccountDeletionStatusResponse) SetStatusCode(v int32) *GetAccountDeletionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountDeletionStatusResponse) SetBody(v *GetAccountDeletionStatusResponseBody) *GetAccountDeletionStatusResponse {
	s.Body = v
	return s
}

func (s *GetAccountDeletionStatusResponse) Validate() error {
	return dara.Validate(s)
}
