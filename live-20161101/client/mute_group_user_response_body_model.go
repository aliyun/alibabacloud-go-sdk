// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MuteGroupUserResponseBody
	GetRequestId() *string
	SetResult(v *MuteGroupUserResponseBodyResult) *MuteGroupUserResponseBody
	GetResult() *MuteGroupUserResponseBodyResult
}

type MuteGroupUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *MuteGroupUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s MuteGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MuteGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *MuteGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MuteGroupUserResponseBody) GetResult() *MuteGroupUserResponseBodyResult {
	return s.Result
}

func (s *MuteGroupUserResponseBody) SetRequestId(v string) *MuteGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteGroupUserResponseBody) SetResult(v *MuteGroupUserResponseBodyResult) *MuteGroupUserResponseBody {
	s.Result = v
	return s
}

func (s *MuteGroupUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type MuteGroupUserResponseBodyResult struct {
	// Indicates whether the mute is successful. Valid values:
	//
	// 	- true: The mute is successful.
	//
	// 	- false: The mute failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MuteGroupUserResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s MuteGroupUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MuteGroupUserResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *MuteGroupUserResponseBodyResult) SetSuccess(v bool) *MuteGroupUserResponseBodyResult {
	s.Success = &v
	return s
}

func (s *MuteGroupUserResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
