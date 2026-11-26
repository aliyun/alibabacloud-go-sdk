// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestHttpEventConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *TestHttpEventConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *TestHttpEventConfigRequest
	GetJsonStr() *string
}

type TestHttpEventConfigRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {
	//
	//     "callbackUrl": "http://123213.com",
	//
	//     "httpMethod": "GET"
	//
	// }
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s TestHttpEventConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s TestHttpEventConfigRequest) GoString() string {
	return s.String()
}

func (s *TestHttpEventConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *TestHttpEventConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *TestHttpEventConfigRequest) SetBaseMeAgentId(v int64) *TestHttpEventConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *TestHttpEventConfigRequest) SetJsonStr(v string) *TestHttpEventConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *TestHttpEventConfigRequest) Validate() error {
	return dara.Validate(s)
}
