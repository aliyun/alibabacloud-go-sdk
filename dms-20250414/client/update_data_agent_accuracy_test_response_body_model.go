// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentAccuracyTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDataAgentAccuracyTestResponseBodyData) *UpdateDataAgentAccuracyTestResponseBody
	GetData() *UpdateDataAgentAccuracyTestResponseBodyData
	SetErrorCode(v string) *UpdateDataAgentAccuracyTestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataAgentAccuracyTestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDataAgentAccuracyTestResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateDataAgentAccuracyTestResponseBody
	GetSuccess() *string
}

type UpdateDataAgentAccuracyTestResponseBody struct {
	// The response struct.
	Data *UpdateDataAgentAccuracyTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-xxx-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataAgentAccuracyTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentAccuracyTestResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentAccuracyTestResponseBody) GetData() *UpdateDataAgentAccuracyTestResponseBodyData {
	return s.Data
}

func (s *UpdateDataAgentAccuracyTestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataAgentAccuracyTestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataAgentAccuracyTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataAgentAccuracyTestResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateDataAgentAccuracyTestResponseBody) SetData(v *UpdateDataAgentAccuracyTestResponseBodyData) *UpdateDataAgentAccuracyTestResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBody) SetErrorCode(v string) *UpdateDataAgentAccuracyTestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBody) SetErrorMessage(v string) *UpdateDataAgentAccuracyTestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBody) SetRequestId(v string) *UpdateDataAgentAccuracyTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBody) SetSuccess(v string) *UpdateDataAgentAccuracyTestResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataAgentAccuracyTestResponseBodyData struct {
	// The ID of the accuracy test item.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTest *string `json:"AccuracyTest,omitempty" xml:"AccuracyTest,omitempty"`
	// Agent Id
	//
	// example:
	//
	// cu0cs*******mf
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The data source. We recommend that you configure this in the custom agent.
	//
	// example:
	//
	// [{\\"DataSourceType\\":\\"database\\",\\"RegionId\\":\\"cn-hangzhou\\",\\"DmsInstanceId\\":\\"27xxxxx\\",\\"DmsDatabaseId\\":\\"752xxxxx\\",\\"Database\\":\\"employees\\",\\"Tables\\":[\\"employees\\",\\"salaries\\",\\"departments\\"]}]
	Dataset *string `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// The accuracy evaluation criteria. An empty value indicates the default criteria.
	//
	// example:
	//
	// null
	EvaluationPrompt *string `json:"EvaluationPrompt,omitempty" xml:"EvaluationPrompt,omitempty"`
	// The file ID.
	//
	// example:
	//
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The analysis mode.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDataAgentAccuracyTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentAccuracyTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) GetAccuracyTest() *string {
	return s.AccuracyTest
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) GetDataset() *string {
	return s.Dataset
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) GetEvaluationPrompt() *string {
	return s.EvaluationPrompt
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) GetMode() *int32 {
	return s.Mode
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) SetAccuracyTest(v string) *UpdateDataAgentAccuracyTestResponseBodyData {
	s.AccuracyTest = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) SetAgentId(v string) *UpdateDataAgentAccuracyTestResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) SetDataset(v string) *UpdateDataAgentAccuracyTestResponseBodyData {
	s.Dataset = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) SetEvaluationPrompt(v string) *UpdateDataAgentAccuracyTestResponseBodyData {
	s.EvaluationPrompt = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) SetFileId(v string) *UpdateDataAgentAccuracyTestResponseBodyData {
	s.FileId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) SetMode(v int32) *UpdateDataAgentAccuracyTestResponseBodyData {
	s.Mode = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) SetWorkspaceId(v string) *UpdateDataAgentAccuracyTestResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
