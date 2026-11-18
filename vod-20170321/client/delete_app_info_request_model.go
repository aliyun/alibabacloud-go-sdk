// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppInfoRequest
	GetAppId() *string
}

type DeleteAppInfoRequest struct {
	// The application ID. You can obtain the value from the AppId response parameter returned by the [Create an application](https://help.aliyun.com/document_detail/113266.html) or [List applications](https://help.aliyun.com/document_detail/114000.html) API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppInfoRequest) SetAppId(v string) *DeleteAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
