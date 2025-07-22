// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRecordTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppRecordTemplateShrinkRequest
	GetAppId() *string
	SetClientToken(v string) *CreateAppRecordTemplateShrinkRequest
	GetClientToken() *string
	SetRecordTemplateShrink(v string) *CreateAppRecordTemplateShrinkRequest
	GetRecordTemplateShrink() *string
}

type CreateAppRecordTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	RecordTemplateShrink *string `json:"RecordTemplate,omitempty" xml:"RecordTemplate,omitempty"`
}

func (s CreateAppRecordTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRecordTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRecordTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppRecordTemplateShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppRecordTemplateShrinkRequest) GetRecordTemplateShrink() *string {
	return s.RecordTemplateShrink
}

func (s *CreateAppRecordTemplateShrinkRequest) SetAppId(v string) *CreateAppRecordTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppRecordTemplateShrinkRequest) SetClientToken(v string) *CreateAppRecordTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRecordTemplateShrinkRequest) SetRecordTemplateShrink(v string) *CreateAppRecordTemplateShrinkRequest {
	s.RecordTemplateShrink = &v
	return s
}

func (s *CreateAppRecordTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
