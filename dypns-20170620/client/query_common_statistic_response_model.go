// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCommonStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCommonStatisticResponse
	GetStatusCode() *int32
	SetBody(v *QueryCommonStatisticResponseBody) *QueryCommonStatisticResponse
	GetBody() *QueryCommonStatisticResponseBody
}

type QueryCommonStatisticResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCommonStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCommonStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonStatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryCommonStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCommonStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCommonStatisticResponse) GetBody() *QueryCommonStatisticResponseBody {
	return s.Body
}

func (s *QueryCommonStatisticResponse) SetHeaders(v map[string]*string) *QueryCommonStatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryCommonStatisticResponse) SetStatusCode(v int32) *QueryCommonStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCommonStatisticResponse) SetBody(v *QueryCommonStatisticResponseBody) *QueryCommonStatisticResponse {
	s.Body = v
	return s
}

func (s *QueryCommonStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
