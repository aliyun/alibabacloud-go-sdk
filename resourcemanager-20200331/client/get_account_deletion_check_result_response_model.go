// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountDeletionCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountDeletionCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountDeletionCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountDeletionCheckResultResponseBody) *GetAccountDeletionCheckResultResponse
	GetBody() *GetAccountDeletionCheckResultResponseBody
}

type GetAccountDeletionCheckResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountDeletionCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountDeletionCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponse) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountDeletionCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountDeletionCheckResultResponse) GetBody() *GetAccountDeletionCheckResultResponseBody {
	return s.Body
}

func (s *GetAccountDeletionCheckResultResponse) SetHeaders(v map[string]*string) *GetAccountDeletionCheckResultResponse {
	s.Headers = v
	return s
}

func (s *GetAccountDeletionCheckResultResponse) SetStatusCode(v int32) *GetAccountDeletionCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponse) SetBody(v *GetAccountDeletionCheckResultResponseBody) *GetAccountDeletionCheckResultResponse {
	s.Body = v
	return s
}

func (s *GetAccountDeletionCheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
