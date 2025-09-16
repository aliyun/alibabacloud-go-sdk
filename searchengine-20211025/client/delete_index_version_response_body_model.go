// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIndexVersionResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteIndexVersionResponseBody
	GetResult() map[string]interface{}
}

type DeleteIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteIndexVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndexVersionResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteIndexVersionResponseBody) SetRequestId(v string) *DeleteIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexVersionResponseBody) SetResult(v map[string]interface{}) *DeleteIndexVersionResponseBody {
	s.Result = v
	return s
}

func (s *DeleteIndexVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
