// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPromptVersionResponseBodyData) *GetPromptVersionResponseBody
	GetData() *GetPromptVersionResponseBodyData
	SetRequestId(v string) *GetPromptVersionResponseBody
	GetRequestId() *string
}

type GetPromptVersionResponseBody struct {
	Data      *GetPromptVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPromptVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPromptVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPromptVersionResponseBody) GetData() *GetPromptVersionResponseBodyData {
	return s.Data
}

func (s *GetPromptVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPromptVersionResponseBody) SetData(v *GetPromptVersionResponseBodyData) *GetPromptVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetPromptVersionResponseBody) SetRequestId(v string) *GetPromptVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPromptVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPromptVersionResponseBodyData struct {
	CommitMsg   *string                                      `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	GmtModified *int64                                       `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Md5         *string                                      `json:"Md5,omitempty" xml:"Md5,omitempty"`
	PromptKey   *string                                      `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
	SrcUser     *string                                      `json:"SrcUser,omitempty" xml:"SrcUser,omitempty"`
	Status      *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Template    *string                                      `json:"Template,omitempty" xml:"Template,omitempty"`
	Variables   []*GetPromptVersionResponseBodyDataVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
	Version     *string                                      `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetPromptVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPromptVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPromptVersionResponseBodyData) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *GetPromptVersionResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetPromptVersionResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *GetPromptVersionResponseBodyData) GetPromptKey() *string {
	return s.PromptKey
}

func (s *GetPromptVersionResponseBodyData) GetSrcUser() *string {
	return s.SrcUser
}

func (s *GetPromptVersionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetPromptVersionResponseBodyData) GetTemplate() *string {
	return s.Template
}

func (s *GetPromptVersionResponseBodyData) GetVariables() []*GetPromptVersionResponseBodyDataVariables {
	return s.Variables
}

func (s *GetPromptVersionResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetPromptVersionResponseBodyData) SetCommitMsg(v string) *GetPromptVersionResponseBodyData {
	s.CommitMsg = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetGmtModified(v int64) *GetPromptVersionResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetMd5(v string) *GetPromptVersionResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetPromptKey(v string) *GetPromptVersionResponseBodyData {
	s.PromptKey = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetSrcUser(v string) *GetPromptVersionResponseBodyData {
	s.SrcUser = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetStatus(v string) *GetPromptVersionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetTemplate(v string) *GetPromptVersionResponseBodyData {
	s.Template = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetVariables(v []*GetPromptVersionResponseBodyDataVariables) *GetPromptVersionResponseBodyData {
	s.Variables = v
	return s
}

func (s *GetPromptVersionResponseBodyData) SetVersion(v string) *GetPromptVersionResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetPromptVersionResponseBodyData) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPromptVersionResponseBodyDataVariables struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPromptVersionResponseBodyDataVariables) String() string {
	return dara.Prettify(s)
}

func (s GetPromptVersionResponseBodyDataVariables) GoString() string {
	return s.String()
}

func (s *GetPromptVersionResponseBodyDataVariables) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetPromptVersionResponseBodyDataVariables) GetDescription() *string {
	return s.Description
}

func (s *GetPromptVersionResponseBodyDataVariables) GetName() *string {
	return s.Name
}

func (s *GetPromptVersionResponseBodyDataVariables) SetDefaultValue(v string) *GetPromptVersionResponseBodyDataVariables {
	s.DefaultValue = &v
	return s
}

func (s *GetPromptVersionResponseBodyDataVariables) SetDescription(v string) *GetPromptVersionResponseBodyDataVariables {
	s.Description = &v
	return s
}

func (s *GetPromptVersionResponseBodyDataVariables) SetName(v string) *GetPromptVersionResponseBodyDataVariables {
	s.Name = &v
	return s
}

func (s *GetPromptVersionResponseBodyDataVariables) Validate() error {
	return dara.Validate(s)
}
