// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotUnsubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRobotUnsubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRobotUnsubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *QueryRobotUnsubscriptionResponseBody) *QueryRobotUnsubscriptionResponse
	GetBody() *QueryRobotUnsubscriptionResponseBody
}

type QueryRobotUnsubscriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotUnsubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRobotUnsubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotUnsubscriptionResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotUnsubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRobotUnsubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRobotUnsubscriptionResponse) GetBody() *QueryRobotUnsubscriptionResponseBody {
	return s.Body
}

func (s *QueryRobotUnsubscriptionResponse) SetHeaders(v map[string]*string) *QueryRobotUnsubscriptionResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotUnsubscriptionResponse) SetStatusCode(v int32) *QueryRobotUnsubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotUnsubscriptionResponse) SetBody(v *QueryRobotUnsubscriptionResponseBody) *QueryRobotUnsubscriptionResponse {
	s.Body = v
	return s
}

func (s *QueryRobotUnsubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
