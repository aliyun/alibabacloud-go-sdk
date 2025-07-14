// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourceStaticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryResourceStaticsRequest
	GetAppId() *string
}

type QueryResourceStaticsRequest struct {
	// 7171a6ca-d1cd-4928-8642-7d5cfe69\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s QueryResourceStaticsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryResourceStaticsRequest) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryResourceStaticsRequest) SetAppId(v string) *QueryResourceStaticsRequest {
	s.AppId = &v
	return s
}

func (s *QueryResourceStaticsRequest) Validate() error {
	return dara.Validate(s)
}
