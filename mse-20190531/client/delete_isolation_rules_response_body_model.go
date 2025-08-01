// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIsolationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteIsolationRulesResponseBody
	GetCode() *int32
	SetData(v []*int64) *DeleteIsolationRulesResponseBody
	GetData() []*int64
	SetHttpStatusCode(v int32) *DeleteIsolationRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteIsolationRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIsolationRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIsolationRulesResponseBody
	GetSuccess() *bool
}

type DeleteIsolationRulesResponseBody struct {
	// example:
	//
	// 200
	Code *int32   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIsolationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIsolationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteIsolationRulesResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *DeleteIsolationRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteIsolationRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIsolationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIsolationRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIsolationRulesResponseBody) SetCode(v int32) *DeleteIsolationRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIsolationRulesResponseBody) SetData(v []*int64) *DeleteIsolationRulesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIsolationRulesResponseBody) SetHttpStatusCode(v int32) *DeleteIsolationRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteIsolationRulesResponseBody) SetMessage(v string) *DeleteIsolationRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIsolationRulesResponseBody) SetRequestId(v string) *DeleteIsolationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIsolationRulesResponseBody) SetSuccess(v bool) *DeleteIsolationRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIsolationRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
