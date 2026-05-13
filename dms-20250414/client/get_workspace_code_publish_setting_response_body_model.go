// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceCodePublishSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkspaceCodePublishSettingResponseBody
	GetCode() *string
	SetData(v *GetWorkspaceCodePublishSettingResponseBodyData) *GetWorkspaceCodePublishSettingResponseBody
	GetData() *GetWorkspaceCodePublishSettingResponseBodyData
	SetErrorCode(v string) *GetWorkspaceCodePublishSettingResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetWorkspaceCodePublishSettingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkspaceCodePublishSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspaceCodePublishSettingResponseBody
	GetRequestId() *string
}

type GetWorkspaceCodePublishSettingResponseBody struct {
	// example:
	//
	// 200
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetWorkspaceCodePublishSettingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// This record is being collected, please wait for a moment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWorkspaceCodePublishSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodePublishSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodePublishSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkspaceCodePublishSettingResponseBody) GetData() *GetWorkspaceCodePublishSettingResponseBodyData {
	return s.Data
}

func (s *GetWorkspaceCodePublishSettingResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkspaceCodePublishSettingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkspaceCodePublishSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceCodePublishSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceCodePublishSettingResponseBody) SetCode(v string) *GetWorkspaceCodePublishSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBody) SetData(v *GetWorkspaceCodePublishSettingResponseBodyData) *GetWorkspaceCodePublishSettingResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBody) SetErrorCode(v string) *GetWorkspaceCodePublishSettingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBody) SetHttpStatusCode(v int32) *GetWorkspaceCodePublishSettingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBody) SetMessage(v string) *GetWorkspaceCodePublishSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBody) SetRequestId(v string) *GetWorkspaceCodePublishSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceCodePublishSettingResponseBodyData struct {
	Exclude []*string `json:"Exclude,omitempty" xml:"Exclude,omitempty" type:"Repeated"`
	// example:
	//
	// false
	LockRepoBranch *bool                                                  `json:"LockRepoBranch,omitempty" xml:"LockRepoBranch,omitempty"`
	Repos          []*GetWorkspaceCodePublishSettingResponseBodyDataRepos `json:"Repos,omitempty" xml:"Repos,omitempty" type:"Repeated"`
}

func (s GetWorkspaceCodePublishSettingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodePublishSettingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodePublishSettingResponseBodyData) GetExclude() []*string {
	return s.Exclude
}

func (s *GetWorkspaceCodePublishSettingResponseBodyData) GetLockRepoBranch() *bool {
	return s.LockRepoBranch
}

func (s *GetWorkspaceCodePublishSettingResponseBodyData) GetRepos() []*GetWorkspaceCodePublishSettingResponseBodyDataRepos {
	return s.Repos
}

func (s *GetWorkspaceCodePublishSettingResponseBodyData) SetExclude(v []*string) *GetWorkspaceCodePublishSettingResponseBodyData {
	s.Exclude = v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBodyData) SetLockRepoBranch(v bool) *GetWorkspaceCodePublishSettingResponseBodyData {
	s.LockRepoBranch = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBodyData) SetRepos(v []*GetWorkspaceCodePublishSettingResponseBodyDataRepos) *GetWorkspaceCodePublishSettingResponseBodyData {
	s.Repos = v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBodyData) Validate() error {
	if s.Repos != nil {
		for _, item := range s.Repos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkspaceCodePublishSettingResponseBodyDataRepos struct {
	// example:
	//
	// main
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// example:
	//
	// /luna-public/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// analyticscomputing/dide
	Repo *string `json:"Repo,omitempty" xml:"Repo,omitempty"`
}

func (s GetWorkspaceCodePublishSettingResponseBodyDataRepos) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodePublishSettingResponseBodyDataRepos) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodePublishSettingResponseBodyDataRepos) GetBranch() *string {
	return s.Branch
}

func (s *GetWorkspaceCodePublishSettingResponseBodyDataRepos) GetPath() *string {
	return s.Path
}

func (s *GetWorkspaceCodePublishSettingResponseBodyDataRepos) GetRepo() *string {
	return s.Repo
}

func (s *GetWorkspaceCodePublishSettingResponseBodyDataRepos) SetBranch(v string) *GetWorkspaceCodePublishSettingResponseBodyDataRepos {
	s.Branch = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBodyDataRepos) SetPath(v string) *GetWorkspaceCodePublishSettingResponseBodyDataRepos {
	s.Path = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBodyDataRepos) SetRepo(v string) *GetWorkspaceCodePublishSettingResponseBodyDataRepos {
	s.Repo = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponseBodyDataRepos) Validate() error {
	return dara.Validate(s)
}
