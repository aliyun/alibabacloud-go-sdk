// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIndexVersionResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyIndexVersionResponseBody
	GetResult() map[string]interface{}
}

type ModifyIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIndexVersionResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyIndexVersionResponseBody) SetRequestId(v string) *ModifyIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexVersionResponseBody) SetResult(v map[string]interface{}) *ModifyIndexVersionResponseBody {
	s.Result = v
	return s
}

func (s *ModifyIndexVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
