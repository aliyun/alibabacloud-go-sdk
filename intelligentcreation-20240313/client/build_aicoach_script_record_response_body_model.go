// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildAICoachScriptRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BuildAICoachScriptRecordResponseBodyData) *BuildAICoachScriptRecordResponseBody
	GetData() *BuildAICoachScriptRecordResponseBodyData
	SetErrorCode(v string) *BuildAICoachScriptRecordResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BuildAICoachScriptRecordResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *BuildAICoachScriptRecordResponseBody
	GetRequestId() *string
	SetScriptRecordId(v string) *BuildAICoachScriptRecordResponseBody
	GetScriptRecordId() *string
	SetSuccess(v bool) *BuildAICoachScriptRecordResponseBody
	GetSuccess() *bool
}

type BuildAICoachScriptRecordResponseBody struct {
	Data *BuildAICoachScriptRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// PARAM_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Deduct task already success,Please do not resubmit.token \\"369e8f2c-d283-424a-96c4-c83efe08c89e\\"
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4830493A-728F-5F19-BBCC-1443292E9C49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	ScriptRecordId *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BuildAICoachScriptRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuildAICoachScriptRecordResponseBody) GoString() string {
	return s.String()
}

func (s *BuildAICoachScriptRecordResponseBody) GetData() *BuildAICoachScriptRecordResponseBodyData {
	return s.Data
}

func (s *BuildAICoachScriptRecordResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BuildAICoachScriptRecordResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BuildAICoachScriptRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuildAICoachScriptRecordResponseBody) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *BuildAICoachScriptRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuildAICoachScriptRecordResponseBody) SetData(v *BuildAICoachScriptRecordResponseBodyData) *BuildAICoachScriptRecordResponseBody {
	s.Data = v
	return s
}

func (s *BuildAICoachScriptRecordResponseBody) SetErrorCode(v string) *BuildAICoachScriptRecordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BuildAICoachScriptRecordResponseBody) SetErrorMessage(v string) *BuildAICoachScriptRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BuildAICoachScriptRecordResponseBody) SetRequestId(v string) *BuildAICoachScriptRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildAICoachScriptRecordResponseBody) SetScriptRecordId(v string) *BuildAICoachScriptRecordResponseBody {
	s.ScriptRecordId = &v
	return s
}

func (s *BuildAICoachScriptRecordResponseBody) SetSuccess(v bool) *BuildAICoachScriptRecordResponseBody {
	s.Success = &v
	return s
}

func (s *BuildAICoachScriptRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BuildAICoachScriptRecordResponseBodyData struct {
	// example:
	//
	// 1234567
	ScriptId *string `json:"scriptId,omitempty" xml:"scriptId,omitempty"`
}

func (s BuildAICoachScriptRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BuildAICoachScriptRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *BuildAICoachScriptRecordResponseBodyData) GetScriptId() *string {
	return s.ScriptId
}

func (s *BuildAICoachScriptRecordResponseBodyData) SetScriptId(v string) *BuildAICoachScriptRecordResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *BuildAICoachScriptRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
