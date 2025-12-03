// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePushRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeletePushRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeletePushRuleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeletePushRuleResponseBody
	GetRequestId() *string
	SetResult(v *DeletePushRuleResponseBodyResult) *DeletePushRuleResponseBody
	GetResult() *DeletePushRuleResponseBodyResult
	SetSuccess(v string) *DeletePushRuleResponseBody
	GetSuccess() *string
}

type DeletePushRuleResponseBody struct {
	// example:
	//
	// InvalidParam.NotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeletePushRuleResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeletePushRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePushRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePushRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeletePushRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeletePushRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePushRuleResponseBody) GetResult() *DeletePushRuleResponseBodyResult {
	return s.Result
}

func (s *DeletePushRuleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeletePushRuleResponseBody) SetErrorCode(v string) *DeletePushRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePushRuleResponseBody) SetErrorMessage(v string) *DeletePushRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePushRuleResponseBody) SetRequestId(v string) *DeletePushRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePushRuleResponseBody) SetResult(v *DeletePushRuleResponseBodyResult) *DeletePushRuleResponseBody {
	s.Result = v
	return s
}

func (s *DeletePushRuleResponseBody) SetSuccess(v string) *DeletePushRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePushRuleResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeletePushRuleResponseBodyResult struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeletePushRuleResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeletePushRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeletePushRuleResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *DeletePushRuleResponseBodyResult) SetResult(v bool) *DeletePushRuleResponseBodyResult {
	s.Result = &v
	return s
}

func (s *DeletePushRuleResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
