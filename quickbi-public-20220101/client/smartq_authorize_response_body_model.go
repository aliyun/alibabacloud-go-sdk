// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqAuthorizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SmartqAuthorizeResponseBody
	GetRequestId() *string
	SetResult(v []*SmartqAuthorizeResponseBodyResult) *SmartqAuthorizeResponseBody
	GetResult() []*SmartqAuthorizeResponseBodyResult
	SetSuccess(v bool) *SmartqAuthorizeResponseBody
	GetSuccess() *bool
}

type SmartqAuthorizeResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 617277C****************ABA47E31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Array of failed user information.
	Result []*SmartqAuthorizeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. The value range is as follows:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SmartqAuthorizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmartqAuthorizeResponseBody) GoString() string {
	return s.String()
}

func (s *SmartqAuthorizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmartqAuthorizeResponseBody) GetResult() []*SmartqAuthorizeResponseBodyResult {
	return s.Result
}

func (s *SmartqAuthorizeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SmartqAuthorizeResponseBody) SetRequestId(v string) *SmartqAuthorizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmartqAuthorizeResponseBody) SetResult(v []*SmartqAuthorizeResponseBodyResult) *SmartqAuthorizeResponseBody {
	s.Result = v
	return s
}

func (s *SmartqAuthorizeResponseBody) SetSuccess(v bool) *SmartqAuthorizeResponseBody {
	s.Success = &v
	return s
}

func (s *SmartqAuthorizeResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SmartqAuthorizeResponseBodyResult struct {
	// Reason for failure.
	//
	// example:
	//
	// INVALID_FILE_FORMAT
	DetailMessage *string `json:"DetailMessage,omitempty" xml:"DetailMessage,omitempty"`
	// Q&A resource ID.
	//
	// example:
	//
	// 617277C****************ABA47E31
	LlmCube *string `json:"LlmCube,omitempty" xml:"LlmCube,omitempty"`
	// Analysis theme ID.
	//
	// example:
	//
	// 617277C****************ABA47E31
	LlmCubeTheme *string `json:"LlmCubeTheme,omitempty" xml:"LlmCubeTheme,omitempty"`
	// User ID.
	//
	// example:
	//
	// 617277C****************ABA47E31
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SmartqAuthorizeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SmartqAuthorizeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartqAuthorizeResponseBodyResult) GetDetailMessage() *string {
	return s.DetailMessage
}

func (s *SmartqAuthorizeResponseBodyResult) GetLlmCube() *string {
	return s.LlmCube
}

func (s *SmartqAuthorizeResponseBodyResult) GetLlmCubeTheme() *string {
	return s.LlmCubeTheme
}

func (s *SmartqAuthorizeResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *SmartqAuthorizeResponseBodyResult) SetDetailMessage(v string) *SmartqAuthorizeResponseBodyResult {
	s.DetailMessage = &v
	return s
}

func (s *SmartqAuthorizeResponseBodyResult) SetLlmCube(v string) *SmartqAuthorizeResponseBodyResult {
	s.LlmCube = &v
	return s
}

func (s *SmartqAuthorizeResponseBodyResult) SetLlmCubeTheme(v string) *SmartqAuthorizeResponseBodyResult {
	s.LlmCubeTheme = &v
	return s
}

func (s *SmartqAuthorizeResponseBodyResult) SetUserId(v string) *SmartqAuthorizeResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *SmartqAuthorizeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
