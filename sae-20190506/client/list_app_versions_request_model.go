// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppVersionsRequest
	GetAppId() *string
}

type ListAppVersionsRequest struct {
	// The returned message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListAppVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppVersionsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppVersionsRequest) SetAppId(v string) *ListAppVersionsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppVersionsRequest) Validate() error {
	return dara.Validate(s)
}
