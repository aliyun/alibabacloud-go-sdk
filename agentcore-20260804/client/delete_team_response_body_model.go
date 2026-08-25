// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTeamResponseBody
	GetCode() *string
	SetData(v *DeleteTeamResponseBodyData) *DeleteTeamResponseBody
	GetData() *DeleteTeamResponseBodyData
	SetHttpStatusCode(v int32) *DeleteTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTeamResponseBody
	GetSuccess() *bool
}

type DeleteTeamResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeleteTeamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTeamResponseBody) GetData() *DeleteTeamResponseBodyData {
	return s.Data
}

func (s *DeleteTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTeamResponseBody) SetCode(v string) *DeleteTeamResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTeamResponseBody) SetData(v *DeleteTeamResponseBodyData) *DeleteTeamResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTeamResponseBody) SetHttpStatusCode(v int32) *DeleteTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTeamResponseBody) SetMessage(v string) *DeleteTeamResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTeamResponseBody) SetRequestId(v string) *DeleteTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTeamResponseBody) SetSuccess(v bool) *DeleteTeamResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTeamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteTeamResponseBodyData struct {
	// example:
	//
	// team-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// tm-123456
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteTeamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteTeamResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteTeamResponseBodyData) GetTeamId() *string {
	return s.TeamId
}

func (s *DeleteTeamResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteTeamResponseBodyData) SetName(v string) *DeleteTeamResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteTeamResponseBodyData) SetStatus(v string) *DeleteTeamResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteTeamResponseBodyData) SetTeamId(v string) *DeleteTeamResponseBodyData {
	s.TeamId = &v
	return s
}

func (s *DeleteTeamResponseBodyData) SetWorkspaceId(v string) *DeleteTeamResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteTeamResponseBodyData) Validate() error {
	return dara.Validate(s)
}
