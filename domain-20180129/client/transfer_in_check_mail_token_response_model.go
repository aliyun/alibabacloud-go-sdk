// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInCheckMailTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferInCheckMailTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferInCheckMailTokenResponse
	GetStatusCode() *int32
	SetBody(v *TransferInCheckMailTokenResponseBody) *TransferInCheckMailTokenResponse
	GetBody() *TransferInCheckMailTokenResponseBody
}

type TransferInCheckMailTokenResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferInCheckMailTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferInCheckMailTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferInCheckMailTokenResponse) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferInCheckMailTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferInCheckMailTokenResponse) GetBody() *TransferInCheckMailTokenResponseBody {
	return s.Body
}

func (s *TransferInCheckMailTokenResponse) SetHeaders(v map[string]*string) *TransferInCheckMailTokenResponse {
	s.Headers = v
	return s
}

func (s *TransferInCheckMailTokenResponse) SetStatusCode(v int32) *TransferInCheckMailTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInCheckMailTokenResponse) SetBody(v *TransferInCheckMailTokenResponseBody) *TransferInCheckMailTokenResponse {
	s.Body = v
	return s
}

func (s *TransferInCheckMailTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
