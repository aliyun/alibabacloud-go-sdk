// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckCreateOrderForDeleteDBNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailures(v *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures) *PreCheckCreateOrderForDeleteDBNodesResponseBody
	GetFailures() *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures
	SetPreCheckResult(v bool) *PreCheckCreateOrderForDeleteDBNodesResponseBody
	GetPreCheckResult() *bool
	SetRequestId(v string) *PreCheckCreateOrderForDeleteDBNodesResponseBody
	GetRequestId() *string
}

type PreCheckCreateOrderForDeleteDBNodesResponseBody struct {
	// The information about the failed order.
	Failures *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures `json:"Failures,omitempty" xml:"Failures,omitempty" type:"Struct"`
	// The precheck result.
	//
	// example:
	//
	// True
	PreCheckResult *bool `json:"PreCheckResult,omitempty" xml:"PreCheckResult,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B415BC6-FE84-5323-A255-42CF330DB99C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreCheckCreateOrderForDeleteDBNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateOrderForDeleteDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBody) GetFailures() *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures {
	return s.Failures
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBody) GetPreCheckResult() *bool {
	return s.PreCheckResult
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBody) SetFailures(v *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures) *PreCheckCreateOrderForDeleteDBNodesResponseBody {
	s.Failures = v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBody) SetPreCheckResult(v bool) *PreCheckCreateOrderForDeleteDBNodesResponseBody {
	s.PreCheckResult = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBody) SetRequestId(v string) *PreCheckCreateOrderForDeleteDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBody) Validate() error {
	if s.Failures != nil {
		if err := s.Failures.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures struct {
	Failures []*PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures `json:"Failures,omitempty" xml:"Failures,omitempty" type:"Repeated"`
}

func (s PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures) GoString() string {
	return s.String()
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures) GetFailures() []*PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures {
	return s.Failures
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures) SetFailures(v []*PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures {
	s.Failures = v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailures) Validate() error {
	if s.Failures != nil {
		for _, item := range s.Failures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures struct {
	// The response code. Valid values:
	//
	// 	- **200**: success
	//
	// 	- **400**: client error
	//
	// 	- **401**: identity authentication failed
	//
	// 	- **404**: requested page not found
	//
	// 	- **500**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) GoString() string {
	return s.String()
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) GetCode() *string {
	return s.Code
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) GetMessage() *string {
	return s.Message
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) SetCode(v string) *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures {
	s.Code = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) SetMessage(v string) *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures {
	s.Message = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesResponseBodyFailuresFailures) Validate() error {
	return dara.Validate(s)
}
