// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetApplicationVariablesRequest
	GetAppId() *string
}

type GetApplicationVariablesRequest struct {
	// example:
	//
	// Q2P4O9YSOKCQ35L9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariablesRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationVariablesRequest) SetAppId(v string) *GetApplicationVariablesRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationVariablesRequest) Validate() error {
	return dara.Validate(s)
}
