// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMuteAllGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelMuteAllGroupUserResponseBody
	GetRequestId() *string
	SetResult(v *CancelMuteAllGroupUserResponseBodyResult) *CancelMuteAllGroupUserResponseBody
	GetResult() *CancelMuteAllGroupUserResponseBodyResult
}

type CancelMuteAllGroupUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *CancelMuteAllGroupUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CancelMuteAllGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteAllGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *CancelMuteAllGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelMuteAllGroupUserResponseBody) GetResult() *CancelMuteAllGroupUserResponseBodyResult {
	return s.Result
}

func (s *CancelMuteAllGroupUserResponseBody) SetRequestId(v string) *CancelMuteAllGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelMuteAllGroupUserResponseBody) SetResult(v *CancelMuteAllGroupUserResponseBodyResult) *CancelMuteAllGroupUserResponseBody {
	s.Result = v
	return s
}

func (s *CancelMuteAllGroupUserResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelMuteAllGroupUserResponseBodyResult struct {
	// Indicates whether the cancellation was successful, with values:
	//
	// - true: Success.
	//
	// - false: Not successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelMuteAllGroupUserResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteAllGroupUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CancelMuteAllGroupUserResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *CancelMuteAllGroupUserResponseBodyResult) SetSuccess(v bool) *CancelMuteAllGroupUserResponseBodyResult {
	s.Success = &v
	return s
}

func (s *CancelMuteAllGroupUserResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
