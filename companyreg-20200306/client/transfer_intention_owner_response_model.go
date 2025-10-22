// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferIntentionOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferIntentionOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferIntentionOwnerResponse
	GetStatusCode() *int32
	SetBody(v *TransferIntentionOwnerResponseBody) *TransferIntentionOwnerResponse
	GetBody() *TransferIntentionOwnerResponseBody
}

type TransferIntentionOwnerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferIntentionOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferIntentionOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferIntentionOwnerResponse) GoString() string {
	return s.String()
}

func (s *TransferIntentionOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferIntentionOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferIntentionOwnerResponse) GetBody() *TransferIntentionOwnerResponseBody {
	return s.Body
}

func (s *TransferIntentionOwnerResponse) SetHeaders(v map[string]*string) *TransferIntentionOwnerResponse {
	s.Headers = v
	return s
}

func (s *TransferIntentionOwnerResponse) SetStatusCode(v int32) *TransferIntentionOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferIntentionOwnerResponse) SetBody(v *TransferIntentionOwnerResponseBody) *TransferIntentionOwnerResponse {
	s.Body = v
	return s
}

func (s *TransferIntentionOwnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
