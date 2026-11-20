// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretTransferRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecretTransferRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecretTransferRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecretTransferRecordResponseBody) *CreateSecretTransferRecordResponse
	GetBody() *CreateSecretTransferRecordResponseBody
}

type CreateSecretTransferRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecretTransferRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecretTransferRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretTransferRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretTransferRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecretTransferRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecretTransferRecordResponse) GetBody() *CreateSecretTransferRecordResponseBody {
	return s.Body
}

func (s *CreateSecretTransferRecordResponse) SetHeaders(v map[string]*string) *CreateSecretTransferRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretTransferRecordResponse) SetStatusCode(v int32) *CreateSecretTransferRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecretTransferRecordResponse) SetBody(v *CreateSecretTransferRecordResponseBody) *CreateSecretTransferRecordResponse {
	s.Body = v
	return s
}

func (s *CreateSecretTransferRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
