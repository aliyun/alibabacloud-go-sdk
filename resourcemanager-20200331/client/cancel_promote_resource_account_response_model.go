// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPromoteResourceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPromoteResourceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPromoteResourceAccountResponse
	GetStatusCode() *int32
	SetBody(v *CancelPromoteResourceAccountResponseBody) *CancelPromoteResourceAccountResponse
	GetBody() *CancelPromoteResourceAccountResponseBody
}

type CancelPromoteResourceAccountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPromoteResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPromoteResourceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPromoteResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CancelPromoteResourceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPromoteResourceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPromoteResourceAccountResponse) GetBody() *CancelPromoteResourceAccountResponseBody {
	return s.Body
}

func (s *CancelPromoteResourceAccountResponse) SetHeaders(v map[string]*string) *CancelPromoteResourceAccountResponse {
	s.Headers = v
	return s
}

func (s *CancelPromoteResourceAccountResponse) SetStatusCode(v int32) *CancelPromoteResourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPromoteResourceAccountResponse) SetBody(v *CancelPromoteResourceAccountResponseBody) *CancelPromoteResourceAccountResponse {
	s.Body = v
	return s
}

func (s *CancelPromoteResourceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
