// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLogstashDescriptionResponseBody
	GetRequestId() *string
	SetResult(v *UpdateLogstashDescriptionResponseBodyResult) *UpdateLogstashDescriptionResponseBody
	GetResult() *UpdateLogstashDescriptionResponseBodyResult
}

type UpdateLogstashDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned.
	Result *UpdateLogstashDescriptionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateLogstashDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLogstashDescriptionResponseBody) GetResult() *UpdateLogstashDescriptionResponseBodyResult {
	return s.Result
}

func (s *UpdateLogstashDescriptionResponseBody) SetRequestId(v string) *UpdateLogstashDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogstashDescriptionResponseBody) SetResult(v *UpdateLogstashDescriptionResponseBodyResult) *UpdateLogstashDescriptionResponseBody {
	s.Result = v
	return s
}

func (s *UpdateLogstashDescriptionResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLogstashDescriptionResponseBodyResult struct {
	// The name of the cluster.
	//
	// example:
	//
	// logstash_name
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateLogstashDescriptionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashDescriptionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateLogstashDescriptionResponseBodyResult) SetDescription(v string) *UpdateLogstashDescriptionResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateLogstashDescriptionResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
