// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHistoryThreeElementsVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HistoryThreeElementsVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HistoryThreeElementsVerificationResponse
	GetStatusCode() *int32
	SetBody(v *HistoryThreeElementsVerificationResponseBody) *HistoryThreeElementsVerificationResponse
	GetBody() *HistoryThreeElementsVerificationResponseBody
}

type HistoryThreeElementsVerificationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HistoryThreeElementsVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HistoryThreeElementsVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s HistoryThreeElementsVerificationResponse) GoString() string {
	return s.String()
}

func (s *HistoryThreeElementsVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HistoryThreeElementsVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HistoryThreeElementsVerificationResponse) GetBody() *HistoryThreeElementsVerificationResponseBody {
	return s.Body
}

func (s *HistoryThreeElementsVerificationResponse) SetHeaders(v map[string]*string) *HistoryThreeElementsVerificationResponse {
	s.Headers = v
	return s
}

func (s *HistoryThreeElementsVerificationResponse) SetStatusCode(v int32) *HistoryThreeElementsVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *HistoryThreeElementsVerificationResponse) SetBody(v *HistoryThreeElementsVerificationResponseBody) *HistoryThreeElementsVerificationResponse {
	s.Body = v
	return s
}

func (s *HistoryThreeElementsVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
