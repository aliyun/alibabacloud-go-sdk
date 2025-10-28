// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebContainerConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetWebContainerConfigRequest
	GetAppId() *string
}

type GetWebContainerConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetWebContainerConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWebContainerConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWebContainerConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetWebContainerConfigRequest) SetAppId(v string) *GetWebContainerConfigRequest {
	s.AppId = &v
	return s
}

func (s *GetWebContainerConfigRequest) Validate() error {
	return dara.Validate(s)
}
