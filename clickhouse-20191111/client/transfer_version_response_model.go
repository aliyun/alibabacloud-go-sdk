// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferVersionResponse
	GetStatusCode() *int32
	SetBody(v *TransferVersionResponseBody) *TransferVersionResponse
	GetBody() *TransferVersionResponseBody
}

type TransferVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferVersionResponse) GoString() string {
	return s.String()
}

func (s *TransferVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferVersionResponse) GetBody() *TransferVersionResponseBody {
	return s.Body
}

func (s *TransferVersionResponse) SetHeaders(v map[string]*string) *TransferVersionResponse {
	s.Headers = v
	return s
}

func (s *TransferVersionResponse) SetStatusCode(v int32) *TransferVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferVersionResponse) SetBody(v *TransferVersionResponseBody) *TransferVersionResponse {
	s.Body = v
	return s
}

func (s *TransferVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
