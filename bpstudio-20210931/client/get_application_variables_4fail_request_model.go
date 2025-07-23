// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationVariables4FailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetApplicationVariables4FailRequest
	GetAppId() *string
}

type GetApplicationVariables4FailRequest struct {
	// example:
	//
	// Q2P4O9YSOKCT35L9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationVariables4FailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariables4FailRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationVariables4FailRequest) SetAppId(v string) *GetApplicationVariables4FailRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationVariables4FailRequest) Validate() error {
	return dara.Validate(s)
}
