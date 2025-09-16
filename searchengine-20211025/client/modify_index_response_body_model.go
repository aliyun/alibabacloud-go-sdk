// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIndexResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *ModifyIndexResponseBody
	GetResult() interface{}
}

type ModifyIndexResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIndexResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *ModifyIndexResponseBody) SetRequestId(v string) *ModifyIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexResponseBody) SetResult(v interface{}) *ModifyIndexResponseBody {
	s.Result = v
	return s
}

func (s *ModifyIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
