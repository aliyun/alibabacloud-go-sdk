// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerRepositoryMirrorSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *TriggerRepositoryMirrorSyncResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *TriggerRepositoryMirrorSyncResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *TriggerRepositoryMirrorSyncResponseBody
	GetRequestId() *string
	SetResult(v *TriggerRepositoryMirrorSyncResponseBodyResult) *TriggerRepositoryMirrorSyncResponseBody
	GetResult() *TriggerRepositoryMirrorSyncResponseBodyResult
	SetSuccess(v bool) *TriggerRepositoryMirrorSyncResponseBody
	GetSuccess() *bool
}

type TriggerRepositoryMirrorSyncResponseBody struct {
	// example:
	//
	// SYSTEM_ILLEGAL_ARGUMENT_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 企业不存在
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 37294673-00CA-5B8B-914F-A8B35511E90A
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *TriggerRepositoryMirrorSyncResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TriggerRepositoryMirrorSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TriggerRepositoryMirrorSyncResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TriggerRepositoryMirrorSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerRepositoryMirrorSyncResponseBody) GetResult() *TriggerRepositoryMirrorSyncResponseBodyResult {
	return s.Result
}

func (s *TriggerRepositoryMirrorSyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetErrorCode(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetErrorMessage(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetRequestId(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetResult(v *TriggerRepositoryMirrorSyncResponseBodyResult) *TriggerRepositoryMirrorSyncResponseBody {
	s.Result = v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetSuccess(v bool) *TriggerRepositoryMirrorSyncResponseBody {
	s.Success = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TriggerRepositoryMirrorSyncResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TriggerRepositoryMirrorSyncResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *TriggerRepositoryMirrorSyncResponseBodyResult) SetResult(v bool) *TriggerRepositoryMirrorSyncResponseBodyResult {
	s.Result = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
