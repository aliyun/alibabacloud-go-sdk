// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSenderStatisticsByTagNameAndBatchIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SenderStatisticsByTagNameAndBatchIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SenderStatisticsByTagNameAndBatchIDResponse
	GetStatusCode() *int32
	SetBody(v *SenderStatisticsByTagNameAndBatchIDResponseBody) *SenderStatisticsByTagNameAndBatchIDResponse
	GetBody() *SenderStatisticsByTagNameAndBatchIDResponseBody
}

type SenderStatisticsByTagNameAndBatchIDResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SenderStatisticsByTagNameAndBatchIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponse) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponse) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) GetBody() *SenderStatisticsByTagNameAndBatchIDResponseBody {
	return s.Body
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) SetHeaders(v map[string]*string) *SenderStatisticsByTagNameAndBatchIDResponse {
	s.Headers = v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) SetStatusCode(v int32) *SenderStatisticsByTagNameAndBatchIDResponse {
	s.StatusCode = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) SetBody(v *SenderStatisticsByTagNameAndBatchIDResponseBody) *SenderStatisticsByTagNameAndBatchIDResponse {
	s.Body = v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
