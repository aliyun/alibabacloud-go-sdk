// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppSecretRequest
	GetAppId() *string
}

type CreateAppSecretRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateAppSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSecretRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppSecretRequest) SetAppId(v string) *CreateAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppSecretRequest) Validate() error {
	return dara.Validate(s)
}
