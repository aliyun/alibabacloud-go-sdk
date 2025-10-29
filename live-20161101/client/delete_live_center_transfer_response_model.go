// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveCenterTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveCenterTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveCenterTransferResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveCenterTransferResponseBody) *DeleteLiveCenterTransferResponse
	GetBody() *DeleteLiveCenterTransferResponseBody
}

type DeleteLiveCenterTransferResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveCenterTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveCenterTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveCenterTransferResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveCenterTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveCenterTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveCenterTransferResponse) GetBody() *DeleteLiveCenterTransferResponseBody {
	return s.Body
}

func (s *DeleteLiveCenterTransferResponse) SetHeaders(v map[string]*string) *DeleteLiveCenterTransferResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveCenterTransferResponse) SetStatusCode(v int32) *DeleteLiveCenterTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveCenterTransferResponse) SetBody(v *DeleteLiveCenterTransferResponseBody) *DeleteLiveCenterTransferResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveCenterTransferResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
