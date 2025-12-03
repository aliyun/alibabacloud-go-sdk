// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteTagResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteTagResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteTagResponseBody
	GetRequestId() *string
	SetResult(v *DeleteTagResponseBodyResult) *DeleteTagResponseBody
	GetResult() *DeleteTagResponseBodyResult
	SetSuccess(v bool) *DeleteTagResponseBody
	GetSuccess() *bool
}

type DeleteTagResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 60945D4F-CF6D-5CFF-89ED-1D1F6657032C
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTagResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTagResponseBody) GetResult() *DeleteTagResponseBodyResult {
	return s.Result
}

func (s *DeleteTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTagResponseBody) SetErrorCode(v string) *DeleteTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTagResponseBody) SetErrorMessage(v string) *DeleteTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTagResponseBody) SetRequestId(v string) *DeleteTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagResponseBody) SetResult(v *DeleteTagResponseBodyResult) *DeleteTagResponseBody {
	s.Result = v
	return s
}

func (s *DeleteTagResponseBody) SetSuccess(v bool) *DeleteTagResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTagResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteTagResponseBodyResult struct {
	// example:
	//
	// v1.0
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s DeleteTagResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteTagResponseBodyResult) GetTagName() *string {
	return s.TagName
}

func (s *DeleteTagResponseBodyResult) SetTagName(v string) *DeleteTagResponseBodyResult {
	s.TagName = &v
	return s
}

func (s *DeleteTagResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
