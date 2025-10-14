// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailListOfOrderNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDetailListOfOrderNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDetailListOfOrderNumResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDetailListOfOrderNumResponseBody) *ChangeDetailListOfOrderNumResponse
	GetBody() *ChangeDetailListOfOrderNumResponseBody
}

type ChangeDetailListOfOrderNumResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDetailListOfOrderNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponse) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDetailListOfOrderNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDetailListOfOrderNumResponse) GetBody() *ChangeDetailListOfOrderNumResponseBody {
	return s.Body
}

func (s *ChangeDetailListOfOrderNumResponse) SetHeaders(v map[string]*string) *ChangeDetailListOfOrderNumResponse {
	s.Headers = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponse) SetStatusCode(v int32) *ChangeDetailListOfOrderNumResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponse) SetBody(v *ChangeDetailListOfOrderNumResponseBody) *ChangeDetailListOfOrderNumResponse {
	s.Body = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
