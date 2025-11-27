// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedeployRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedeployRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RedeployRCInstanceResponseBody) *RedeployRCInstanceResponse
	GetBody() *RedeployRCInstanceResponseBody
}

type RedeployRCInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedeployRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedeployRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RedeployRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *RedeployRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedeployRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedeployRCInstanceResponse) GetBody() *RedeployRCInstanceResponseBody {
	return s.Body
}

func (s *RedeployRCInstanceResponse) SetHeaders(v map[string]*string) *RedeployRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *RedeployRCInstanceResponse) SetStatusCode(v int32) *RedeployRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RedeployRCInstanceResponse) SetBody(v *RedeployRCInstanceResponseBody) *RedeployRCInstanceResponse {
	s.Body = v
	return s
}

func (s *RedeployRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
