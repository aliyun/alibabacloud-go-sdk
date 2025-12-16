// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveAppResponseBody
	GetRequestId() *string
	SetResult(v []*int32) *RemoveAppResponseBody
	GetResult() []*int32
}

type RemoveAppResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 33477D76-C380-2D84-A4AD-043F52876CB1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAppResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAppResponseBody) GetResult() []*int32 {
	return s.Result
}

func (s *RemoveAppResponseBody) SetRequestId(v string) *RemoveAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAppResponseBody) SetResult(v []*int32) *RemoveAppResponseBody {
	s.Result = v
	return s
}

func (s *RemoveAppResponseBody) Validate() error {
	return dara.Validate(s)
}
