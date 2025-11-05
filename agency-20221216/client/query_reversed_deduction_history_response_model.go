// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReversedDeductionHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReversedDeductionHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReversedDeductionHistoryResponse
	GetStatusCode() *int32
	SetBody(v *QueryReversedDeductionHistoryResponseBody) *QueryReversedDeductionHistoryResponse
	GetBody() *QueryReversedDeductionHistoryResponseBody
}

type QueryReversedDeductionHistoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReversedDeductionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReversedDeductionHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReversedDeductionHistoryResponse) GoString() string {
	return s.String()
}

func (s *QueryReversedDeductionHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReversedDeductionHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReversedDeductionHistoryResponse) GetBody() *QueryReversedDeductionHistoryResponseBody {
	return s.Body
}

func (s *QueryReversedDeductionHistoryResponse) SetHeaders(v map[string]*string) *QueryReversedDeductionHistoryResponse {
	s.Headers = v
	return s
}

func (s *QueryReversedDeductionHistoryResponse) SetStatusCode(v int32) *QueryReversedDeductionHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponse) SetBody(v *QueryReversedDeductionHistoryResponseBody) *QueryReversedDeductionHistoryResponse {
	s.Body = v
	return s
}

func (s *QueryReversedDeductionHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
