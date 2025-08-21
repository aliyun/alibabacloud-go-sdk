// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogstashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLogstashResponseBody
	GetRequestId() *string
	SetResult(v *Logstash) *CreateLogstashResponseBody
	GetResult() *Logstash
}

type CreateLogstashResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE*****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *Logstash `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateLogstashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogstashResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogstashResponseBody) GetResult() *Logstash {
	return s.Result
}

func (s *CreateLogstashResponseBody) SetRequestId(v string) *CreateLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogstashResponseBody) SetResult(v *Logstash) *CreateLogstashResponseBody {
	s.Result = v
	return s
}

func (s *CreateLogstashResponseBody) Validate() error {
	return dara.Validate(s)
}
