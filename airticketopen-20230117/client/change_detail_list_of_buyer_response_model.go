// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailListOfBuyerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDetailListOfBuyerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDetailListOfBuyerResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDetailListOfBuyerResponseBody) *ChangeDetailListOfBuyerResponse
	GetBody() *ChangeDetailListOfBuyerResponseBody
}

type ChangeDetailListOfBuyerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDetailListOfBuyerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDetailListOfBuyerResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponse) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDetailListOfBuyerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDetailListOfBuyerResponse) GetBody() *ChangeDetailListOfBuyerResponseBody {
	return s.Body
}

func (s *ChangeDetailListOfBuyerResponse) SetHeaders(v map[string]*string) *ChangeDetailListOfBuyerResponse {
	s.Headers = v
	return s
}

func (s *ChangeDetailListOfBuyerResponse) SetStatusCode(v int32) *ChangeDetailListOfBuyerResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponse) SetBody(v *ChangeDetailListOfBuyerResponseBody) *ChangeDetailListOfBuyerResponse {
	s.Body = v
	return s
}

func (s *ChangeDetailListOfBuyerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
