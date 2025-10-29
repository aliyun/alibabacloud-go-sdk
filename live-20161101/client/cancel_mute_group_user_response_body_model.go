// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMuteGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelMuteGroupUserResponseBody
	GetRequestId() *string
	SetResult(v *CancelMuteGroupUserResponseBodyResult) *CancelMuteGroupUserResponseBody
	GetResult() *CancelMuteGroupUserResponseBodyResult
}

type CancelMuteGroupUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *CancelMuteGroupUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CancelMuteGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *CancelMuteGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelMuteGroupUserResponseBody) GetResult() *CancelMuteGroupUserResponseBodyResult {
	return s.Result
}

func (s *CancelMuteGroupUserResponseBody) SetRequestId(v string) *CancelMuteGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelMuteGroupUserResponseBody) SetResult(v *CancelMuteGroupUserResponseBodyResult) *CancelMuteGroupUserResponseBody {
	s.Result = v
	return s
}

func (s *CancelMuteGroupUserResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelMuteGroupUserResponseBodyResult struct {
	// Indicates whether the members are unmuted. Valid values:
	//
	// 	- true: The members are unmuted.
	//
	// 	- false: The members failed to be unmuted.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelMuteGroupUserResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteGroupUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CancelMuteGroupUserResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *CancelMuteGroupUserResponseBodyResult) SetSuccess(v bool) *CancelMuteGroupUserResponseBodyResult {
	s.Success = &v
	return s
}

func (s *CancelMuteGroupUserResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
