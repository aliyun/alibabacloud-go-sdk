// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferPayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferPayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferPayTypeResponse
	GetStatusCode() *int32
	SetBody(v *TransferPayTypeResponseBody) *TransferPayTypeResponse
	GetBody() *TransferPayTypeResponseBody
}

type TransferPayTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferPayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferPayTypeResponse) GoString() string {
	return s.String()
}

func (s *TransferPayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferPayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferPayTypeResponse) GetBody() *TransferPayTypeResponseBody {
	return s.Body
}

func (s *TransferPayTypeResponse) SetHeaders(v map[string]*string) *TransferPayTypeResponse {
	s.Headers = v
	return s
}

func (s *TransferPayTypeResponse) SetStatusCode(v int32) *TransferPayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferPayTypeResponse) SetBody(v *TransferPayTypeResponseBody) *TransferPayTypeResponse {
	s.Body = v
	return s
}

func (s *TransferPayTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
