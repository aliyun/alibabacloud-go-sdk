// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAliasResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteAliasResponseBody
	GetResult() map[string]interface{}
}

type DeleteAliasResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAliasResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAliasResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteAliasResponseBody) SetRequestId(v string) *DeleteAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAliasResponseBody) SetResult(v map[string]interface{}) *DeleteAliasResponseBody {
	s.Result = v
	return s
}

func (s *DeleteAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
