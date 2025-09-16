// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIndexResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteIndexResponseBody
	GetResult() map[string]interface{}
}

type DeleteIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndexResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteIndexResponseBody) SetRequestId(v string) *DeleteIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexResponseBody) SetResult(v map[string]interface{}) *DeleteIndexResponseBody {
	s.Result = v
	return s
}

func (s *DeleteIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
