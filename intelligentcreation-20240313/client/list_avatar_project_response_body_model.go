// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryAvatarProjectResultList(v []*ListAvatarProjectResponseBodyQueryAvatarProjectResultList) *ListAvatarProjectResponseBody
	GetQueryAvatarProjectResultList() []*ListAvatarProjectResponseBodyQueryAvatarProjectResultList
	SetRequestId(v string) *ListAvatarProjectResponseBody
	GetRequestId() *string
}

type ListAvatarProjectResponseBody struct {
	QueryAvatarProjectResultList []*ListAvatarProjectResponseBodyQueryAvatarProjectResultList `json:"queryAvatarProjectResultList,omitempty" xml:"queryAvatarProjectResultList,omitempty" type:"Repeated"`
	// example:
	//
	// D7F2B74F-63F2-5DD6-95E4-F408EAD6617E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAvatarProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectResponseBody) GetQueryAvatarProjectResultList() []*ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	return s.QueryAvatarProjectResultList
}

func (s *ListAvatarProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvatarProjectResponseBody) SetQueryAvatarProjectResultList(v []*ListAvatarProjectResponseBodyQueryAvatarProjectResultList) *ListAvatarProjectResponseBody {
	s.QueryAvatarProjectResultList = v
	return s
}

func (s *ListAvatarProjectResponseBody) SetRequestId(v string) *ListAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvatarProjectResponseBody) Validate() error {
	if s.QueryAvatarProjectResultList != nil {
		for _, item := range s.QueryAvatarProjectResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvatarProjectResponseBodyQueryAvatarProjectResultList struct {
	// example:
	//
	// 1000206
	AgentId  *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 12826084562688
	ProjectId   *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// DEPLOYING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAvatarProjectResponseBodyQueryAvatarProjectResultList) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarProjectResponseBodyQueryAvatarProjectResultList) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) GetStatus() *string {
	return s.Status
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetAgentId(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.AgentId = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetErrorMsg(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.ErrorMsg = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetProjectId(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.ProjectId = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetProjectName(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.ProjectName = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) SetStatus(v string) *ListAvatarProjectResponseBodyQueryAvatarProjectResultList {
	s.Status = &v
	return s
}

func (s *ListAvatarProjectResponseBodyQueryAvatarProjectResultList) Validate() error {
	return dara.Validate(s)
}
