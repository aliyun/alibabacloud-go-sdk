// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySoarStrategySubscribeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySoarStrategySubscribeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySoarStrategySubscribeResponse
	GetStatusCode() *int32
	SetBody(v *ModifySoarStrategySubscribeResponseBody) *ModifySoarStrategySubscribeResponse
	GetBody() *ModifySoarStrategySubscribeResponseBody
}

type ModifySoarStrategySubscribeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySoarStrategySubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySoarStrategySubscribeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySoarStrategySubscribeResponse) GoString() string {
	return s.String()
}

func (s *ModifySoarStrategySubscribeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySoarStrategySubscribeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySoarStrategySubscribeResponse) GetBody() *ModifySoarStrategySubscribeResponseBody {
	return s.Body
}

func (s *ModifySoarStrategySubscribeResponse) SetHeaders(v map[string]*string) *ModifySoarStrategySubscribeResponse {
	s.Headers = v
	return s
}

func (s *ModifySoarStrategySubscribeResponse) SetStatusCode(v int32) *ModifySoarStrategySubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySoarStrategySubscribeResponse) SetBody(v *ModifySoarStrategySubscribeResponseBody) *ModifySoarStrategySubscribeResponse {
	s.Body = v
	return s
}

func (s *ModifySoarStrategySubscribeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
