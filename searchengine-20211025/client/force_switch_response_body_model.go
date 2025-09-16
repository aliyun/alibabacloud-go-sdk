// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForceSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ForceSwitchResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ForceSwitchResponseBody
	GetResult() map[string]interface{}
}

type ForceSwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0B1FF998-BB8D-5182-BFC0-E471AA77095A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index information.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ForceSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForceSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ForceSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForceSwitchResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ForceSwitchResponseBody) SetRequestId(v string) *ForceSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForceSwitchResponseBody) SetResult(v map[string]interface{}) *ForceSwitchResponseBody {
	s.Result = v
	return s
}

func (s *ForceSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
