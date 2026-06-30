// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UploadDataV4Request
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UploadDataV4Request
	GetJsonStr() *string
}

type UploadDataV4Request struct {
	// The workspace ID. Use this parameter to select a specific workspace when you have multiple workspaces. If you do not specify this parameter, the default workspace is used.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The JSON string that contains all the custom parameters for this API. For more information, see the description of jsonStr properties below.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataV4Request) String() string {
	return dara.Prettify(s)
}

func (s UploadDataV4Request) GoString() string {
	return s.String()
}

func (s *UploadDataV4Request) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UploadDataV4Request) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UploadDataV4Request) SetBaseMeAgentId(v int64) *UploadDataV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadDataV4Request) SetJsonStr(v string) *UploadDataV4Request {
	s.JsonStr = &v
	return s
}

func (s *UploadDataV4Request) Validate() error {
	return dara.Validate(s)
}
