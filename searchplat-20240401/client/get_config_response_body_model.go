// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetConfigResponseBody
	GetRequestId() *string
	SetResult(v *GetConfigResponseBodyResult) *GetConfigResponseBody
	GetResult() *GetConfigResponseBodyResult
}

type GetConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7CC54C38-D721-4C55-A410-2A94B5A6BE0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *GetConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigResponseBody) GetResult() *GetConfigResponseBodyResult {
	return s.Result
}

func (s *GetConfigResponseBody) SetRequestId(v string) *GetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigResponseBody) SetResult(v *GetConfigResponseBodyResult) *GetConfigResponseBody {
	s.Result = v
	return s
}

func (s *GetConfigResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConfigResponseBodyResult struct {
	// The configuration content.
	ConfigData map[string]interface{} `json:"configData,omitempty" xml:"configData,omitempty"`
	// The configuration type.
	//
	// - prompt
	//
	// - lark
	//
	// example:
	//
	// lark
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// The time when the configuration was created.
	//
	// example:
	//
	// 2026-06-18T07:04:42.877040
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The time when the configuration was last updated.
	//
	// example:
	//
	// 2026-06-18T07:04:42.877040
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1201721
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetConfigResponseBodyResult) GetConfigData() map[string]interface{} {
	return s.ConfigData
}

func (s *GetConfigResponseBodyResult) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetConfigResponseBodyResult) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetConfigResponseBodyResult) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetConfigResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetConfigResponseBodyResult) SetConfigData(v map[string]interface{}) *GetConfigResponseBodyResult {
	s.ConfigData = v
	return s
}

func (s *GetConfigResponseBodyResult) SetConfigType(v string) *GetConfigResponseBodyResult {
	s.ConfigType = &v
	return s
}

func (s *GetConfigResponseBodyResult) SetCreatedAt(v int64) *GetConfigResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetConfigResponseBodyResult) SetUpdatedAt(v int64) *GetConfigResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *GetConfigResponseBodyResult) SetWorkspaceId(v string) *GetConfigResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *GetConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
