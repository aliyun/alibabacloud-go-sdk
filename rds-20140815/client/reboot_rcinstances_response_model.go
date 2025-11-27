// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootRCInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootRCInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RebootRCInstancesResponseBody) *RebootRCInstancesResponse
	GetBody() *RebootRCInstancesResponseBody
}

type RebootRCInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootRCInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootRCInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstancesResponse) GoString() string {
	return s.String()
}

func (s *RebootRCInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootRCInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootRCInstancesResponse) GetBody() *RebootRCInstancesResponseBody {
	return s.Body
}

func (s *RebootRCInstancesResponse) SetHeaders(v map[string]*string) *RebootRCInstancesResponse {
	s.Headers = v
	return s
}

func (s *RebootRCInstancesResponse) SetStatusCode(v int32) *RebootRCInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootRCInstancesResponse) SetBody(v *RebootRCInstancesResponseBody) *RebootRCInstancesResponse {
	s.Body = v
	return s
}

func (s *RebootRCInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
