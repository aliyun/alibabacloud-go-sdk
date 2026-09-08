// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventHouseWithTimeRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventHouseWithTimeRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventHouseWithTimeRangeResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventHouseWithTimeRangeResponseBody) *QueryEventHouseWithTimeRangeResponse
	GetBody() *QueryEventHouseWithTimeRangeResponseBody
}

type QueryEventHouseWithTimeRangeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventHouseWithTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventHouseWithTimeRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseWithTimeRangeResponse) GoString() string {
	return s.String()
}

func (s *QueryEventHouseWithTimeRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventHouseWithTimeRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventHouseWithTimeRangeResponse) GetBody() *QueryEventHouseWithTimeRangeResponseBody {
	return s.Body
}

func (s *QueryEventHouseWithTimeRangeResponse) SetHeaders(v map[string]*string) *QueryEventHouseWithTimeRangeResponse {
	s.Headers = v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponse) SetStatusCode(v int32) *QueryEventHouseWithTimeRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponse) SetBody(v *QueryEventHouseWithTimeRangeResponseBody) *QueryEventHouseWithTimeRangeResponse {
	s.Body = v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
