// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GenerateLabelRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GenerateLabelRequest
	GetJsonStr() *string
}

type GenerateLabelRequest struct {
	// The ID of the business space.
	//
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following details.
	//
	// example:
	//
	// “”
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GenerateLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateLabelRequest) GoString() string {
	return s.String()
}

func (s *GenerateLabelRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GenerateLabelRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GenerateLabelRequest) SetBaseMeAgentId(v int64) *GenerateLabelRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GenerateLabelRequest) SetJsonStr(v string) *GenerateLabelRequest {
	s.JsonStr = &v
	return s
}

func (s *GenerateLabelRequest) Validate() error {
	return dara.Validate(s)
}
