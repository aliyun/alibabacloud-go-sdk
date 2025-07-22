// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRecordTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppRecordTemplateShrinkRequest
	GetAppId() *string
	SetClientToken(v string) *ModifyAppRecordTemplateShrinkRequest
	GetClientToken() *string
	SetRecordTemplateShrink(v string) *ModifyAppRecordTemplateShrinkRequest
	GetRecordTemplateShrink() *string
}

type ModifyAppRecordTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	RecordTemplateShrink *string `json:"RecordTemplate,omitempty" xml:"RecordTemplate,omitempty"`
}

func (s ModifyAppRecordTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppRecordTemplateShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppRecordTemplateShrinkRequest) GetRecordTemplateShrink() *string {
	return s.RecordTemplateShrink
}

func (s *ModifyAppRecordTemplateShrinkRequest) SetAppId(v string) *ModifyAppRecordTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRecordTemplateShrinkRequest) SetClientToken(v string) *ModifyAppRecordTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppRecordTemplateShrinkRequest) SetRecordTemplateShrink(v string) *ModifyAppRecordTemplateShrinkRequest {
	s.RecordTemplateShrink = &v
	return s
}

func (s *ModifyAppRecordTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
