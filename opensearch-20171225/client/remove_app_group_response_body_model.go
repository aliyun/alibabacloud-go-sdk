// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveAppGroupResponseBody
	GetRequestId() *string
	SetResult(v []*int32) *RemoveAppGroupResponseBody
	GetResult() []*int32
}

type RemoveAppGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3AA29D02-54F3-8569-F71A-90E1B7BE4E7E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAppGroupResponseBody) GetResult() []*int32 {
	return s.Result
}

func (s *RemoveAppGroupResponseBody) SetRequestId(v string) *RemoveAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAppGroupResponseBody) SetResult(v []*int32) *RemoveAppGroupResponseBody {
	s.Result = v
	return s
}

func (s *RemoveAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
