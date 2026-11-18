// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckVideoCharacterReplaceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PrecheckVideoCharacterReplaceJobResponseBody
	GetRequestId() *string
	SetResult(v []*PrecheckVideoCharacterReplaceJobResponseBodyResult) *PrecheckVideoCharacterReplaceJobResponseBody
	GetResult() []*PrecheckVideoCharacterReplaceJobResponseBodyResult
	SetStatus(v string) *PrecheckVideoCharacterReplaceJobResponseBody
	GetStatus() *string
}

type PrecheckVideoCharacterReplaceJobResponseBody struct {
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of items that failed the validation. This is an empty array if the validation passes.
	Result []*PrecheckVideoCharacterReplaceJobResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The pre-validation status. A value of Success indicates that the validation passed. A value of Failed indicates that one or more items failed the validation.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PrecheckVideoCharacterReplaceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrecheckVideoCharacterReplaceJobResponseBody) GoString() string {
	return s.String()
}

func (s *PrecheckVideoCharacterReplaceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrecheckVideoCharacterReplaceJobResponseBody) GetResult() []*PrecheckVideoCharacterReplaceJobResponseBodyResult {
	return s.Result
}

func (s *PrecheckVideoCharacterReplaceJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *PrecheckVideoCharacterReplaceJobResponseBody) SetRequestId(v string) *PrecheckVideoCharacterReplaceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponseBody) SetResult(v []*PrecheckVideoCharacterReplaceJobResponseBodyResult) *PrecheckVideoCharacterReplaceJobResponseBody {
	s.Result = v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponseBody) SetStatus(v string) *PrecheckVideoCharacterReplaceJobResponseBody {
	s.Status = &v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponseBody) Validate() error {
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

type PrecheckVideoCharacterReplaceJobResponseBodyResult struct {
	// The validation error code for the input item. For example, ImageCheckFailed indicates that the image validation failed, and AudioCheckFailed indicates that the audio validation failed.
	//
	// example:
	//
	// ImageCheckFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The name of the input parameter that failed the validation, which matches the corresponding field name in AppParams.
	//
	// example:
	//
	// LoadImage.1.TargetImage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s PrecheckVideoCharacterReplaceJobResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s PrecheckVideoCharacterReplaceJobResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PrecheckVideoCharacterReplaceJobResponseBodyResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PrecheckVideoCharacterReplaceJobResponseBodyResult) GetKey() *string {
	return s.Key
}

func (s *PrecheckVideoCharacterReplaceJobResponseBodyResult) SetErrorCode(v string) *PrecheckVideoCharacterReplaceJobResponseBodyResult {
	s.ErrorCode = &v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponseBodyResult) SetKey(v string) *PrecheckVideoCharacterReplaceJobResponseBodyResult {
	s.Key = &v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
