// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAliasResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreateAliasResponseBody
	GetResult() map[string]interface{}
}

type CreateAliasResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAliasResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreateAliasResponseBody) SetRequestId(v string) *CreateAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAliasResponseBody) SetResult(v map[string]interface{}) *CreateAliasResponseBody {
	s.Result = v
	return s
}

func (s *CreateAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
