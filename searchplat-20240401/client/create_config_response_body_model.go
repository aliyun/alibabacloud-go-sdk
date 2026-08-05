// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateConfigResponseBody
	GetRequestId() *string
	SetResult(v *CreateConfigResponseBodyResult) *CreateConfigResponseBody
	GetResult() *CreateConfigResponseBodyResult
}

type CreateConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AAD430CC-D016-10BF-B837-8DA1EED87E94
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result object.
	Result *CreateConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigResponseBody) GetResult() *CreateConfigResponseBodyResult {
	return s.Result
}

func (s *CreateConfigResponseBody) SetRequestId(v string) *CreateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigResponseBody) SetResult(v *CreateConfigResponseBodyResult) *CreateConfigResponseBody {
	s.Result = v
	return s
}

func (s *CreateConfigResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConfigResponseBodyResult struct {
	// The configuration content.
	ConfigData map[string]interface{} `json:"configData,omitempty" xml:"configData,omitempty"`
	// The configuration type.
	//
	// example:
	//
	// prompt
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-06-18T07:04:42.877040
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-06-18T07:04:42.877040
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-001
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateConfigResponseBodyResult) GetConfigData() map[string]interface{} {
	return s.ConfigData
}

func (s *CreateConfigResponseBodyResult) GetConfigType() *string {
	return s.ConfigType
}

func (s *CreateConfigResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateConfigResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateConfigResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateConfigResponseBodyResult) SetConfigData(v map[string]interface{}) *CreateConfigResponseBodyResult {
	s.ConfigData = v
	return s
}

func (s *CreateConfigResponseBodyResult) SetConfigType(v string) *CreateConfigResponseBodyResult {
	s.ConfigType = &v
	return s
}

func (s *CreateConfigResponseBodyResult) SetCreatedAt(v string) *CreateConfigResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateConfigResponseBodyResult) SetUpdatedAt(v string) *CreateConfigResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *CreateConfigResponseBodyResult) SetWorkspaceId(v string) *CreateConfigResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *CreateConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
