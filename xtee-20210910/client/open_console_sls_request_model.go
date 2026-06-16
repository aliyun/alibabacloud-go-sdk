// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenConsoleSlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *OpenConsoleSlsRequest
	GetLang() *string
	SetRegId(v string) *OpenConsoleSlsRequest
	GetRegId() *string
	SetScene(v string) *OpenConsoleSlsRequest
	GetScene() *string
}

type OpenConsoleSlsRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The scenario.
	//
	// example:
	//
	// SAF_DE_SERVICE
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s OpenConsoleSlsRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenConsoleSlsRequest) GoString() string {
	return s.String()
}

func (s *OpenConsoleSlsRequest) GetLang() *string {
	return s.Lang
}

func (s *OpenConsoleSlsRequest) GetRegId() *string {
	return s.RegId
}

func (s *OpenConsoleSlsRequest) GetScene() *string {
	return s.Scene
}

func (s *OpenConsoleSlsRequest) SetLang(v string) *OpenConsoleSlsRequest {
	s.Lang = &v
	return s
}

func (s *OpenConsoleSlsRequest) SetRegId(v string) *OpenConsoleSlsRequest {
	s.RegId = &v
	return s
}

func (s *OpenConsoleSlsRequest) SetScene(v string) *OpenConsoleSlsRequest {
	s.Scene = &v
	return s
}

func (s *OpenConsoleSlsRequest) Validate() error {
	return dara.Validate(s)
}
