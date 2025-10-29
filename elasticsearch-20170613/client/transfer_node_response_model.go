// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferNodeResponse
	GetStatusCode() *int32
	SetBody(v *TransferNodeResponseBody) *TransferNodeResponse
	GetBody() *TransferNodeResponseBody
}

type TransferNodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferNodeResponse) GoString() string {
	return s.String()
}

func (s *TransferNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferNodeResponse) GetBody() *TransferNodeResponseBody {
	return s.Body
}

func (s *TransferNodeResponse) SetHeaders(v map[string]*string) *TransferNodeResponse {
	s.Headers = v
	return s
}

func (s *TransferNodeResponse) SetStatusCode(v int32) *TransferNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferNodeResponse) SetBody(v *TransferNodeResponseBody) *TransferNodeResponse {
	s.Body = v
	return s
}

func (s *TransferNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
