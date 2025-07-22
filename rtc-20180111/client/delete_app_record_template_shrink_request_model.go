// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRecordTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppRecordTemplateShrinkRequest
	GetAppId() *string
	SetClientToken(v string) *DeleteAppRecordTemplateShrinkRequest
	GetClientToken() *string
	SetTemplateShrink(v string) *DeleteAppRecordTemplateShrinkRequest
	GetTemplateShrink() *string
}

type DeleteAppRecordTemplateShrinkRequest struct {
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
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DeleteAppRecordTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRecordTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRecordTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppRecordTemplateShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAppRecordTemplateShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *DeleteAppRecordTemplateShrinkRequest) SetAppId(v string) *DeleteAppRecordTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppRecordTemplateShrinkRequest) SetClientToken(v string) *DeleteAppRecordTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAppRecordTemplateShrinkRequest) SetTemplateShrink(v string) *DeleteAppRecordTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *DeleteAppRecordTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
