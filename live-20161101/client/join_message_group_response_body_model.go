// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *JoinMessageGroupResponseBody
	GetRequestId() *string
	SetResult(v *JoinMessageGroupResponseBodyResult) *JoinMessageGroupResponseBody
	GetResult() *JoinMessageGroupResponseBodyResult
}

type JoinMessageGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *JoinMessageGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s JoinMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinMessageGroupResponseBody) GetResult() *JoinMessageGroupResponseBodyResult {
	return s.Result
}

func (s *JoinMessageGroupResponseBody) SetRequestId(v string) *JoinMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinMessageGroupResponseBody) SetResult(v *JoinMessageGroupResponseBodyResult) *JoinMessageGroupResponseBody {
	s.Result = v
	return s
}

func (s *JoinMessageGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type JoinMessageGroupResponseBodyResult struct {
	// Indicates whether the users successfully joined the message group. Valid values:
	//
	// 	- true: The users successfully joined the message group.
	//
	// 	- false: The users failed to join the message group.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinMessageGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s JoinMessageGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *JoinMessageGroupResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *JoinMessageGroupResponseBodyResult) SetSuccess(v bool) *JoinMessageGroupResponseBodyResult {
	s.Success = &v
	return s
}

func (s *JoinMessageGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
