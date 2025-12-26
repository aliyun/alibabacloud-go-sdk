// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddFaceRecordResponseBody
	GetCode() *string
	SetMessage(v string) *AddFaceRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddFaceRecordResponseBody
	GetRequestId() *string
	SetResult(v *AddFaceRecordResponseBodyResult) *AddFaceRecordResponseBody
	GetResult() *AddFaceRecordResponseBodyResult
}

type AddFaceRecordResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7A0D192A-CC0C-5DE5-A3B6-A14CF45508F2
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddFaceRecordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s AddFaceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFaceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddFaceRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddFaceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFaceRecordResponseBody) GetResult() *AddFaceRecordResponseBodyResult {
	return s.Result
}

func (s *AddFaceRecordResponseBody) SetCode(v string) *AddFaceRecordResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceRecordResponseBody) SetMessage(v string) *AddFaceRecordResponseBody {
	s.Message = &v
	return s
}

func (s *AddFaceRecordResponseBody) SetRequestId(v string) *AddFaceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceRecordResponseBody) SetResult(v *AddFaceRecordResponseBodyResult) *AddFaceRecordResponseBody {
	s.Result = v
	return s
}

func (s *AddFaceRecordResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFaceRecordResponseBodyResult struct {
	// example:
	//
	// {
	//
	// "faceQuality": "HIGH"
	//
	// }
	ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s AddFaceRecordResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddFaceRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddFaceRecordResponseBodyResult) GetExtFaceInfo() *string {
	return s.ExtFaceInfo
}

func (s *AddFaceRecordResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *AddFaceRecordResponseBodyResult) SetExtFaceInfo(v string) *AddFaceRecordResponseBodyResult {
	s.ExtFaceInfo = &v
	return s
}

func (s *AddFaceRecordResponseBodyResult) SetPassed(v string) *AddFaceRecordResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *AddFaceRecordResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
