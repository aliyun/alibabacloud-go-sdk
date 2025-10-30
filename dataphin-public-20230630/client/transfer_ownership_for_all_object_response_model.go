// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferOwnershipForAllObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferOwnershipForAllObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferOwnershipForAllObjectResponse
	GetStatusCode() *int32
	SetBody(v *TransferOwnershipForAllObjectResponseBody) *TransferOwnershipForAllObjectResponse
	GetBody() *TransferOwnershipForAllObjectResponseBody
}

type TransferOwnershipForAllObjectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferOwnershipForAllObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferOwnershipForAllObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferOwnershipForAllObjectResponse) GoString() string {
	return s.String()
}

func (s *TransferOwnershipForAllObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferOwnershipForAllObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferOwnershipForAllObjectResponse) GetBody() *TransferOwnershipForAllObjectResponseBody {
	return s.Body
}

func (s *TransferOwnershipForAllObjectResponse) SetHeaders(v map[string]*string) *TransferOwnershipForAllObjectResponse {
	s.Headers = v
	return s
}

func (s *TransferOwnershipForAllObjectResponse) SetStatusCode(v int32) *TransferOwnershipForAllObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferOwnershipForAllObjectResponse) SetBody(v *TransferOwnershipForAllObjectResponseBody) *TransferOwnershipForAllObjectResponse {
	s.Body = v
	return s
}

func (s *TransferOwnershipForAllObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
