// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyAppPluginConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceBizId(v string) *CopyAppPluginConfigRequest
	GetSourceBizId() *string
	SetTargetBizId(v string) *CopyAppPluginConfigRequest
	GetTargetBizId() *string
}

type CopyAppPluginConfigRequest struct {
	// example:
	//
	// 31104757
	SourceBizId *string `json:"SourceBizId,omitempty" xml:"SourceBizId,omitempty"`
	// example:
	//
	// 31104758
	TargetBizId *string `json:"TargetBizId,omitempty" xml:"TargetBizId,omitempty"`
}

func (s CopyAppPluginConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyAppPluginConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyAppPluginConfigRequest) GetSourceBizId() *string {
	return s.SourceBizId
}

func (s *CopyAppPluginConfigRequest) GetTargetBizId() *string {
	return s.TargetBizId
}

func (s *CopyAppPluginConfigRequest) SetSourceBizId(v string) *CopyAppPluginConfigRequest {
	s.SourceBizId = &v
	return s
}

func (s *CopyAppPluginConfigRequest) SetTargetBizId(v string) *CopyAppPluginConfigRequest {
	s.TargetBizId = &v
	return s
}

func (s *CopyAppPluginConfigRequest) Validate() error {
	return dara.Validate(s)
}
