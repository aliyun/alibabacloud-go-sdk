// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonCustInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCommonCustInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCommonCustInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryCommonCustInfoResponseBody) *QueryCommonCustInfoResponse
	GetBody() *QueryCommonCustInfoResponseBody
}

type QueryCommonCustInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCommonCustInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCommonCustInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonCustInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCommonCustInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCommonCustInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCommonCustInfoResponse) GetBody() *QueryCommonCustInfoResponseBody {
	return s.Body
}

func (s *QueryCommonCustInfoResponse) SetHeaders(v map[string]*string) *QueryCommonCustInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCommonCustInfoResponse) SetStatusCode(v int32) *QueryCommonCustInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCommonCustInfoResponse) SetBody(v *QueryCommonCustInfoResponseBody) *QueryCommonCustInfoResponse {
	s.Body = v
	return s
}

func (s *QueryCommonCustInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
