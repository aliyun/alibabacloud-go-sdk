// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLogstashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartLogstashResponseBody
	GetRequestId() *string
	SetResult(v *Logstash) *RestartLogstashResponseBody
	GetResult() *Logstash
}

type RestartLogstashResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Logstash `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RestartLogstashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *RestartLogstashResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartLogstashResponseBody) GetResult() *Logstash {
	return s.Result
}

func (s *RestartLogstashResponseBody) SetRequestId(v string) *RestartLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartLogstashResponseBody) SetResult(v *Logstash) *RestartLogstashResponseBody {
	s.Result = v
	return s
}

func (s *RestartLogstashResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
