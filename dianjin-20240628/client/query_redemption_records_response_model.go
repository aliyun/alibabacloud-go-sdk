// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRedemptionRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRedemptionRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRedemptionRecordsResponse
	GetStatusCode() *int32
	SetBody(v *QueryRedemptionRecordsResponseBody) *QueryRedemptionRecordsResponse
	GetBody() *QueryRedemptionRecordsResponseBody
}

type QueryRedemptionRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRedemptionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRedemptionRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRedemptionRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryRedemptionRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRedemptionRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRedemptionRecordsResponse) GetBody() *QueryRedemptionRecordsResponseBody {
	return s.Body
}

func (s *QueryRedemptionRecordsResponse) SetHeaders(v map[string]*string) *QueryRedemptionRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryRedemptionRecordsResponse) SetStatusCode(v int32) *QueryRedemptionRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRedemptionRecordsResponse) SetBody(v *QueryRedemptionRecordsResponseBody) *QueryRedemptionRecordsResponse {
	s.Body = v
	return s
}

func (s *QueryRedemptionRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
