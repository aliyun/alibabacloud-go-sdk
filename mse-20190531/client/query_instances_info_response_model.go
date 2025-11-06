// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstancesInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstancesInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstancesInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstancesInfoResponseBody) *QueryInstancesInfoResponse
	GetBody() *QueryInstancesInfoResponseBody
}

type QueryInstancesInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstancesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstancesInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancesInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryInstancesInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstancesInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstancesInfoResponse) GetBody() *QueryInstancesInfoResponseBody {
	return s.Body
}

func (s *QueryInstancesInfoResponse) SetHeaders(v map[string]*string) *QueryInstancesInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryInstancesInfoResponse) SetStatusCode(v int32) *QueryInstancesInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstancesInfoResponse) SetBody(v *QueryInstancesInfoResponseBody) *QueryInstancesInfoResponse {
	s.Body = v
	return s
}

func (s *QueryInstancesInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
