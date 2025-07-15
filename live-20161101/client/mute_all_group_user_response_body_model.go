// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MuteAllGroupUserResponseBody
	GetRequestId() *string
	SetResult(v *MuteAllGroupUserResponseBodyResult) *MuteAllGroupUserResponseBody
	GetResult() *MuteAllGroupUserResponseBodyResult
}

type MuteAllGroupUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *MuteAllGroupUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s MuteAllGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MuteAllGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *MuteAllGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MuteAllGroupUserResponseBody) GetResult() *MuteAllGroupUserResponseBodyResult {
	return s.Result
}

func (s *MuteAllGroupUserResponseBody) SetRequestId(v string) *MuteAllGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteAllGroupUserResponseBody) SetResult(v *MuteAllGroupUserResponseBodyResult) *MuteAllGroupUserResponseBody {
	s.Result = v
	return s
}

func (s *MuteAllGroupUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type MuteAllGroupUserResponseBodyResult struct {
	// Indicates whether the mute was successful. Valid values:
	//
	// 	- true: The mute was successful.
	//
	// 	- false: The mute failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MuteAllGroupUserResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s MuteAllGroupUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MuteAllGroupUserResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *MuteAllGroupUserResponseBodyResult) SetSuccess(v bool) *MuteAllGroupUserResponseBodyResult {
	s.Success = &v
	return s
}

func (s *MuteAllGroupUserResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
