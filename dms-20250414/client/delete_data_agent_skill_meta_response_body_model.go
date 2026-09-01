// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentSkillMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDataAgentSkillMetaResponseBodyData) *DeleteDataAgentSkillMetaResponseBody
	GetData() *DeleteDataAgentSkillMetaResponseBodyData
	SetErrorCode(v string) *DeleteDataAgentSkillMetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataAgentSkillMetaResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataAgentSkillMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataAgentSkillMetaResponseBody
	GetSuccess() *bool
}

type DeleteDataAgentSkillMetaResponseBody struct {
	// The response struct.
	Data *DeleteDataAgentSkillMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentSkillMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentSkillMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentSkillMetaResponseBody) GetData() *DeleteDataAgentSkillMetaResponseBodyData {
	return s.Data
}

func (s *DeleteDataAgentSkillMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataAgentSkillMetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataAgentSkillMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAgentSkillMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentSkillMetaResponseBody) SetData(v *DeleteDataAgentSkillMetaResponseBodyData) *DeleteDataAgentSkillMetaResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataAgentSkillMetaResponseBody) SetErrorCode(v string) *DeleteDataAgentSkillMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataAgentSkillMetaResponseBody) SetErrorMessage(v string) *DeleteDataAgentSkillMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataAgentSkillMetaResponseBody) SetRequestId(v string) *DeleteDataAgentSkillMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAgentSkillMetaResponseBody) SetSuccess(v bool) *DeleteDataAgentSkillMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentSkillMetaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDataAgentSkillMetaResponseBodyData struct {
	// The skill ID.
	//
	// example:
	//
	// ski-04pomiln*************j0
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - **true**: The operation was successful.
	//
	// - **false**: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentSkillMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentSkillMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentSkillMetaResponseBodyData) GetSkillId() *string {
	return s.SkillId
}

func (s *DeleteDataAgentSkillMetaResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentSkillMetaResponseBodyData) SetSkillId(v string) *DeleteDataAgentSkillMetaResponseBodyData {
	s.SkillId = &v
	return s
}

func (s *DeleteDataAgentSkillMetaResponseBodyData) SetSuccess(v bool) *DeleteDataAgentSkillMetaResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentSkillMetaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
