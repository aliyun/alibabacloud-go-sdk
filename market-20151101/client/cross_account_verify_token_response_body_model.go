// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCrossAccountVerifyTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CrossAccountVerifyTokenResponseBody
	GetCode() *string
	SetMessage(v string) *CrossAccountVerifyTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CrossAccountVerifyTokenResponseBody
	GetRequestId() *string
	SetResult(v *CrossAccountVerifyTokenResponseBodyResult) *CrossAccountVerifyTokenResponseBody
	GetResult() *CrossAccountVerifyTokenResponseBodyResult
	SetSuccess(v bool) *CrossAccountVerifyTokenResponseBody
	GetSuccess() *bool
}

type CrossAccountVerifyTokenResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// C19D103F-EA2D-50A5-8441-0267CE9FBA56
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CrossAccountVerifyTokenResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CrossAccountVerifyTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CrossAccountVerifyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CrossAccountVerifyTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CrossAccountVerifyTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CrossAccountVerifyTokenResponseBody) GetResult() *CrossAccountVerifyTokenResponseBodyResult {
	return s.Result
}

func (s *CrossAccountVerifyTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CrossAccountVerifyTokenResponseBody) SetCode(v string) *CrossAccountVerifyTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetMessage(v string) *CrossAccountVerifyTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetRequestId(v string) *CrossAccountVerifyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetResult(v *CrossAccountVerifyTokenResponseBodyResult) *CrossAccountVerifyTokenResponseBody {
	s.Result = v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetSuccess(v bool) *CrossAccountVerifyTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type CrossAccountVerifyTokenResponseBodyResult struct {
	AuthRoles []*string `json:"AuthRoles,omitempty" xml:"AuthRoles,omitempty" type:"Repeated"`
	// example:
	//
	// 1676974108866
	AuthTime *int64 `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	// example:
	//
	// marketplace_wangxiao_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1744526877246715
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CrossAccountVerifyTokenResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CrossAccountVerifyTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenResponseBodyResult) GetAuthRoles() []*string {
	return s.AuthRoles
}

func (s *CrossAccountVerifyTokenResponseBodyResult) GetAuthTime() *int64 {
	return s.AuthTime
}

func (s *CrossAccountVerifyTokenResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CrossAccountVerifyTokenResponseBodyResult) GetUid() *string {
	return s.Uid
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetAuthRoles(v []*string) *CrossAccountVerifyTokenResponseBodyResult {
	s.AuthRoles = v
	return s
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetAuthTime(v int64) *CrossAccountVerifyTokenResponseBodyResult {
	s.AuthTime = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetName(v string) *CrossAccountVerifyTokenResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetUid(v string) *CrossAccountVerifyTokenResponseBodyResult {
	s.Uid = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
