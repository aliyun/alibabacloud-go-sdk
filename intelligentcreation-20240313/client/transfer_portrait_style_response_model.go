// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferPortraitStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferPortraitStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferPortraitStyleResponse
	GetStatusCode() *int32
	SetBody(v *TransferPortraitStyleResponseBody) *TransferPortraitStyleResponse
	GetBody() *TransferPortraitStyleResponseBody
}

type TransferPortraitStyleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferPortraitStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferPortraitStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferPortraitStyleResponse) GoString() string {
	return s.String()
}

func (s *TransferPortraitStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferPortraitStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferPortraitStyleResponse) GetBody() *TransferPortraitStyleResponseBody {
	return s.Body
}

func (s *TransferPortraitStyleResponse) SetHeaders(v map[string]*string) *TransferPortraitStyleResponse {
	s.Headers = v
	return s
}

func (s *TransferPortraitStyleResponse) SetStatusCode(v int32) *TransferPortraitStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferPortraitStyleResponse) SetBody(v *TransferPortraitStyleResponseBody) *TransferPortraitStyleResponse {
	s.Body = v
	return s
}

func (s *TransferPortraitStyleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
