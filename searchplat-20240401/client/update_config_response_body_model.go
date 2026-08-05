// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConfigResponseBody
	GetRequestId() *string
	SetResult(v *UpdateConfigResponseBodyResult) *UpdateConfigResponseBody
	GetResult() *UpdateConfigResponseBodyResult
}

type UpdateConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AAD430CC-D016-10BF-B837-8DA1EED87E94
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result object.
	Result *UpdateConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigResponseBody) GetResult() *UpdateConfigResponseBodyResult {
	return s.Result
}

func (s *UpdateConfigResponseBody) SetRequestId(v string) *UpdateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigResponseBody) SetResult(v *UpdateConfigResponseBodyResult) *UpdateConfigResponseBody {
	s.Result = v
	return s
}

func (s *UpdateConfigResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConfigResponseBodyResult struct {
	// The configuration content.
	ConfigData map[string]interface{} `json:"configData,omitempty" xml:"configData,omitempty"`
	// The configuration category.
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

func (s UpdateConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponseBodyResult) GetConfigData() map[string]interface{} {
	return s.ConfigData
}

func (s *UpdateConfigResponseBodyResult) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateConfigResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateConfigResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateConfigResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateConfigResponseBodyResult) SetConfigData(v map[string]interface{}) *UpdateConfigResponseBodyResult {
	s.ConfigData = v
	return s
}

func (s *UpdateConfigResponseBodyResult) SetConfigType(v string) *UpdateConfigResponseBodyResult {
	s.ConfigType = &v
	return s
}

func (s *UpdateConfigResponseBodyResult) SetCreatedAt(v string) *UpdateConfigResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateConfigResponseBodyResult) SetUpdatedAt(v string) *UpdateConfigResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateConfigResponseBodyResult) SetWorkspaceId(v string) *UpdateConfigResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
